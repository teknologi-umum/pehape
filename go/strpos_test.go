package pehape_test

import (
	"errors"
	"testing"

	PHP "github.com/teknologi-umum/pehape/go"
)

func TestStrpos(t *testing.T) {
	type param struct {
		str    string
		find   string
		offset int
	}
	type test struct {
		param  param
		expect int
	}

	t.Run("It should return correct error if the given offset is invalid", func(t *testing.T) {
		tests := []test{
			{
				param: param{
					str:    "abcdef",
					find:   "c",
					offset: 6,
				},
			},
			{
				param: param{
					str:    "abcdef",
					find:   "c",
					offset: -7,
				},
			},
		}

		for _, test := range tests {
			_, err := PHP.Strpos(test.param.str, test.param.find, test.param.offset)

			if err == nil {
				t.Errorf("expect error invalid offset, but got nil")
			}
			if !errors.Is(err, PHP.ErrStrposInvalidOffset) {
				t.Errorf("expect error %s, but got %s", PHP.ErrStrposInvalidOffset, err)
			}
		}
	})

	t.Run("It should return correct position of the given string if the offset is empty", func(t *testing.T) {
		test := &test{
			param: param{
				str:  "I love GoLang, I love GoLang too!",
				find: "Go",
			},
			expect: 7,
		}

		res, err := PHP.Strpos(test.param.str, test.param.find)

		if err != nil {
			t.Errorf("expect error nil, but got %s", err)
		}
		if res != test.expect {
			t.Errorf("expect %d, but got %d", test.expect, res)
		}
	})

	t.Run("It should return correct position of the given string if the offset is not empty", func(t *testing.T) {
		tests := []test{
			{
				param: param{
					str:    "I love GoLang, I love GoLang too!",
					find:   "Go",
					offset: 3,
				},
				expect: 7,
			},
			{
				param: param{
					str:    "I love GoLang, I love GoLang too!",
					find:   "Go",
					offset: 9,
				},
				expect: 22,
			},
			{
				param: param{
					str:    "I love GoLang, I love GoLang too!",
					find:   "Go",
					offset: -11,
				},
				expect: 22,
			},
			{
				param: param{
					str:    "I love GoLang, I love GoLang too!",
					find:   "Go",
					offset: -27,
				},
				expect: 7,
			},
			{
				param: param{
					str:    "hello",
					find:   "o",
					offset: -1,
				},
				expect: 4,
			},
			{
				param: param{
					str:    "hello",
					find:   "o",
					offset: 4,
				},
				expect: 4,
			},
		}

		for _, test := range tests {
			res, err := PHP.Strpos(test.param.str, test.param.find, test.param.offset)

			if err != nil {
				t.Errorf("expect error nil, but got %s", err)
			}
			if res != test.expect {
				t.Errorf("expect %d, but got %d", test.expect, res)
			}
		}
	})

	t.Run("It should return correct error if string not found", func(t *testing.T) {
		tests := []test{
			{
				param: param{
					str:    "I love GoLang, I love GoLang too!",
					find:   "hahahah",
					offset: 0,
				},
			},
			{
				param: param{
					str:    "I love GoLang, I love GoLang too!",
					find:   "Go",
					offset: 23,
				},
			},
			{
				param: param{
					str:    "I love GoLang, I love GoLang too!",
					find:   "Go",
					offset: -10,
				},
			},
		}

		for _, test := range tests {
			_, err := PHP.Strpos(test.param.str, test.param.find, test.param.offset)

			if err == nil {
				t.Errorf("expect error, but got nil")
			}
			if !errors.Is(err, PHP.ErrStrposStringNotFound) {
				t.Errorf("expect error %s, but got %s", PHP.ErrStrposStringNotFound, err)
			}
		}
	})

}
