package options

import "github.com/aprice/telnet"

// TerminalType Telnet Option - https://tools.ietf.org/html/rfc857

// TerminalTypeOption enables NAWS negotiation on a Server.
func TerminalTypeOption(c *telnet.Connection) telnet.Negotiator {
	return &TerminalTypeHandler{client: false}
}

// TerminalTypeHandler negotiates TerminalType for a specific connection.
type TerminalTypeHandler struct {
	client bool
}

// OptionCode returns with the code used to negotiate TerminalType modes.
func (e *TerminalTypeHandler) OptionCode() byte {
	return telnet.TeloptTTYPE
}

// Offer is called when a new connection is initiated. It offers the handler
// an opportunity to advertise or request an option.
func (e *TerminalTypeHandler) Offer(c *telnet.Connection) {
	if !e.client {
		c.Conn.Write([]byte{telnet.IAC, telnet.WILL, e.OptionCode()})
	}
}

// HandleDo is called when an IAC DO command is received for this option,
// indicating the client is requesting the option to be enabled.
func (e *TerminalTypeHandler) HandleDo(c *telnet.Connection) {

}

// HandleWill is called when an IAC WILL command is received for this
// option, indicating the client is willing to enable this option.
func (e *TerminalTypeHandler) HandleWill(c *telnet.Connection) {

}

// HandleSB is called when a subnegotiation command is received for this
// option. body contains the bytes between `IAC SB <OptionCode>` and `IAC
// SE`.
func (e *TerminalTypeHandler) HandleSB(c *telnet.Connection, body []byte) {

}
