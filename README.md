# Package telnet

The [`telnet` package](http://godoc.org/github.com/aprice/telnet) provides basic
telnet client and server implementations for Go, including handling of IACs and
extensible telnet option negotiation.

Currently both the basic client and server are implemented, as well as a NAWS
client/server handler as a working example. Additional handlers may be added to
the core library over time (feel free to submit a PR if you've written one you'd
like to see added!)

## Usage

Running a server:
```go
svr := telnet.NewServer(":9999", telnet.HandleFunc(func(c *telnet.Connection){
	log.Printf("Connection received: %s", c.RemoteAddr())
	c.Write([]byte("Hello world!\r\n"))
	c.Close()
}))
svr.ListenAndServe()
```

The server API is modeled after the `net/http` API, so it should be easy to get
your bearings; of course, telnet and HTTP are very different beasts, so the
similarities are somewhat limited. The server listens on a TCP address for new
connections. Whenever a new connection is received, the connection handler is
called with the connection object. This object is a wrapper for the underlying
TCP connection, which aims to transparently handle IAC. There is a slightly
more complex example located in the `example` package.

Running a client is pretty simple:
```go
conn, err := telnet.Dial("127.0.0.1:9999")
```

This is really straightforward - dial out, get a telnet connection handler back.
Again, this handles IAC transparently, and like the Server, can take a list of
optional IAC handlers. Bear in mind that some handlers - for example, the
included NAWS handler - use different Option functions to register them with a
client versus a server; this is because they may behave differently at each end.
See the documentation for the options for more details.

## Linereader

A sub-package, `linereader`, exposes a simple reader intended to be run in a
Goroutine, which consumes lines from an `io.Reader` and sends them over a
channel for asynchronous handling.
