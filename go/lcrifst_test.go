package pehape_test

import (
	"testing"

	PHP "github.com/teknologi-umum/pehape/go"
)

func TestLcfirst(t *testing.T) {
	type test struct {
		param  string
		expect string
	}

	t.Run("It should convert the given string correct", func(t *testing.T) {
		tests := []test{
			{
				param:  "Hello",
				expect: "hello",
			},
			{
				param:  "GoLang",
				expect: "goLang",
			},
			{
				param:  "golang",
				expect: "golang",
			},
			{
				param:  "G",
				expect: "g",
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

		for _, test := range tests {
			if res := PHP.Lcfirst(test.param); res != test.expect {
				t.Errorf("expect %s, but got %s", test.expect, res)
			}
		}
	})

	t.Run("It should return empty string if the given parameter is empty string", func(t *testing.T) {
		test := &test{
			param:  "",
			expect: "",
		}

		if res := PHP.Lcfirst(test.param); res != test.expect {
			t.Errorf("expect empty string, but got %s", res)
		}
	})
}
