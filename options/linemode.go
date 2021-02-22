package options

import "github.com/aprice/telnet"

// ECHO Telnet Echo Option - https://tools.ietf.org/html/rfc857

// LinemodeOption enables Linemode negotiation on a Server.
func LinemodeOption(c *telnet.Connection) telnet.Negotiator {
	return &LinemodeHandler{client: false}
}

// LinemodeHandler negotiates ECHO for a specific connection.
type LinemodeHandler struct {
	client bool
}

// OptionCode returns with the code used to negotiate ECHO modes.
func (e *LinemodeHandler) OptionCode() byte {
	return telnet.TeloptLINEMODE
}

// Offer is called when a new connection is initiated. It offers the handler
// an opportunity to advertise or request an option.
func (e *LinemodeHandler) Offer(c *telnet.Connection) {
	if !e.client {
		c.Conn.Write([]byte{telnet.IAC, telnet.WONT, e.OptionCode()})
	}
}

// HandleDo is called when an IAC DO command is received for this option,
// indicating the client is requesting the option to be enabled.
func (e *LinemodeHandler) HandleDo(c *telnet.Connection) {

}

// HandleWill is called when an IAC WILL command is received for this
// option, indicating the client is willing to enable this option.
func (e *LinemodeHandler) HandleWill(c *telnet.Connection) {

}

// HandleSB is called when a subnegotiation command is received for this
// option. body contains the bytes between `IAC SB <OptionCode>` and `IAC
// SE`.
func (e *LinemodeHandler) HandleSB(c *telnet.Connection, body []byte) {

}
