package sam3

import (
	"errors"
	"net"
	"time"
)

// The DatagramSession implements net.PacketConn. It works almost like ordinary
// UDP, except that datagrams may be much larger (max 31kB). These datagrams are
// also end-to-end encrypted, signed and includes replay-protection. And they 
// are also built to be surveillance-resistant (yey!).
type DatagramSession struct {
	Addr		I2PAddr
	Priv		I2PKeys
	lport		int
}

// Implements net.PacketConn
func (s *DatagramSession) ReadFrom(b []byte) (n int, addr net.Addr, err error) {
	return 0, I2PAddr(""), errors.New("Not implemented.")
}

// Implements net.PacketConn
func (s *DatagramSession) WriteTo(b []byte, addr net.Addr) (n int, err error) {
	return 0, errors.New("Not implemented.")
}

// Implements net.PacketConn
func (s *DatagramSession) Close() error {
	return errors.New("Not implemented.")
}

// Implements net.PacketConn
func (s *DatagramSession) LocalAddr() net.Addr {
	return I2PAddr("")
}

// Implements net.PacketConn
func (s *DatagramSession) SetDeadline(t time.Time) error {
	return errors.New("Not implemented.")
}

// Implements net.PacketConn
func (s *DatagramSession) SetReadDeadline(t time.Time) error {
	return errors.New("Not implemented.")
}

// Implements net.PacketConn
func (s *DatagramSession) SetWriteDeadline(t time.Time) error {
	return errors.New("Not implemented.")
}


