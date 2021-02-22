package options

import "github.com/aprice/telnet"

// ECHO Telnet Echo Option - https://tools.ietf.org/html/rfc857

// EchoOption enables NAWS negotiation on a Server.
func EchoOption(c *telnet.Connection) telnet.Negotiator {
	return &EchoHandler{client: false}
}

// EchoHandler negotiates ECHO for a specific connection.
type EchoHandler struct {
	client bool
}

// OptionCode returns with the code used to negotiate ECHO modes.
func (e *EchoHandler) OptionCode() byte {
	return telnet.TeloptECHO
}

// Offer is called when a new connection is initiated. It offers the handler
// an opportunity to advertise or request an option.
func (e *EchoHandler) Offer(c *telnet.Connection) {
	if !e.client {
		c.Conn.Write([]byte{telnet.IAC, telnet.WILL, e.OptionCode()})
	}
}

// HandleDo is called when an IAC DO command is received for this option,
// indicating the client is requesting the option to be enabled.
func (e *EchoHandler) HandleDo(c *telnet.Connection) {

}

// HandleWill is called when an IAC WILL command is received for this
// option, indicating the client is willing to enable this option.
func (e *EchoHandler) HandleWill(c *telnet.Connection) {

}

// HandleSB is called when a subnegotiation command is received for this
// option. body contains the bytes between `IAC SB <OptionCode>` and `IAC
// SE`.
func (e *EchoHandler) HandleSB(c *telnet.Connection, body []byte) {

}
