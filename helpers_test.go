package cream_test

import (
	"encoding/hex"
	"testing"

	"github.com/batmac/cream"
)

func Test_NewRand(t *testing.T) {
	testCases := []struct {
		name        string
		fn          func() []byte
		expectedLen int
	}{
		{"NewIV", cream.NewIV, 16},
		{"NewKey128", cream.NewKey128, 16},
		{"NewKey192", cream.NewKey192, 24},
		{"NewKey256", cream.NewKey256, 32},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			keyFn := tt.fn
			runs := 100
			history := make(map[string]bool)

			for i := 0; i < runs; i++ {
				k := string(keyFn())
				if len(k) != tt.expectedLen {
					t.Fatalf("expected %d, got %d", tt.expectedLen, len(k))
				}
				if _, ok := history[k]; ok {
					t.Fatalf("duplicate key: %s", hex.EncodeToString([]byte(k)))
				}
				history[k] = true

			}
		})
	}
}
