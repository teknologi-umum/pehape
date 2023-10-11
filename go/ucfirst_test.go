package pehape_test

import (
	"testing"

	PHP "github.com/teknologi-umum/pehape/go"
)

func TestUcfirst(t *testing.T) {
	type test struct {
		param  string
		expect string
	}

	testCases := []test{
		{
			param:  "Hello",
			expect: "Hello",
		},
		{
			param:  "HELLO",
			expect: "HELLO",
		},
		{
			param:  "GoLang",
			expect: "GoLang",
		},
		{
			param:  "golang",
			expect: "Golang",
		},
		{
			param:  "gOLANG",
			expect: "GOLANG",
		},
		{
			param:  "g",
			expect: "G",
		},
		{
			param:  "!@@#4#Golang",
			expect: "!@@#4#Golang",
		},
		{
			param:  "每一次搜索如何在",
			expect: "每一次搜索如何在",
		},
		{
			param:  "🙂🤣",
			expect: "🙂🤣",
		},
	}

	for _, test := range testCases {
		if res := PHP.Ucfirst(test.param); res != test.expect {
			t.Errorf("expect %s, but got %s", test.expect, res)
		}
	}

	t.Run("It should return empty string if the given parameter is empty string", func(t *testing.T) {
		test := test{
			param:  "",
			expect: "",
		}

		if res := PHP.Ucfirst(test.param); res != test.expect {
			t.Errorf("expect empty string, but got %s", res)
		}
	})
}
