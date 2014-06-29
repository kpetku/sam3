package sam3



import (
	"fmt"
	"strings"
	"testing"
	"time"
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
	keys, err := sam.NewKeys()
	if err != nil {
		fmt.Println(err.Error())
		t.Fail()
	} else {
		fmt.Println("\tAddress created: " + keys.Addr().Base32())
		fmt.Println("\tI2PKeys: " + string(keys.both)[:50] + "(...etc)")
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
	keys, err := sam.NewKeys()
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
	keys, err := sam.NewKeys()
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
	fmt.Println("\tNotice: This may fail if your I2P node is not well integrated in the I2P network.")
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
	keys, err := sam.NewKeys()
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
		keys, err := sam2.NewKeys()
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
	defer l.Close()
	w <- true
	fmt.Println("\tServer: Accept()ing on tunnel")
	conn, err := l.Accept()
	if err != nil {
		t.Fail()
		fmt.Println("Failed to Accept(): " + err.Error())
		return
	}
	defer conn.Close()
	buf := make([]byte, 512)
	n,err := conn.Read(buf)
	fmt.Printf("\tClient exited successfully: %t\n", <-c)
	fmt.Println("\tServer: received from Client: " + string(buf[:n]))
}



func Test_DatagramServerClient(t *testing.T) {
	if testing.Short() {
		return
	}

	fmt.Println("Test_DatagramServerClient")
	sam, err := NewSAM(yoursam)
	if err != nil {
		t.Fail()
		return
	}
	defer sam.Close()
	keys, err := sam.NewKeys()
	if err != nil {
		t.Fail()
		return
	}
//	fmt.Println("\tServer: My address: " + keys.Addr().Base32())
	fmt.Println("\tServer: Creating tunnel")
	ds, err := sam.NewDatagramSession("DGserverTun", keys, []string{"inbound.length=0", "outbound.length=0", "inbound.lengthVariance=0", "outbound.lengthVariance=0", "inbound.quantity=1", "outbound.quantity=1"}, 0)
	if err != nil {
		fmt.Println("Server: Failed to create tunnel: " + err.Error())
		t.Fail()
		return
	}
	c, w := make(chan bool), make(chan bool)
	go func(c, w chan(bool)) { 
		sam2, err := NewSAM(yoursam)
		if err != nil {
			c <- false
			return
		}
		defer sam2.Close()
		keys, err := sam2.NewKeys()
		if err != nil {
			c <- false
			return
		}
		fmt.Println("\tClient: Creating tunnel")
		ds2, err := sam2.NewDatagramSession("DGclientTun", keys, []string{"inbound.length=0", "outbound.length=0", "inbound.lengthVariance=0", "outbound.lengthVariance=0", "inbound.quantity=1", "outbound.quantity=1"}, 0)
		if err != nil {
			c <- false
			return
		}
		defer ds2.Close()
//		fmt.Println("\tClient: Servers address: " + ds.LocalAddr().Base32())
//		fmt.Println("\tClient: Clients address: " + ds2.LocalAddr().Base32())
		fmt.Println("\tClient: Tries to send datagram to server")
		for {
			select {
				default :
					_, err = ds2.WriteTo([]byte("Hello datagram-world! <3 <3 <3 <3 <3 <3"), ds.LocalAddr())
					if err != nil {
						fmt.Println("\tClient: Failed to send datagram: " + err.Error())
						c <- false
						return
					}
					time.Sleep(5 * time.Second)
				case <-w :
					fmt.Println("\tClient: Sent datagram, quitting.")
					return
			}
		}
		c <- true
	}(c, w)
	buf := make([]byte, 512)
	fmt.Println("\tServer: ReadFrom() waiting...")
	n, _, err := ds.ReadFrom(buf)
	w <- true
	if err != nil {
		fmt.Println("\tServer: Failed to ReadFrom(): " + err.Error())
		t.Fail()
		return
	}
	fmt.Println("\tServer: Received datagram: " + string(buf[:n]))
//	fmt.Println("\tServer: Senders address was: " + saddr.Base32())
}



func Test_RawServerClient(t *testing.T) {
	if testing.Short() {
		return
	}

	fmt.Println("Test_RawServerClient")
	sam, err := NewSAM(yoursam)
	if err != nil {
		t.Fail()
		return
	}
	defer sam.Close()
	keys, err := sam.NewKeys()
	if err != nil {
		t.Fail()
		return
	}
	fmt.Println("\tServer: Creating tunnel")
	rs, err := sam.NewDatagramSession("RAWserverTun", keys, []string{"inbound.length=0", "outbound.length=0", "inbound.lengthVariance=0", "outbound.lengthVariance=0", "inbound.quantity=1", "outbound.quantity=1"}, 0)
	if err != nil {
		fmt.Println("Server: Failed to create tunnel: " + err.Error())
		t.Fail()
		return
	}
	c, w := make(chan bool), make(chan bool)
	go func(c, w chan(bool)) { 
		sam2, err := NewSAM(yoursam)
		if err != nil {
			c <- false
			return
		}
		defer sam2.Close()
		keys, err := sam2.NewKeys()
		if err != nil {
			c <- false
			return
		}
		fmt.Println("\tClient: Creating tunnel")
		rs2, err := sam2.NewDatagramSession("RAWclientTun", keys, []string{"inbound.length=0", "outbound.length=0", "inbound.lengthVariance=0", "outbound.lengthVariance=0", "inbound.quantity=1", "outbound.quantity=1"}, 0)
		if err != nil {
			c <- false
			return
		}
		defer rs2.Close()
		fmt.Println("\tClient: Tries to send raw datagram to server")
		for {
			select {
				default :
					_, err = rs2.WriteTo([]byte("Hello raw-world! <3 <3 <3 <3 <3 <3"), rs.LocalAddr())
					if err != nil {
						fmt.Println("\tClient: Failed to send raw datagram: " + err.Error())
						c <- false
						return
					}
					time.Sleep(5 * time.Second)
				case <-w :
					fmt.Println("\tClient: Sent raw datagram, quitting.")
					return
			}
		}
		c <- true
	}(c, w)
	buf := make([]byte, 512)
	fmt.Println("\tServer: Read() waiting...")
	n, _, err := rs.ReadFrom(buf)
	w <- true
	if err != nil {
		fmt.Println("\tServer: Failed to Read(): " + err.Error())
		t.Fail()
		return
	}
	fmt.Println("\tServer: Received datagram: " + string(buf[:n]))
//	fmt.Println("\tServer: Senders address was: " + saddr.Base32())
}

