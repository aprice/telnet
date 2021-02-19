package options

import "github.com/aprice/telnet"

// SUPPRESS-GO-AHEAD Telnet Option - https://tools.ietf.org/html/rfc857

// SuppressGoAheadOption will enable GO-AHEAD suppression negotiation on a Server.
func SuppressGoAheadOption(c *telnet.Connection) telnet.Negotiator {
	return &SuppressGoAheadHandler{client: false}
}

// SuppressGoAheadHandler negotiates ECHO for a specific connection.
type SuppressGoAheadHandler struct {
	client bool
}

// OptionCode returns with the code used to negotiate ECHO modes.
func (e *SuppressGoAheadHandler) OptionCode() byte {
	return telnet.TeloptSGA
}

// Offer is called when a new connection is initiated. It offers the handler
// an opportunity to advertise or request an option.
func (e *SuppressGoAheadHandler) Offer(c *telnet.Connection) {
	if !e.client {
		c.Conn.Write([]byte{telnet.IAC, telnet.WILL, e.OptionCode()})
	}
}

// HandleDo is called when an IAC DO command is received for this option,
// indicating the client is requesting the option to be enabled.
func (e *SuppressGoAheadHandler) HandleDo(c *telnet.Connection) {

}

// HandleWill is called when an IAC WILL command is received for this
// option, indicating the client is willing to enable this option.
func (e *SuppressGoAheadHandler) HandleWill(c *telnet.Connection) {

}

// HandleSB is called when a subnegotiation command is received for this
// option. body contains the bytes between `IAC SB <OptionCode>` and `IAC
// SE`.
func (e *SuppressGoAheadHandler) HandleSB(c *telnet.Connection, body []byte) {

}
