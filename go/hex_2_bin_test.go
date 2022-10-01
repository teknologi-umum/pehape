package pehape_test

import (
	"encoding/hex"
	"errors"
	"testing"

	PHP "github.com/teknologi-umum/pehape/go"
)

func TestHex2Bin(t *testing.T) {
	type test struct {
		param  string
		expect string
	}
	t.Run("It should decode hexadecimal characters into correct ASCII string", func(t *testing.T) {
		tt := &test{
			param:  "48656c6c6f20576f726c642121",
			expect: "Hello World!!",
		}
		if res, _ := PHP.Hex2Bin(tt.param); res != tt.expect {
			t.Errorf("expected %s, but got %s", tt.expect, res)
		}
	})

	t.Run("It should return error if given parameter is not even length", func(t *testing.T) {
		tt := &test{
			param: "4",
		}
		_, err := PHP.Hex2Bin(tt.param)
		if err == nil {
			t.Errorf("expected error, but got nil")
		}
		if !errors.Is(err, hex.ErrLength) {
			t.Errorf("expected %s, but got %s", hex.ErrLength, err)
		}
	})

	t.Run("It should return error if given parameter is not hexadecimal characters", func(t *testing.T) {
		tt := &test{
			param: "Incorrect params",
		}
		_, err := PHP.Hex2Bin(tt.param)
		if err == nil {
			t.Errorf("expected error, but got nil")
		}
	})
}
