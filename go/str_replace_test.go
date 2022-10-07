package pehape_test

import (
	"errors"
	"reflect"
	"testing"

	PHP "github.com/teknologi-umum/pehape/go"
)

func TestStrReplace(t *testing.T) {
	type param struct {
		find, replace, str interface{}
	}
	type test struct {
		param       param
		expect      interface{}
		expectCount int
	}

	t.Run("It should return error if the given parameters is invalid", func(t *testing.T) {
		tests := []test{
			// if 'find' or 'replace' has nested array
			{
				param: param{
					find: []interface{}{
						[]string{"world", "have"},
						"day",
					},
					replace: "people",
					str:     "Hello world, have a nice day!",
				},
			},
			{
				param: param{
					find: []interface{}{
						[]string{"world", "have"},
						"day",
					},
					replace: []string{"we", "eee"},
					str:     "Hello world, have a nice day!",
				},
			},
			{
				param: param{
					find:    []string{"day"},
					replace: []string{"we", "eee"},
					str:     []int{1, 2},
				},
			},
			{
				param: param{
					find: []string{
						"world", "have",
					},
					replace: []interface{}{
						[]string{
							"people", "has",
						},
					},
					str: "Hello world, have a nice day!",
				},
			},
			// if parameter contain non slice and string type
			{
				param: param{
					find:    1,
					replace: "people",
					str:     "Hello world 123",
				},
			},
			{
				param: param{
					find:    "hello",
					replace: 1,
					str:     "Hello world 123",
				},
			},
			{
				param: param{
					find:    []int{1},
					replace: 1,
					str:     false,
				},
			},
			{
				param: param{
					find:    1.90,
					replace: []int{1},
					str:     "Hello world 123",
				},
			},
			{
				param: param{
					find:    "hello",
					replace: false,
					str:     1,
				},
			},
			{
				param: param{
					find:    false,
					replace: "people",
					str:     []int{1},
				},
			},
			{
				param: param{
					find:    []bool{false, true},
					replace: "people",
					str:     []string{"we", "he"},
				},
			},
			{
				param: param{
					find:    "world",
					replace: "people",
					str: []interface{}{
						[]string{"hello", "world"},
						"have a nice day",
					},
				},
			},
			// the type of 'find' must be an array if the type of 'replace' is also an array
			{
				param: param{
					find: "world",
					replace: []string{
						"people",
					},
					str: "Hello world, have a nice day!",
				},
			},
		}

		for _, test := range tests {
			_, _, err := PHP.StrReplace(test.param.find, test.param.replace, test.param.str)

			if err == nil {
				t.Errorf("expect error, but got nil")
			}
			if !errors.Is(err, PHP.ErrStrReplaceInvalidParameter) {
				t.Errorf("expect error %s, but got %s", PHP.ErrStrReplaceInvalidParameter, err)
			}
		}
	})

	t.Run("It should replace the given string correctly", func(t *testing.T) {
		tests := []test{
			{
				param: param{
					find:    []string{"world", "nice"},
					replace: []string{"people", "bad"},
					str:     "Hello world! have a nice day",
				},
				expect:      "Hello people! have a bad day",
				expectCount: 2,
			},
			{
				param: param{
					find:    []string{"world", "nice", "day"},
					replace: []string{"people", "bad"},
					str:     "Hello world! have a nice day",
				},
				expect:      "Hello people! have a bad ",
				expectCount: 3,
			},
			{
				param: param{
					find:    []string{"world", "nice"},
					replace: "cute",
					str:     "Hello world! have a nice day",
				},
				expect:      "Hello cute! have a cute day",
				expectCount: 2,
			},
			{
				param: param{
					find:    []string{"world", ""},
					replace: "cute",
					str:     "Hello world! have a nice day",
				},
				expect:      "Hello cute! have a nice day",
				expectCount: 1,
			},
			{
				param: param{
					find:    "world",
					replace: "baby",
					str:     "Hello world! have a nice day 123",
				},
				expect:      "Hello baby! have a nice day 123",
				expectCount: 1,
			},
			{
				param: param{
					find:    "e",
					replace: "a",
					str:     []string{"wee", "he", "waa"},
				},
				expect:      []string{"waa", "ha", "waa"},
				expectCount: 3,
			},
			{
				param: param{
					find:    []string{"w", "e"},
					replace: "e",
					str:     "wee",
				},
				expect:      "eee",
				expectCount: 4,
			},
			{
				param: param{
					find:    "",
					replace: "a",
					str:     "abc",
				},
				expect:      "abc",
				expectCount: 0,
			},
			{
				param: param{
					find:    []string{"e", "w", "e"},
					replace: "e",
					str:     []string{"wee", "eee"},
				},
				expect:      []string{"eee", "eee"},
				expectCount: 12,
			},
			{
				param: param{
					find:    []string{"", "w", "e"},
					replace: "e",
					str:     []string{"wee", "eee"},
				},
				expect:      []string{"eee", "eee"},
				expectCount: 7,
			},
			{
				param: param{
					find:    "",
					replace: "e",
					str:     []string{"wee", "eee"},
				},
				expect:      []string{"wee", "eee"},
				expectCount: 0,
			},
			// if string not found
			{
				param: param{
					find:    "wkwkkww",
					replace: "baby",
					str:     "Hello world! have a nice day 123",
				},
				expect:      "Hello world! have a nice day 123",
				expectCount: 0,
			},
			{
				param: param{
					find:    []string{"wkwkw", ""},
					replace: []string{"baby", "beautiful"},
					str:     "Hello world! have a nice day 123",
				},
				expect:      "Hello world! have a nice day 123",
				expectCount: 0,
			},
			{
				param: param{
					find:    []string{"a", "e"},
					replace: []string{"b", "d"},
					str:     []string{"abcde", "edcba"},
				},
				expect:      []string{"bbcdd", "ddcbb"},
				expectCount: 4,
			},
			{
				param: param{
					find:    []string{"a", ""},
					replace: []string{"b", "d"},
					str:     []string{"abcde", "edcba"},
				},
				expect:      []string{"bbcde", "edcbb"},
				expectCount: 2,
			},
			{
				param: param{
					find:    []string{"a", "e"},
					replace: []string{"b"},
					str:     []string{"abcde", "edcba"},
				},
				expect:      []string{"bbcd", "dcbb"},
				expectCount: 4,
			},
		}

		for _, test := range tests {
			res, count, err := PHP.StrReplace(test.param.find, test.param.replace, test.param.str)

			if err != nil {
				t.Errorf("\n\tfind: %v\n\treplace: %v\n\tstr: %v\nexpect error nil, but got %s\n",
					test.param.find, test.param.replace, test.param.str, err)
			}
			if reflect.TypeOf(res) != reflect.TypeOf(test.expect) {
				t.Errorf("expect return %T, but got %T", test.expect, res)
			}
			if !reflect.DeepEqual(res, test.expect) {
				t.Errorf("\n\tfind: %v\n\treplace: %v\n\tstr: %v\nexpect %s, but got %s\n",
					test.param.find, test.param.replace, test.param.str, test.expect, res)
			}
			if count != test.expectCount {
				t.Errorf("\n\tfind: %v\n\treplace: %v\n\tstr: %v\nexpect count %d, but got %d\n",
					test.param.find, test.param.replace, test.param.str, test.expectCount, count)
			}
		}
	})
}
