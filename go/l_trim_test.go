package pehape_test

import (
	"testing"

	PHP "github.com/teknologi-umum/pehape/go"
)

func TestLTrim(t *testing.T) {
	type param struct {
		str       string
		charLists []string
	}
	type test struct {
		param  param
		expect string
	}

	t.Run("It should remove predefined characters at the left side of the string if the given charlist is empty", func(t *testing.T) {
		tt := &test{
			param: param{
				str: "	\n\x0BHello World",
			},
			expect: "Hello World",
		}

		if res := PHP.LTrim(tt.param.str); res != tt.expect {
			t.Errorf("expected %s, but got %s", tt.expect, res)
		}
	})

	t.Run("It should remove given characters at the left side of the string if the given charlist is not empty", func(t *testing.T) {
		tt := &test{
			param: param{
				str:       "123456\t    \n\x0BHello World",
				charLists: []string{"123456", " ", "\t\n\x0B"},
			},
			expect: "Hello World",
		}

		if res := PHP.LTrim(tt.param.str, tt.param.charLists...); res != tt.expect {
			t.Errorf("expected %s, but got %s", tt.expect, res)
		}
	})
}
