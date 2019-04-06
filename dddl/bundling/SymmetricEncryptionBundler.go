package bundling

import (
	"io"

	"dddl/crypto"
)

// SymmetricEncryptionBundler encrypts messages using AES-GCM.  It feeds the
// result through the default length prefix bundler before returning.
type SymmetricEncryptionBundler struct {
	key                      []byte
	underlyingLengthPrefixer *LengthPrefixedBundler
}

// NewSymmetricEncryptionBundler creates a new symmetric encryption bundler.
func NewSymmetricEncryptionBundler(key []byte) *SymmetricEncryptionBundler {
	return &SymmetricEncryptionBundler{
		key:                      key,
		underlyingLengthPrefixer: NewLengthPrefixedBundler(),
	}
}

// Bundle adds the length-prefixing.
func (b *SymmetricEncryptionBundler) Bundle(msg []byte, w io.Writer) error {
	ct, err := crypto.AESGCMEncrypt(msg, b.key)
	if err != nil {
		return err
	}

	return b.underlyingLengthPrefixer.Bundle(ct, w)
}

// Unbundle removes the length-prefixing.
func (b *SymmetricEncryptionBundler) Unbundle(r io.Reader) ([]byte, error) {
	pt, err := b.underlyingLengthPrefixer.Unbundle(r)
	if err != nil {
		return nil, err
	}

	return crypto.AESGCMDecrypt(pt, b.key)
}
