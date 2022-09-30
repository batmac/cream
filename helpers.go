package cream

import "crypto/rand"

// NewRand returns a random byte slice of the given length.
func NewRand(length int) []byte {
	key := make([]byte, length)
	if _, err := rand.Read(key); err != nil {
		panic(err)
	}
	return key
}

// NewIV returns a random byte slice of 16 bytes.
func NewIV() []byte {
	return NewRand(16)
}

// NewKey128 returns a random byte slice of 16 bytes (for AES-128).
func NewKey128() []byte {
	return NewRand(16)
}

// NewKey192 returns a random byte slice of 24 bytes (for AES-192).
func NewKey192() []byte {
	return NewRand(24)
}

// NewKey256 returns a random byte slice of 32 bytes (for AES-256).
func NewKey256() []byte {
	return NewRand(32)
}
