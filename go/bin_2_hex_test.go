package pehape

import (
	"testing"
)

func TestBin2Hex(t *testing.T) {
	type test struct {
		param  string
		expect string
	}
	t.Run("It should encode string into correct value", func(t *testing.T) {
		tt := &test{
			param:  "Hello World!!",
			expect: "48656c6c6f20576f726c642121",
		}

		if res := Bin2Hex(tt.param); res != tt.expect {
			t.Errorf("expected %s, but got %s", tt.expect, res)
		}
	})

	t.Run("It should encode string into correct value", func(t *testing.T) {
		tt := &test{
			param:  "hacktoberfest1234!@#",
			expect: "6861636b746f6265726665737431323334214023",
		}

		if res := Bin2Hex(tt.param); res != tt.expect {
			t.Errorf("expected %s, but got %s", tt.expect, res)
		}
	})
}
