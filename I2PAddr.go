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


// The public and private keys associated with an I2P destination. I2P hides the
// details of exactly what this is, so treat them as blobs, but generally: One
// pair of DSA keys, one pair of ElGamal keys, and sometimes (almost never) also
// a certificate. Addr() and Priv() returns you the full content of I2PKeys, 
// which you could/should store to disk if you want to be able to use the I2P 
// destination again in the future.
type I2PKeys struct {
	addr    I2PAddr
	priv    string
}

// Creates I2PKeys from an I2PAddr and a private key. The I2PAddr obviously must
// be the public key belonging to the private key. Performs no error checking.
func NewKeys(addr I2PAddr, priv string) I2PKeys {
	return I2PKeys{addr, priv}
}

// Returns the public keys of the I2PKeys.
func (k I2PKeys) Addr() I2PAddr {
	return k.addr
}

// Returns the private keys, in I2Ps base64 format.
func (k I2PKeys) Priv() string {
	return k.priv
}

// I2PAddr represents an I2P destination, almost equivalent to an IP address.
// This is the humongously huge base64 representation of such an address, which
// really is just a pair of public keys and also maybe a certificate. (I2P hides
// the details of exactly what it is. Read the I2P specifications for more info.)
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

// Creates a new I2P address from a base64-encoded string. Checks if the address
// addr is in correct format. (If you know for sure it is, use I2PAddr(addr).)
func NewI2PAddrFromString(addr string) (I2PAddr, error) {
	// very basic check
	if len(addr) > 4096 || len(addr) < 516 {
		return I2PAddr(""), errors.New("Not an I2P address")
	}
	buf := make([]byte, i2pB64enc.DecodedLen(len(addr)))
	if _, err := i2pB64enc.Decode(buf, []byte(addr)); err != nil {
		return I2PAddr(""), errors.New("Address is not base64-encoded")
	}
	return I2PAddr(addr), nil
}

// Creates a new I2P address from a byte array. The inverse of ToBytes().
func NewI2PAddrFromBytes(addr []byte) (I2PAddr, error) {
	if len(addr) > 4096 || len(addr) < 384 {
		return I2PAddr(""), errors.New("Not an I2P address")
	}
	buf := make([]byte, i2pB64enc.EncodedLen(len(addr)))
	i2pB64enc.Encode(buf, addr)
	return I2PAddr(string(buf)), nil
}

// Turns an I2P address to a byte array. The inverse of NewI2PAddrFromBytes().
func (addr I2PAddr) ToBytes() ([]byte, error) {
	buf := make([]byte, i2pB64enc.DecodedLen(len(addr)))
	if _, err := i2pB64enc.Decode(buf, []byte(addr)); err != nil {
		return buf, errors.New("Address is not base64-encoded")
	}
	return buf, nil
}

// Returns the *.b32.i2p address of the I2P address. It is supposed to be a 
// somewhat human-manageable 64 character long pseudo-domain name equivalent of 
// the 516+ characters long default base64-address (the I2PAddr format). It is 
// not possible to turn the base32-address back into a usable I2PAddr without 
// performing a Lookup(). Lookup only works if you are using the I2PAddr from
// which the b32 address was generated.
func (addr I2PAddr) Base32() string {
	hash := sha256.New()
	hash.Write([]byte(string(addr)))
	digest := hash.Sum(nil)
	b32addr := make([]byte, 56)
	i2pB32enc.Encode(b32addr, digest)
	return string(b32addr[:52]) + ".b32.i2p"
}

// Makes any string into a *.b32.i2p human-readable I2P address. This makes no
// sense, unless "anything" is an I2P destination of some sort.
func Base32(anything string) string {
	return I2PAddr(anything).Base32()
}
