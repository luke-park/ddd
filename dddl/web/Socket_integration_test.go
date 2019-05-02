package web

import (
	"testing"
	"dddl/bundling"
	"net"
)

func TestShouldCorrectlyCommunicate(t *testing.T) {
	l, err := net.Listen("tcp", ":555")
	if err != nil {
		t.Error(err)
		return
	}

	go func () {
		c, err := l.Accept()
		if err != nil {
			t.Error(err)
			return
		}

		s := NewSocket(c, bundling.NewLengthPrefixedBundler())
		msg, err := s.Receive()
		if err != nil {
			t.Error(err)
			return
		}

		if len(msg) != 2 || msg[0] != 26 || msg[1] != 254 {
			t.Errorf("got: %v, expected: [ 26, 254 ]", msg)
			return
		}

		s.Send([]byte { 250, 249, 248 })
		s.Close()
	}()

	c, err := net.Dial("tcp", "127.0.0.1:555")
	if err != nil {
		t.Error(err)
		return
	}

	s := NewSocket(c, bundling.NewLengthPrefixedBundler())
	s.Send([]byte { 26, 254 })
	
	msg, err := s.Receive()
	if err != nil {
		t.Error(err)
		return
	}

	if len(msg) != 3 || msg[0] != 250 || msg[1] != 249 || msg[2] != 248 {
		t.Errorf("got: %v, expected: [ 250, 249, 248 ]", msg)
		return
	}

	s.Close()
}