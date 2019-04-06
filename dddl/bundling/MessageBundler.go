package bundling

import "io"

// MessageBundler defines the methods that a message bundler must implement.
type MessageBundler interface {
	Bundle(raw []byte, w io.Writer) error
	Unbundle(r io.Reader) ([]byte, error)
}
