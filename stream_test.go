package cream_test

import (
	"io"
	"log"
	"testing"

	"github.com/batmac/cream"
)

func TestNew(t *testing.T) {
	testCases := []struct {
		name        string
		key         []byte
		iv          []byte
		expectedErr bool
	}{
		{"128bit", []byte("1234567890123456"), []byte("1234567890123456"), false},
		{"192bit", []byte("123456789012345678901234"), []byte("1234567890123456"), false},
		{"256bit", []byte("12345678901234567890123456789012"), []byte("1234567890123456"), false},
		{"invalid key", []byte("123456789012345678901234567890123"), []byte("1234567890123456"), true},
		{"invalid IV", []byte("1234567890123456"), []byte("12345678901234567890123456789012"), true},
	}

	plaintext := []byte("hello world")

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			r, w := io.Pipe()
			defer func() {
				_ = r.Close()
				_ = w.Close()
			}()
			sr, err := cream.NewReader(tt.key, tt.iv, r)
			if tt.expectedErr {
				if err == nil {
					t.Fatal(err)
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}

			go func() {
				sw, err := cream.NewWriter(tt.key, tt.iv, w)
				if err != nil {
					log.Fatal(err)
				}
				_, err = sw.Write(plaintext)
				if err != nil {
					log.Fatal(err)
				}
				sw.Close()
			}()

			str, err := io.ReadAll(sr)
			if err != nil {
				t.Fatal(err)
			}
			if string(str) != string(plaintext) {
				t.Fatal("not equal")
			}
		})
	}
}
