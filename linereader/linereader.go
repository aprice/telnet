package linereader

import (
	"bufio"
	"io"
)

// LineReader reads lines from an io.Reader and passes them to a channel.
type LineReader struct {
	C chan []byte
}

// New constructs a new LineReader. While the zero value is usable, it may
// cause a race condition; it is better to call New() to initialize the channel
// before calling ReadLines, to ensure that the channel is not nil when it is
// first read.
func New() *LineReader {
	return &LineReader{C: make(chan []byte)}
}

// ReadLines continuously reads lines from the Reader and sends them on C. It
// will not return until it encounters an error (including io.EOF).
func (r *LineReader) ReadLines(in io.Reader) (err error) {
	if r.C == nil {
		r.C = make(chan []byte)
	}
	defer close(r.C)
	rdr := bufio.NewReader(in)
	var line, cont []byte
	var prefix bool
	for {
		line, prefix, err = rdr.ReadLine()
		for prefix && err == nil {
			cont, prefix, err = rdr.ReadLine()
			line = append(line, cont...)
		}
		if line != nil {
			r.C <- line
		}
		if err == io.EOF {
			err = nil
			break
		}
		if err != nil {
			break
		}
	}
	return
}
