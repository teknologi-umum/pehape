package pehape_test

import (
	"errors"
	"testing"

	PHP "github.com/teknologi-umum/pehape/go"
)

func TestStrchr(t *testing.T) {
	type param struct {
		str          string
		search       interface{}
		beforeSearch bool
	}
	type test struct {
		param  param
		expect string
	}

	t.Run("It should return string correctly if the given parameters is valid", func(t *testing.T) {
		tests := []test{
			{
				param: param{
					str:          "Hello World",
					search:       "o",
					beforeSearch: false,
				},
				expect: "o World",
			},
			{
				param: param{
					str:          "Hello World",
					search:       "Wor",
					beforeSearch: false,
				},
				expect: "World",
			},
			{
				param: param{
					str:          "Hello World",
					search:       111,
					beforeSearch: false,
				},
				expect: "o World",
			},
			{
				param: param{
					str:          "Hello World",
					search:       "\x6F",
					beforeSearch: true,
				},
				expect: "Hell",
			},
			{
				param: param{
					str:          "Hello World",
					search:       "o",
					beforeSearch: true,
				},
				expect: "Hell",
			},
		}

		for _, test := range tests {
			res, err := PHP.Strchr(test.param.str, test.param.search, test.param.beforeSearch)

			if err != nil {
				t.Errorf("expect error nil, but got %s", err)
			}
			if res != test.expect {
				t.Errorf("\nParameters:\n\tstr: %s\n\tsearch: %v\n\tbeforeSearch:%t\nExpect %s, but got %s",
					test.param.str, test.param.search, test.param.beforeSearch, test.expect, res,
				)
			}
		}
	})

	t.Run("It should return error correctly if the given parameters is invalid", func(t *testing.T) {
		tests := []test{
			{
				param: param{
					str:          "Hello World",
					search:       false,
					beforeSearch: true,
				},
			},
			{
				param: param{
					str:          "Hello World",
					search:       []string{"o"},
					beforeSearch: false,
				},
			},
		}

		for _, test := range tests {
			_, err := PHP.Strchr(test.param.str, test.param.search, test.param.beforeSearch)

			if err == nil {
				t.Errorf("expect error, but got nil")
			}
			if !errors.Is(err, PHP.ErrStrchrInvalidParameterType) {
				t.Errorf("expect error type %s, but got %s", PHP.ErrStrchrInvalidParameterType, err)
			}
		}
	})

	t.Run("It should return error correctly if string not found", func(t *testing.T) {
		tests := []test{
			{
				param: param{
					str:          "Hello World",
					search:       "O",
					beforeSearch: false,
				},
			},
			{
				param: param{
					str:          "Hello World",
					search:       "Wol",
					beforeSearch: true,
				},
			},
		}

		for _, test := range tests {
			_, err := PHP.Strchr(test.param.str, test.param.search, test.param.beforeSearch)

			if err == nil {
				t.Errorf("expect error, but got nil")
			}
			if !errors.Is(err, PHP.ErrStrchrStringNotFound) {
				t.Errorf("expect error type %s, but got %s", PHP.ErrStrchrStringNotFound, err)
			}
		}
	})

	t.Run("It should return string correctly if the given beforeSearch parameter is empty", func(t *testing.T) {
		test := &test{
			param: param{
				str:    "Hello World",
				search: "o",
			},
			expect: "o World",
		}

		res, err := PHP.Strchr(test.param.str, test.param.search)

		if err != nil {
			t.Errorf("expect error nil, but got %s", err)
		}
		if res != test.expect {
			t.Errorf("\nParameters:\n\tstr: %s\n\tsearch: %v\n\tbeforeSearch:%t\nExpect %s, but got %s",
				test.param.str, test.param.search, test.param.beforeSearch, test.expect, res,
			)
		}
	})
}
