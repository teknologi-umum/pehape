package pehape_test

import (
	"testing"

	PHP "github.com/teknologi-umum/pehape/go"
)

func TestMd5(t *testing.T) {
	type param struct {
		str string
		raw bool
	}
	type test struct {
		param  param
		expect string
	}

	t.Run("It should calculate string correctly if the given raw is not empty", func(t *testing.T) {
		tests := []test{
			{
				param: param{
					str: "Hello",
					raw: false,
				},
				expect: "8b1a9953c4611296a827abf8c47804d7",
			},
			{
				param: param{
					str: "Hello World!12345678",
					raw: true,
				},
				expect: string([]byte{216, 70, 166, 13, 80, 18, 226, 71, 163, 33, 175, 141, 94, 206, 62, 134}),
			},
		}

		for _, tt := range tests {
			if res := PHP.Md5(tt.param.str, tt.param.raw); res != tt.expect {
				t.Errorf("\nParameters:\n\tstr: %s\n\traw:%t\nexpect %s, but got %s", tt.param.str, tt.param.raw, tt.expect, res)
			}
		}
	})

	t.Run("It should calculate string correctly if the given raw is emtpy", func(t *testing.T) {
		tt := &test{
			param: param{
				str: "Hello",
			},
			expect: "8b1a9953c4611296a827abf8c47804d7",
		}

		if res := PHP.Md5(tt.param.str, tt.param.raw); res != tt.expect {
			t.Errorf("\nParameters:\n\tstr: %s\n\traw:%t\nexpect %s, but got %s", tt.param.str, tt.param.raw, tt.expect, res)
		}
	})
}
