package bundling

import (
	"bytes"
	"strings"
	"testing"
)

func TestLengthPrefixedBundler(t *testing.T) {
	testCases := []struct {
		msg                 string
		expectedFirst4Bytes []byte
	}{
		{msg: "", expectedFirst4Bytes: []byte{0, 0, 0, 0}},
		{msg: "H", expectedFirst4Bytes: []byte{0, 0, 0, 1}},
		{msg: "Hello, World!", expectedFirst4Bytes: []byte{0, 0, 0, 13}},
		{msg: strings.Repeat("a", 255), expectedFirst4Bytes: []byte{0, 0, 0, 255}},
		{msg: strings.Repeat("a", 256), expectedFirst4Bytes: []byte{0, 0, 1, 0}},
		{msg: strings.Repeat("a", 257), expectedFirst4Bytes: []byte{0, 0, 1, 1}},
		{msg: strings.Repeat("a", 65535), expectedFirst4Bytes: []byte{0, 0, 255, 255}},
		{msg: strings.Repeat("a", 65536), expectedFirst4Bytes: []byte{0, 1, 0, 0}},
		{msg: strings.Repeat("a", 65537), expectedFirst4Bytes: []byte{0, 1, 0, 1}},
		{msg: strings.Repeat("a", 16777215), expectedFirst4Bytes: []byte{0, 255, 255, 255}},
		{msg: strings.Repeat("a", 16777216), expectedFirst4Bytes: []byte{1, 0, 0, 0}},
		{msg: strings.Repeat("a", 16777217), expectedFirst4Bytes: []byte{1, 0, 0, 1}},
	}

	for i, v := range testCases {
		bundler := NewLengthPrefixedBundler()
		buf := &bytes.Buffer{}

		err := bundler.Bundle([]byte(v.msg), buf)
		if err != nil {
			t.Error(err)
			return
		}

		got := buf.Bytes()
		if len(got) < 4 || len(v.expectedFirst4Bytes) != 4 {
			t.Errorf("len(got): %v, len(expectedFirst4Bytes): %v", len(got), len(v.expectedFirst4Bytes))
			return
		}

		if bytes.Compare(got[:4], v.expectedFirst4Bytes) != 0 {
			t.Errorf("%v != %v (test case index %v)", got[:4], v.expectedFirst4Bytes, i)
			return
		}

		unbundleGot, err := bundler.Unbundle(buf)
		if err != nil {
			t.Error(err)
			return
		}

		if string(unbundleGot) != v.msg {
			t.Errorf("len(string(unbundleGot)): %v, len(v.msg): %v", len(unbundleGot), len(v.msg))
			return
		}
	}
}
