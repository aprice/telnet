package telnet

import (
	"net"
)

// Option functions add handling of a telnet option to a Server. The Option
// function takes a connection (which it can store but needn't) and returns a
// Negotiator; it is up to the Option function whether a single instance of
// the Negotiator is reused or if a new instance is created for each connection.
type Option func(c *Connection) Negotiator

// Handler is a telnet connection handler. The Handler passed to a server will
// be called for all incoming connections.
type Handler interface {
	HandleTelnet(conn *Connection)
}

// HandleFunc makes it easy to pass a function as a Handler instead of a full
// type.
type HandleFunc func(conn *Connection)

// HandleTelnet implements Handler, and simply calls the function.
func (f HandleFunc) HandleTelnet(conn *Connection) {
	f(conn)
}

// Server listens for telnet connections.
type Server struct {
	// Address is the addres the Server listens on.
	Address  string
	handler  Handler
	options  []Option
	listener net.Listener
	quitting bool
}

// NewServer constructs a new telnet server.
func NewServer(addr string, handler Handler, options ...Option) *Server {
	return &Server{Address: addr, handler: handler, options: options}
}

// Serve runs the telnet server. This function does not return and
// should probably be run in a goroutine.
func (s *Server) Serve(l net.Listener) error {
	s.listener = l
	s.Address = s.listener.Addr().String()
	for {
		c, err := s.listener.Accept()
		if err != nil {
			if s.quitting {
				return nil
			}
			return err
		}
		conn := NewConnection(c, s.options)
		go func() {
			s.handler.HandleTelnet(conn)
			conn.Conn.Close()
		}()
	}
}

// ListenAndServe runs the telnet server by creating a new Listener using the
// current Server.Address, and then calling Serve().
func (s *Server) ListenAndServe() error {
	l, err := net.Listen("tcp", s.Address)
	if err != nil {
		return err
	}

	return s.Serve(l)
}

// Stop the telnet server. This stops listening for new connections, but does
// not affect any active connections already opened.
func (s *Server) Stop() {
	if s.quitting {
		return
	}
	s.quitting = true
	s.listener.Close()
}
