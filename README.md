# README #

go library for the I2P SAMv3 bridge, used to build anonymous/pseudonymous end-to-end encrypted sockets.

This library is much better than ccondom (that use BOB), much more stable and much easier to maintain.

## Support/TODO ##

**What works:**

* Utils
    * Resolving URLs to I2P destinations
    * .b32.i2p hashes
    * Generating keys/i2p destinations
* Streaming
    * DialI2P() - Connecting to stuff in I2P
    * Listen()/Accept() - Handling incomming connections
    * Implements net.Conn and net.Listener
* Datagrams
    * Implements net.PacketConn

**Does not work:**

* Raw packets

## Documentation ##

* Enter `godoc -http=:8081` into your terminal and hit enter.
* Goto http://localhost:8081, click packages, and navigate to sam3

## Testing ##

* `go test` runs the whole suite
* `go test -short` runs the shorter variant

## License ##

Public domain.

## Author ##

Kalle Vedin `kalle.vedin@fripost.org`
