package sam3

import (
	"time"
	"net"
)

type SAMConn struct {
	laddr  I2PAddr
	raddr  I2PAddr
	conn   net.Conn
}

func (sc SAMConn) Read(buf []byte) (int, error) {
	n, err := sc.conn.Read(buf)
	return n, err
}

func (sc SAMConn) Write(buf []byte) (int, error) {
	n, err := sc.conn.Write(buf)
	return n, err
}

func (sc SAMConn) Close() error {
	return sc.conn.Close()
}

func (sc SAMConn) LocalAddr() I2PAddr {
	return sc.laddr
}

func (sc SAMConn) RemoteAddr() I2PAddr {
	return sc.raddr
}

func (sc SAMConn) SetDeadline(t time.Time) error {
	return sc.conn.SetDeadline(t)
}

func (sc SAMConn) SetReadDeadline(t time.Time) error {
	return sc.conn.SetReadDeadline(t)
}

func (sc SAMConn) SetWriteDeadline(t time.Time) error {
	return sc.conn.SetWriteDeadline(t)
}


