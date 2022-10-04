package pehape_test

import (
	"errors"
	"testing"

	PHP "github.com/teknologi-umum/pehape/go"
)

func TestTrim(t *testing.T) {
	type param struct {
		str   string
		chars []string
	}
	type test struct {
		param  param
		expect string
	}

	t.Run("It should return an error if the given range is invalid", func(t *testing.T) {
		test := &test{
			param: param{
				str:   "abcdef",
				chars: []string{"z..a"},
			},
		}

		_, err := PHP.Trim(test.param.str, test.param.chars...)

		if err == nil {
			t.Errorf("expect error %s, but got nil", PHP.ErrTrimInvalidRange)
		}
		if !errors.Is(err, PHP.ErrTrimInvalidRange) {
			t.Errorf("expect error %s, but got %s", PHP.ErrTrimInvalidRange, err)
		}
	})

	t.Run("It should trim the given string correctly", func(t *testing.T) {
		tests := []test{
			{
				param: param{
					str:   "abcdef!@#pqrstu",
					chars: []string{"a..z"},
				},
				expect: "!@#",
			},
			{
				param: param{
					str:   "221Hello World1!23\n",
					chars: []string{"23", "\n"},
				},
				expect: "1Hello World1!",
			},
			{
				param: param{
					str:   "\t\n Hello World\r\x0B",
					chars: []string{"\t\x0B", " ", "\n\r"},
				},
				expect: "Hello World",
			},
			{
				param: param{
					str: "\t\n Hello World\r\x0B",
				},
				expect: "Hello World",
			},
			{
				param: param{
					str:   "Hello World!!!",
					chars: []string{"\x21"},
				},
				expect: "Hello World",
			},
			{
				param: param{
					str:   "hELLOhaihai wORLD",
					chars: []string{"h", "A..Z", "ai", " "},
				},
				expect: "w",
			},
		}

		for _, test := range tests {
			res, err := PHP.Trim(test.param.str, test.param.chars...)

			if err != nil {
				t.Errorf("expect error nil, but got %s", err)
			}
			if res != test.expect {
				t.Errorf("expect %s, but got %s", test.expect, res)
			}
		}
	})
}
