package sam3

import (
	"crypto/sha256"
	"encoding/base32"
	"encoding/base64"
	"errors"
)



var (
	i2pB64enc *base64.Encoding = base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-~")
	i2pB32enc *base32.Encoding = base32.NewEncoding("abcdefghijklmnopqrstuvwxyz234567")
)




// Base64 representation of the private keys
type I2PKeys struct {
	addr    I2PAddr
	priv    string
}

// Returns the string of I2PKeys (Base64-encoded)
func (k I2PKeys) String() string {
	return k.priv
}

// Returns the I2PAddr (I2P destination) belonging to the I2PKeys.
func (k I2PKeys) Addr() I2PAddr {
	return k.addr
}

// Base64 representation of the public keys
type I2PAddr string

// Returns the base64 representation of the I2PAddr
func (a I2PAddr) Base64() string {
	return string(a)
}

// Returns the I2P destination (base64-encoded)
func (a I2PAddr) String() string {
	return string(a)
}

// Returns "I2P"
func (a I2PAddr) Network() string {
	return "I2P"
}

// Creates a new I2P address from a base64-encoded string.
func NewI2PAddrFromString(addr string) (I2PAddr, error) {
	// very basic check
	if len(addr) > 4096 || len(addr) < 516 {
		return I2PAddr(""), errors.New("Not an I2P address")
	}
	buf := make([]byte, 4096)
	if _, err := i2pB64enc.Decode(buf, []byte(addr)); err != nil {
		return I2PAddr(""), errors.New("Address is not base64-encoded")
	}
	return I2PAddr(addr), nil
}

// Creates a new I2P address from a byte array (pure binary, as you would have 
// if you have *not* base64-encoded it.)
func NewI2PAddrFromBytes(addr []byte) (I2PAddr, error) {
	if len(addr) > 4096 || len(addr) < 384 {
		return I2PAddr(""), errors.New("Not an I2P address")
	}
	buf := make([]byte, i2pB64enc.EncodedLen(len(addr)))
	i2pB64enc.Encode(buf, addr)
	return I2PAddr(string(buf)), nil
}

// Returns the *.b32.i2p address of the I2P address.
func (addr I2PAddr) Base32() string {
	hash := sha256.New()
	hash.Write([]byte(string(addr)))
	digest := hash.Sum(nil)
	b32addr := make([]byte, 56)
	i2pB32enc.Encode(b32addr, digest)
	return string(b32addr[:52]) + ".b32.i2p"
}

// Shortcut to I2PAddr.Base32()
func Base32(addr I2PAddr) string {
	return addr.Base32()
}
