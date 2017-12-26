package telnet

import (
	"encoding/binary"
	"os"
	"time"

	"golang.org/x/crypto/ssh/terminal"
)

const (
	ECHO    = byte(1)
	TTYPE   = byte(24)
	NAWS    = byte(31)
	ENCRYPT = byte(38)
	EOR     = byte(239)
)

// NAWSOption enables NAWS negotiation on a Server.
func NAWSOption(c *Connection) Negotiator {
	return &NAWSHandler{client: false}
}

// ExposeNAWS enables NAWS negotiation on a Client.
func ExposeNAWS(c *Connection) Negotiator {
	width, height, _ := terminal.GetSize(int(os.Stdin.Fd()))
	return &NAWSHandler{Width: uint16(width), Height: uint16(height), client: true}
}

// NAWSHandler negotiates NAWS for a specific connection.
type NAWSHandler struct {
	Width  uint16
	Height uint16

	client bool
}

func (n *NAWSHandler) OptionCode() byte {
	return NAWS
}

func (n *NAWSHandler) Offer(c *Connection) {
	if !n.client {
		c.Conn.Write([]byte{IAC, DO, n.OptionCode()})
	}
}

func (n *NAWSHandler) HandleWill(c *Connection) {}

func (n *NAWSHandler) HandleDo(c *Connection) {
	if n.client {
		c.Conn.Write([]byte{IAC, WILL, n.OptionCode()})
		n.writeSize(c)
		go n.monitorTTYSize(c)
	} else {
		c.Conn.Write([]byte{IAC, WONT, n.OptionCode()})
	}
}

func (n *NAWSHandler) monitorTTYSize(c *Connection) {
	t := time.NewTicker(time.Second)
	for range t.C {
		w, h, err := terminal.GetSize(int(os.Stdin.Fd()))
		if err != nil {
			continue
		}
		width := uint16(w)
		height := uint16(h)
		if width != n.Width || height != n.Height {
			n.Width = width
			n.Height = height
			n.writeSize(c)
		}
	}
}

func (n *NAWSHandler) writeSize(c *Connection) {
	c.Conn.Write([]byte{IAC, SB, n.OptionCode()})
	payload := make([]byte, 4)
	binary.BigEndian.PutUint16(payload, n.Width)
	binary.BigEndian.PutUint16(payload[2:], n.Height)
	// Normal write - we want inadvertent IACs to be escaped in body
	c.Write(payload)
	c.Conn.Write([]byte{IAC, SE})
}

func (n *NAWSHandler) HandleSB(c *Connection, b []byte) {
	if !n.client {
		n.Width = binary.BigEndian.Uint16(b[0:2])
		n.Height = binary.BigEndian.Uint16(b[2:4])
	}
}
