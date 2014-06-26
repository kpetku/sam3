package sam3



import (
	"fmt"
	"strings"
	"testing"
//	"time"
)



const yoursam = "127.0.0.1:7656"



func Test_Basic(t *testing.T) {
	fmt.Println("Test_Basic")
	fmt.Println("\tAttaching to SAM at " + yoursam)
	sam, err := NewSAM(yoursam)
	if err != nil {
		fmt.Println(err.Error)
		t.Fail()
		return
	}
	
	fmt.Println("\tCreating new keys...")
	addr, keys, err := sam.NewKeys()
	if err != nil {
		fmt.Println(err.Error())
		t.Fail()
	} else {
		fmt.Println("\tAddress created: " + addr.Base32())
		fmt.Println("\tI2PKeys: " + string(keys.priv)[:50] + "(...etc)")
	}
	
	addr2, err := sam.Lookup("zzz.i2p")
	if err != nil {
		fmt.Println(err.Error())
		t.Fail()
	} else {
		fmt.Println("\tzzz.i2p = " + addr2.Base32())
	}		

	if err := sam.Close(); err != nil {
		fmt.Println(err.Error())
		t.Fail()
	}
}


/*
func Test_GenericSession(t *testing.T) {
	if testing.Short() {
		return
	}
	fmt.Println("Test_GenericSession")
	sam, err := NewSAM(yoursam)
	if err != nil {
		fmt.Println(err.Error)
		t.Fail()
		return
	}
	_, keys, err := sam.NewKeys()
	if err != nil {
		fmt.Println(err.Error())
		t.Fail()
	} else {
		conn1, err := sam.newGenericSession("STREAM", "testTun", keys, []string{})
		if err != nil {
			fmt.Println(err.Error())
			t.Fail()
		} else {
			conn1.Close()
		}
		conn2, err := sam.newGenericSession("STREAM", "testTun", keys, []string{"inbound.length=1", "outbound.length=1", "inbound.lengthVariance=1", "outbound.lengthVariance=1", "inbound.quantity=1", "outbound.quantity=1"})
		if err != nil {
			fmt.Println(err.Error())
			t.Fail()
		} else {
			conn2.Close()
		}
		conn3, err := sam.newGenericSession("DATAGRAM", "testTun", keys, []string{"inbound.length=1", "outbound.length=1", "inbound.lengthVariance=1", "outbound.lengthVariance=1", "inbound.quantity=1", "outbound.quantity=1"})
		if err != nil {
			fmt.Println(err.Error())
			t.Fail()
		} else {
			conn3.Close()
		}
	}
	if err := sam.Close(); err != nil {
		fmt.Println(err.Error())
		t.Fail()
	}
}
*/


func Test_StreamingDial(t *testing.T) {
	if testing.Short() {
		return
	}
	fmt.Println("Test_StreamingDial")
	sam, err := NewSAM(yoursam)
	if err != nil {
		fmt.Println(err.Error)
		t.Fail()
		return
	}
	defer sam.Close()
	_, keys, err := sam.NewKeys()
	if err != nil {
		fmt.Println(err.Error())
		t.Fail()
		return
	}
	fmt.Println("\tBuilding tunnel")
	ss, err := sam.NewStreamSession("streamTun", keys, []string{"inbound.length=0", "outbound.length=0", "inbound.lengthVariance=0", "outbound.lengthVariance=0", "inbound.quantity=1", "outbound.quantity=1"})
	if err != nil {
		fmt.Println(err.Error())
		t.Fail()
		return
	}
	fmt.Println("\tLooking up forum.i2p")
	forumAddr, err := sam.Lookup("forum.i2p")
	if err != nil {
		fmt.Println(err.Error())
		t.Fail()
		return
	}
	fmt.Println("\tDialing forum.i2p")
	conn, err := ss.DialI2P(forumAddr)
	if err != nil {
		fmt.Println(err.Error())
		t.Fail()
		return
	}
	defer conn.Close()
	fmt.Println("\tSending HTTP GET /")
	if _, err := conn.Write([]byte("GET /\n")); err != nil {
		fmt.Println(err.Error())
		t.Fail()
		return	
	}
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if !strings.Contains(strings.ToLower(string(buf[:n])), "http") && !strings.Contains(strings.ToLower(string(buf[:n])), "html") {
		fmt.Printf("\tProbably failed to StreamSession.DialI2P(forum.i2p)? It replied %d bytes, but nothing that looked like http/html", n)
	} else {
		fmt.Println("\tRead HTTP/HTML from forum.i2p")
	}
}

func Test_StreamingServerClient(t *testing.T) {
	if testing.Short() {
		return
	}
	fmt.Println("Test_StreamingServerClient")
	sam, err := NewSAM(yoursam)
	if err != nil {
		t.Fail()
		return
	}
	defer sam.Close()
	_, keys, err := sam.NewKeys()
	if err != nil {
		t.Fail()
		return
	}
	fmt.Println("\tServer: Creating tunnel")
	ss, err := sam.NewStreamSession("serverTun", keys, []string{"inbound.length=0", "outbound.length=0", "inbound.lengthVariance=0", "outbound.lengthVariance=0", "inbound.quantity=1", "outbound.quantity=1"})
	if err != nil {
		return
	}
	c, w := make(chan bool), make(chan bool)
	go func(c, w chan(bool)) { 
		if !(<-w) {
			return
		}
		sam2, err := NewSAM(yoursam)
		if err != nil {
			c <- false
			return
		}
		defer sam2.Close()
		_, keys, err := sam2.NewKeys()
		if err != nil {
			c <- false
			return
		}
		fmt.Println("\tClient: Creating tunnel")
		ss2, err := sam2.NewStreamSession("clientTun", keys, []string{"inbound.length=0", "outbound.length=0", "inbound.lengthVariance=0", "outbound.lengthVariance=0", "inbound.quantity=1", "outbound.quantity=1"})
		if err != nil {
			c <- false
			return
		}
		fmt.Println("\tClient: Connecting to server")
		conn, err := ss2.DialI2P(ss.Addr())
		if err != nil {
			c <- false
			return
		}
		fmt.Println("\tClient: Connected to tunnel")
		defer conn.Close()
		_, err = conn.Write([]byte("Hello world <3 <3 <3 <3 <3 <3"))
		if err != nil {
			c <- false
			return
		}
		c <- true
	}(c, w)
	l, err := ss.Listen()
	if err != nil {
		fmt.Println("ss.Listen(): " + err.Error())
		t.Fail()
		w <- false
		return
	}
	w <- true
	fmt.Println("\tServer: Accept()ing on tunnel")
	conn, err := l.Accept()
	buf := make([]byte, 512)
	n,err := conn.Read(buf)
	fmt.Printf("\tClient exited successfully: %t\n", <-c)
	fmt.Println("\tServer: received from Client: " + string(buf[:n]))
}



