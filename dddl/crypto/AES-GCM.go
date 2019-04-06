package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
)

// AESGCMEncrypt encrypts the specified plaintext with the given key.
func AESGCMEncrypt(plaintext, key []byte) ([]byte, error) {
	nonce := make([]byte, 12)
	_, err := rand.Read(nonce)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	cipher, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	ciphertext := cipher.Seal(nil, nonce, plaintext, nil)
	ciphertextAndNonce := make([]byte, 0)

	ciphertextAndNonce = append(ciphertextAndNonce, nonce...)
	ciphertextAndNonce = append(ciphertextAndNonce, ciphertext...)

	return ciphertextAndNonce, nil
}

// AESGCMDecrypt decrypts the specified ciphertext and nonce with the given key.
func AESGCMDecrypt(ciphertextAndNonce, key []byte) ([]byte, error) {
	nonce := ciphertextAndNonce[:12]
	ciphertext := ciphertextAndNonce[12:]

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	cipher, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	plaintext, err := cipher.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
