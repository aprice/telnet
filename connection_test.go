package telnet_test

import (
	"bytes"
	"net"
	"testing"

	"github.com/aprice/telnet"
)

func TestConnection_Write(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected []byte
	}{
		{
			name:     "plain",
			input:    []byte("hello world"),
			expected: []byte("hello world"),
		},
		{
			name:     "iac",
			input:    []byte("hello \xffworld"),
			expected: []byte("hello \xff\xffworld"),
		},
		{
			name:     "doubleiac",
			input:    []byte("hello \xff\xffworld"),
			expected: []byte("hello \xff\xff\xff\xffworld"),
		},
	}
	buf := bytes.NewBuffer(nil)
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			buf.Reset()
			client, server := net.Pipe()
			go func() {
				conn := telnet.NewConnection(server, nil)
				_, err := conn.Write(test.input)
				if err != nil {
					t.Error(err)
				}
				conn.Close()
			}()
			buf.ReadFrom(client)
			client.Close()
			if !bytes.Equal(test.expected, buf.Bytes()) {
				t.Errorf("Expected %v, got %v", test.expected, buf.Bytes())
			}
		})
	}
}

func TestConnection_Read(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected []byte
	}{
		{
			name:     "plain",
			input:    []byte("hello world"),
			expected: []byte("hello world"),
		},
		{
			name:     "iac",
			input:    []byte("hello \xff\xffworld"),
			expected: []byte("hello \xffworld"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client, server := net.Pipe()
			go func() {
				client.Write(test.input)
				t.Log("Finished writing")
				client.Close()
				t.Log("Closed client")
			}()
			conn := telnet.NewConnection(server, nil)
			b := make([]byte, len(test.expected))
			_, err := conn.Read(b)
			if err != nil {
				t.Error(err)
			}
			conn.Close()
			if !bytes.Equal(test.expected, b) {
				t.Errorf("Expected %v, got %v", test.expected, b)
			}
		})
	}
}

type closerBuf struct {
	*bytes.Buffer
}

func (c *closerBuf) Close() error { return nil }
