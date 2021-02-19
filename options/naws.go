package options

// NAWS - Negotiate About Window Size - https://tools.ietf.org/html/rfc1073

import (
	"encoding/binary"
	"os"
	"time"

	"github.com/aprice/telnet"
	"golang.org/x/crypto/ssh/terminal"
)

// NAWSOption enables NAWS negotiation on a Server.
func NAWSOption(c *telnet.Connection) telnet.Negotiator {
	return &NAWSHandler{client: false}
}

// ExposeNAWS enables NAWS negotiation on a Client.
func ExposeNAWS(c *telnet.Connection) telnet.Negotiator {
	width, height, _ := terminal.GetSize(int(os.Stdin.Fd()))
	return &NAWSHandler{Width: uint16(width), Height: uint16(height), client: true}
}

// NAWSHandler negotiates NAWS for a specific connection.
type NAWSHandler struct {
	Width  uint16
	Height uint16

	client bool
}

// OptionCode returns the IAC code for NAWS.
func (n *NAWSHandler) OptionCode() byte {
	return telnet.TeloptNAWS
}

// Offer sends the IAC DO NAWS command to the client.
func (n *NAWSHandler) Offer(c *telnet.Connection) {
	if !n.client {
		c.Conn.Write([]byte{telnet.IAC, telnet.DO, n.OptionCode()})
	}
}

// HandleWill is not implemented for NAWS.
func (n *NAWSHandler) HandleWill(c *telnet.Connection) {}

// HandleDo processes the monitor size options for NAWS.
func (n *NAWSHandler) HandleDo(c *telnet.Connection) {
	if n.client {
		c.Conn.Write([]byte{telnet.IAC, telnet.WILL, n.OptionCode()})
		n.writeSize(c)
		go n.monitorTTYSize(c)
	} else {
		c.Conn.Write([]byte{telnet.IAC, telnet.WONT, n.OptionCode()})
	}
}

func (n *NAWSHandler) monitorTTYSize(c *telnet.Connection) {
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

func (n *NAWSHandler) writeSize(c *telnet.Connection) {
	c.Conn.Write([]byte{telnet.IAC, telnet.SB, n.OptionCode()})
	payload := make([]byte, 4)
	binary.BigEndian.PutUint16(payload, n.Width)
	binary.BigEndian.PutUint16(payload[2:], n.Height)
	// Normal write - we want inadvertent telnet.IACs to be escaped in body
	c.Write(payload)
	c.Conn.Write([]byte{telnet.IAC, telnet.SE})
}

// HandleSB processes the information about window size sent from the client to the server.
func (n *NAWSHandler) HandleSB(c *telnet.Connection, b []byte) {
	if !n.client {
		n.Width = binary.BigEndian.Uint16(b[0:2])
		n.Height = binary.BigEndian.Uint16(b[2:4])
	}
}
