// Package cream provides a simple interface to encrypt and decrypt streams
// using AES-CTR from the standard library.
// It does not provide any authentication or integrity checks, nor does it
// manage IVs.
// This package is intended to be used as a building block for other
// implementations.
package cream

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io"
)

// NewWriter returns a new cipher.StreamWriter that encrypts the data written
func NewWriter(key, iv []byte, w io.Writer) (*cipher.StreamWriter, error) {
	stream, err := newStream(key, iv)
	if err != nil {
		return nil, err
	}
	return &cipher.StreamWriter{S: stream, W: w}, nil
}

// NewReader returns a new cipher.StreamReader that decrypts the data read
func NewReader(key, iv []byte, r io.Reader) (*cipher.StreamReader, error) {
	stream, err := newStream(key, iv)
	if err != nil {
		return nil, err
	}
	return &cipher.StreamReader{S: stream, R: r}, nil
}

func newStream(key, iv []byte) (cipher.Stream, error) {
	// build aes from key
	switch len(key) {
	case 16, 24, 32:
	default: // invalid key length
		return nil, fmt.Errorf("invalid key length: %d, should be 16, 24 or 32",
			len(key))
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(iv) != block.BlockSize() {
		return nil, fmt.Errorf("invalid IV length: %d, should be %d",
			len(iv), block.BlockSize())
	}
	// create the CTR stream
	stream := cipher.NewCTR(block, iv)
	return stream, nil
}
