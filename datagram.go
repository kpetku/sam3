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
	keys        I2PKeys
	lport       int
}

// Creates a new datagram session.
func (sam *SAM) NewDatagramSession(tunnelName string, keys I2PKeys, options []string) (*DatagramSession, error) {
	return nil, errors.New("Not implemented.")
}

// Reads one datagram sent to the destination of the DatagramSession. Returns 
// the number of bytes read, from what address it was sent, or an error.
func (s *DatagramSession) ReadFrom(b []byte) (n int, addr I2PAddr, err error) {
	return 0, I2PAddr(""), errors.New("Not implemented.")
}

// Sends one signed datagram to the destination specified. At the time of 
// writing, maximum size is 31 kilobyte, but this may change in the future.
// Implements net.PacketConn.
func (s *DatagramSession) WriteTo(b []byte, addr I2PAddr) (n int, err error) {
	return 0, errors.New("Not implemented.")
}

// Closes the DatagramSession. Implements net.PacketConn
func (s *DatagramSession) Close() error {
	return errors.New("Not implemented.")
}

// Returns the I2P destination of the DatagramSession. Implements net.PacketConn
func (s *DatagramSession) LocalAddr() net.Addr {
	return s.keys.Addr()
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


