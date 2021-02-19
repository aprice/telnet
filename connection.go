package telnet

import (
	"fmt"
	"net"
)

// Negotiator defines the requirements for a telnet option handler.
type Negotiator interface {
	// OptionCode returns the 1-byte option code that indicates this option.
	OptionCode() byte
	// Offer is called when a new connection is initiated. It offers the handler
	// an opportunity to advertise or request an option.
	Offer(conn *Connection)
	// HandleDo is called when an IAC DO command is received for this option,
	// indicating the client is requesting the option to be enabled.
	HandleDo(conn *Connection)
	// HandleWill is called when an IAC WILL command is received for this
	// option, indicating the client is willing to enable this option.
	HandleWill(conn *Connection)
	// HandleSB is called when a subnegotiation command is received for this
	// option. body contains the bytes between `IAC SB <OptionCode>` and `IAC
	// SE`.
	HandleSB(conn *Connection, body []byte)
}

// Connection to the telnet server. This lightweight TCPConn wrapper handles
// telnet control sequences transparently in reads and writes, and provides
// handling of supported options.
type Connection struct {
	// The underlying network connection.
	net.Conn

	// OptionHandlers handle IAC options; the key is the IAC option code.
	OptionHandlers map[byte]Negotiator

	// Read buffer
	buf  []byte
	r, w int // buf read and write positions

	// IAC handling
	iac    bool
	cmd    byte
	option byte

	// Known client wont/dont
	clientWont map[byte]bool
	clientDont map[byte]bool
}

// NewConnection initializes a new Connection for this given TCPConn. It will
// register all the given Option handlers and call Offer() on each, in order.
func NewConnection(c net.Conn, options []Option) *Connection {
	conn := &Connection{
		Conn:           c,
		OptionHandlers: make(map[byte]Negotiator, len(options)),
		buf:            make([]byte, 256),
		clientWont:     make(map[byte]bool),
		clientDont:     make(map[byte]bool),
	}
	for _, o := range options {
		h := o(conn)
		conn.OptionHandlers[h.OptionCode()] = h
		h.Offer(conn)
	}
	return conn
}

// Write to the connection, escaping IAC as necessary.
func (c *Connection) Write(b []byte) (n int, err error) {
	var nn, lastWrite int
	for i, ch := range b {
		if ch == IAC {
			if lastWrite < i-1 {
				nn, err = c.Conn.Write(b[lastWrite:i])
				n += nn
				if err != nil {
					return
				}
			}
			lastWrite = i + 1
			nn, err = c.Conn.Write([]byte{IAC, IAC})
			n += nn
			if err != nil {
				return
			}
		}
	}
	if lastWrite < len(b) {
		nn, err = c.Conn.Write(b[lastWrite:])
		n += nn
	}
	return
}

// RawWrite writes raw data to the connection, without escaping done by Write.
// Use of RawWrite over Conn.Write allows Connection to do any additional
// handling necessary, so long as it does not modify the raw data sent.
func (c *Connection) RawWrite(b []byte) (n int, err error) {
	return c.Conn.Write(b)
}

const maxReadAttempts = 10

// Read from the connection, transparently removing and handling IAC control
// sequences. It may attempt multiple reads against the underlying connection if
// it receives back only IAC which gets stripped out of the stream.
func (c *Connection) Read(b []byte) (n int, err error) {
	for i := 0; i < maxReadAttempts && n == 0 && len(b) > 0; i++ {
		n, err = c.read(b)
	}
	return
}

func (c *Connection) read(b []byte) (n int, err error) {
	err = c.fill(len(b))
	if err != nil {
		return
	}
	var lastWrite, subStart int
	var ignoreIAC bool
	write := func(end int) int {
		if c.r == end {
			return 0
		}
		nn := copy(b[lastWrite:], c.buf[c.r:end])
		n += nn
		lastWrite += nn
		c.r += nn
		return nn
	}
	endIAC := func(i int) {
		subStart = 0
		c.iac = false
		c.cmd = 0
		c.option = 0
		c.r = i + 1
	}
	for i := c.r; i < c.w && lastWrite < len(b); i++ {
		ch := c.buf[i]
		if ch == IAC && !ignoreIAC {
			if c.iac && c.cmd == 0 {
				// Escaped IAC in text
				write(i)
				c.r++
				c.iac = false
				continue
			} else if c.iac && c.buf[i-1] == IAC {
				// Escaped IAC inside IAC sequence
				copy(c.buf[:i], c.buf[i+1:])
				i--
				ignoreIAC = true
				continue
			} else if !c.iac {
				// Start of IAC sequence
				write(i)
				c.iac = true
				continue
			}
		}

		ignoreIAC = false

		if c.iac && c.cmd == 0 {
			c.cmd = ch
			if ch == SB {
				subStart = i + 2
			}
			continue
		} else if c.iac && c.option == 0 {
			c.option = ch
			if c.cmd != SB {
				c.handleNegotiation()
				endIAC(i)
			}
			continue
		} else if c.iac && c.cmd == SB && ch == SE && c.buf[i-1] == IAC {
			if h, ok := c.OptionHandlers[c.option]; ok {
				h.HandleSB(c, c.buf[subStart:i-1])
			}
			endIAC(i)
			continue
		}
	}

	nn := copy(b[lastWrite:], c.buf[c.r:c.w])
	n += nn
	c.r += nn
	return
}

func (c *Connection) fill(requestedBytes int) error {
	if c.r > 0 {
		copy(c.buf, c.buf[c.r:])
		c.w -= c.r
		c.r = 0
	}

	if len(c.buf) < requestedBytes {
		buf := make([]byte, requestedBytes)
		c.w = copy(buf, c.buf[c.r:c.w])
		c.r = 0
		c.buf = buf
	}

	nn, err := c.Conn.Read(c.buf[c.w:])
	c.w += nn
	return err
}

// SetWindowTitle attempts to set the client's telnet window title. Clients may
// or may not support this.
func (c *Connection) SetWindowTitle(title string) {
	fmt.Fprintf(c, TitleBarFmt, title)
}

func (c *Connection) handleNegotiation() {
	switch c.cmd {
	case WILL:
		if h, ok := c.OptionHandlers[c.option]; ok {
			h.HandleWill(c)
		} else {
			c.Conn.Write([]byte{IAC, DONT, c.option})
		}
	case WONT:
		c.clientWont[c.option] = true
	case DO:
		if h, ok := c.OptionHandlers[c.option]; ok {
			h.HandleDo(c)
		} else {
			c.Conn.Write([]byte{IAC, WONT, c.option})
		}
	case DONT:
		c.clientDont[c.option] = true
	}
}
