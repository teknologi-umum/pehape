package pehape_test

import (
	"errors"
	"testing"

	PHP "github.com/teknologi-umum/pehape/go"
)

func TestStrSplit(t *testing.T) {
	type param struct {
		str    string
		length []int
	}
	type test struct {
		param  param
		expect []string
	}

	t.Run("It should return error if the given length is less than 1", func(t *testing.T) {
		test := &test{
			param: param{
				str:    "Hello",
				length: []int{0},
			},
		}

		_, err := PHP.StrSplit(test.param.str, test.param.length...)
		if err == nil {
			t.Errorf("expect error, but got error nil")
		}
		if !errors.Is(err, PHP.ErrInvalidSegmentLength) {
			t.Errorf("expect error %s, but got %s", PHP.ErrInvalidSegmentLength, err)
		}
	})

	t.Run("It should split the given string correctly if the given length is empty", func(t *testing.T) {
		test := &test{
			param: param{
				str: "Hello",
			},
			expect: []string{"H", "e", "l", "l", "o"},
		}

		res, err := PHP.StrSplit(test.param.str)

		if err != nil {
			t.Errorf("expect error nil, but got %s", err)
		}
		if len(res[0]) != 1 {
			t.Errorf("expect %v, but got %v", test.expect, res)
		}
		if len(res) != len(test.expect) {
			t.Errorf("expect %v, but got %v", test.expect, res)
		}
		for i, v := range res {
			if test.expect[i] != v {
				t.Errorf("expect %v, but got %v", test.expect, res)
				return
			}
		}
	})

	t.Run("It should split the given string correctly if the given length is not empty", func(t *testing.T) {
		test := &test{
			param: param{
				str:    "Hello",
				length: []int{2},
			},
			expect: []string{"He", "ll", "o"},
		}

		res, err := PHP.StrSplit(test.param.str, test.param.length...)

		if err != nil {
			t.Errorf("expect error nil, but got %s", err)
		}
		if len(res[0]) != test.param.length[0] {
			t.Errorf("expect %v, but got %v", test.expect, res)
		}
		if len(res) != len(test.expect) {
			t.Errorf("expect %v, but got %v", test.expect, res)
		}
		for i, v := range res {
			if test.expect[i] != v {
				t.Errorf("expect %v, but got %v", test.expect, res)
				return
			}
		}
	})
}