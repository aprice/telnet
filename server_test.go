package telnet_test

import (
	"sync"
	"testing"
	"time"

	"github.com/aprice/telnet"
)

func TestServer_ListenAndServe(t *testing.T) {
	wg := new(sync.WaitGroup)
	s := telnet.NewServer("127.0.0.1:0", telnet.HandleFunc(func(c *telnet.Connection) {
		wg.Add(1)
		_, err := c.Write([]byte("Hello!"))
		if err != nil {
			t.Error(err)
		}
		err = c.Close()
		if err != nil {
			t.Error(err)
		}
		wg.Done()
	}))
	wg.Add(1)
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			t.Error(err)
		}
		wg.Done()
	}()
	time.Sleep(time.Millisecond)
	client, err := telnet.Dial(s.Address)
	if err != nil {
		t.Error(err)
	}
	err = client.Close()
	if err != nil {
		t.Error(err)
	}
	s.Stop()
	wg.Wait()
}
