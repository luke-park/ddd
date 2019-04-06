package bundling

import (
	"bytes"
	"strings"
	"testing"
)

func TestSymmetricEncryptionBundler(t *testing.T) {
	testCases := [][]byte{
		[]byte{},
		[]byte("H"),
		[]byte("Hello, World!"),
		[]byte(strings.Repeat("a", 66000)),
	}

	for i, v := range testCases {
		bundler := NewSymmetricEncryptionBundler([]byte("1234567812345678"))
		buf := &bytes.Buffer{}

		err := bundler.Bundle(v, buf)
		if err != nil {
			t.Error(err)
			return
		}

		got, err := bundler.Unbundle(buf)
		if err != nil {
			t.Error(err)
			return
		}

		if bytes.Compare(v, got) != 0 {
			t.Errorf("len(v): %v, len(got): %v (test case %v)", len(v), len(got), i)
		}
	}
}
