package telnet

import (
	"net"
)

// Dial establishes a telnet connection with the remote host specified by addr
// in host:port format. Any specified option handlers will be applied to the
// connection if it is successful.
func Dial(addr string, options ...Option) (conn *Connection, err error) {
	c, err := net.Dial("tcp", addr)
	if err != nil {
		return
	}
	conn = NewConnection(c, options)
	return
}
