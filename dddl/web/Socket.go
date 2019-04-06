package web

import (
	"dddl/bundling"
	"net"
)

// Socket is a wrapper around net.Conn that simplifies message sending and
// receiving.
type Socket struct {
	conn    net.Conn
	bundler bundling.MessageBundler
}

// NewSocket creates a new socket with the provided connection.
func NewSocket(conn net.Conn, bundler bundling.MessageBundler) *Socket {
	return &Socket{
		conn:    conn,
		bundler: bundler,
	}
}

// Send sends some data to the remote socket.
func (s *Socket) Send(data []byte) error {
	return s.bundler.Bundle(data, s.conn)
}

// SendError sends an error to the remote connection.
func (s *Socket) SendError(err error) error {
	return s.Send([]byte(err.Error()))
}

// Receive receives some data from the remote socket.
func (s *Socket) Receive() ([]byte, error) {
	return s.bundler.Unbundle(s.conn)
}

// Close closes the socket.
func (s *Socket) Close() {
	s.conn.Close()
}
