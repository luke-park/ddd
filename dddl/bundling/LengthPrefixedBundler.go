package bundling

import (
	"errors"
	"io"
)

// LengthPrefixedBundler is the default message bundler.  It simply length
// prefixes messages with a 4-byte unsigned integer.
type LengthPrefixedBundler struct{}

// ErrInvalidLengthPrefix is returned when a msg to be unbundle has an invalid
// length prefix.
var ErrInvalidLengthPrefix = errors.New("the provided msg cannot be unbundled with LengthPrefixedBundler")

// NewLengthPrefixedBundler creates a new length prefixed bundler.
func NewLengthPrefixedBundler() *LengthPrefixedBundler {
	return &LengthPrefixedBundler{}
}

// Bundle adds the length-prefixing.
func (b *LengthPrefixedBundler) Bundle(msg []byte, w io.Writer) error {
	msgLen := uint(len(msg))

	rawLen := make([]byte, 4)
	rawLen[0] = byte(msgLen >> 24)
	rawLen[1] = byte(msgLen << 8 >> 24)
	rawLen[2] = byte(msgLen << 16 >> 24)
	rawLen[3] = byte(msgLen << 24 >> 24)

	_, err := w.Write(append(rawLen, msg...))
	return err
}

// Unbundle removes the length-prefixing.
func (b *LengthPrefixedBundler) Unbundle(r io.Reader) ([]byte, error) {
	rawLen := make([]byte, 4)
	_, err := io.ReadFull(r, rawLen)
	if err != nil {
		return nil, err
	}

	msgLen := 0
	msgLen += int(rawLen[3])
	msgLen += int(rawLen[2]) << 8
	msgLen += int(rawLen[1]) << 16
	msgLen += int(rawLen[0]) << 24

	raw := make([]byte, msgLen)
	_, err = io.ReadFull(r, raw)
	if err != nil {
		return nil, err
	}

	return raw, nil
}
