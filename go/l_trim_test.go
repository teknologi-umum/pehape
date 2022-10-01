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
		res, err := PHP.LTrim(tt.param.str)
		if err != nil {
			t.Errorf("expected error nil, but got %s", err)
		}
		if res != tt.expect {
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
		res, err := PHP.LTrim(tt.param.str, tt.param.charLists...)
		if err != nil {
			t.Errorf("expected error nil, but got %s", err)
		}
		if res != tt.expect {
			t.Errorf("expected %s, but got %s", tt.expect, res)
		}
	})

	t.Run("It should remove the characters in charlists range", func(t *testing.T) {
		tt := &test{
			param: param{
				str:       "defghijklmnopqrstuvwxy.abc",
				charLists: []string{"a..z"},
			},
			expect: ".abc",
		}
		res, err := PHP.LTrim(tt.param.str, tt.param.charLists...)
		if err != nil {
			t.Errorf("expected error nil, but got %s", err)
		}
		if res != tt.expect {
			t.Errorf("expected %s, but got %s", tt.expect, res)
		}
	})

	t.Run("It should return an error if the given range pattern is invalid", func(t *testing.T) {
		tt := &test{
			param: param{
				str:       "abc.defghijklmnopqrstuvwxy",
				charLists: []string{"z..a"},
			},
		}
		_, err := PHP.LTrim(tt.param.str, tt.param.charLists...)
		if err == nil {
			t.Errorf("expected error but got nil")
		}
	})

	t.Run("It should remove characters correctly form the given pattern", func(t *testing.T) {
		tt := &test{
			param: param{
				str:       "abc.defghijklmnopqrstuvwxyzHELLO WORLD !",
				charLists: []string{"a..z\x20", "\x21..R"},
			},
			expect: "WORLD !",
		}
		res, err := PHP.LTrim(tt.param.str, tt.param.charLists...)
		if err != nil {
			t.Errorf("expected error nil but got %s", err)
		}
		if res != tt.expect {
			t.Errorf("expected %s, but got %s", tt.expect, res)
		}
	})
}
