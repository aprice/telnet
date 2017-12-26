/*
The telnet package provides basic telnet client and server implementations. It
includes handling of IACs and extensible telnet option negotiation for both
clients and servers.

Running a server:

	svr := telnet.NewServer(":9999", telnet.HandleFunc(func(c *telnet.Connection){
		log.Printf("Connection received: %s", c.RemoteAddr())
		c.Write([]byte("Hello world!\r\n"))
		c.Close()
	}))
	svr.ListenAndServe()

The server API is modeled after the net/http API, so it should be easy to get
your bearings; of course, telnet and HTTP are very different beasts, so the
similarities are somewhat limited. The server listens on a TCP address for new
connections. Whenever a new connection is received, the connection handler is
called with the connection object. This object is a wrapper for the underlying
TCP connection, which aims to transparently handle IAC.

Running a client is pretty simple:

	conn, err := telnet.Dial("127.0.0.1:9999")

This is really straightforward - dial out, get a telnet connection handler back.
Again, this handles IAC transparently, and like the Server, can take a list of
optional IAC handlers. Bear in mind that some handlers - for example, the
included NAWS handler - use different Option functions to register them with a
client versus a server; this is because they may behave differently at each end.
See the documentation for the options for more details.

*/
package telnet
