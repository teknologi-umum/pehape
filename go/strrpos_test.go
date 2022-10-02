package pehape_test

import (
	"errors"
	"testing"

	PHP "github.com/teknologi-umum/pehape/go"
)

func TestStrrpos(t *testing.T) {
	type param struct {
		str    string
		find   string
		offset int
	}
	type test struct {
		param  param
		expect int
	}

	t.Run("It should return error if the given offset is invalid", func(t *testing.T) {
		tt := &test{
			param: param{
				str:    "GoLang",
				find:   "Go",
				offset: 6,
			},
		}
		_, err := PHP.Strrpos(tt.param.str, tt.param.find, tt.param.offset)
		if err == nil {
			t.Errorf("expect error, but got nil")
		}
		if !errors.Is(err, PHP.ErrSttrposInvalidOffset) {
			t.Errorf("expect error %s, but got %s", PHP.ErrSttrposInvalidOffset, err)
		}
	})

	t.Run("It should return error if the given offset is invalid", func(t *testing.T) {
		tt := &test{
			param: param{
				str:    "GoLang",
				find:   "Go",
				offset: -7,
			},
		}
		_, err := PHP.Strrpos(tt.param.str, tt.param.find, tt.param.offset)
		if err == nil {
			t.Errorf("expect error, but got nil")
		}
		if !errors.Is(err, PHP.ErrSttrposInvalidOffset) {
			t.Errorf("expect error %s, but got %s", PHP.ErrSttrposInvalidOffset, err)
		}
	})

	t.Run("It should return the position of the given string correctly if the given offset empty", func(t *testing.T) {
		tt := &test{
			param: param{
				str:  "I love Go, I love Go too!",
				find: "Go",
			},
			expect: 18,
		}

		res, err := PHP.Strrpos(tt.param.str, tt.param.find)

		if err != nil {
			t.Errorf("expect error nil, but got %s", err)
		}
		if res != tt.expect {
			t.Errorf("expect %d, but got %d", tt.expect, res)
		}
	})

	t.Run("It should return the position of the given string correctly if the given offset is not empty", func(t *testing.T) {
		tt := &test{
			param: param{
				str:    "I love Go, I love Go too!",
				find:   "Go",
				offset: 10,
			},
			expect: 18,
		}

		res, err := PHP.Strrpos(tt.param.str, tt.param.find, tt.param.offset)

		if err != nil {
			t.Errorf("expect error nil, but got %s", err)
		}
		if res != tt.expect {
			t.Errorf("expect %d, but got %d", tt.expect, res)
		}
	})

	t.Run("It should return the position of the given string correctly if the given offset is not empty", func(t *testing.T) {
		tt := &test{
			param: param{
				str:    "I love Go, I love Go too!",
				find:   "Go",
				offset: -10,
			},
			expect: 7,
		}

		res, err := PHP.Strrpos(tt.param.str, tt.param.find, tt.param.offset)

		if err != nil {
			t.Errorf("expect error nil, but got %s", err)
		}
		if res != tt.expect {
			t.Errorf("expect %d, but got %d", tt.expect, res)
		}
	})

	t.Run("It should return error if string not found", func(t *testing.T) {
		tt := &test{
			param: param{
				str:    "I love Go, I love Go too!",
				find:   "Go",
				offset: -25,
			},
		}

		_, err := PHP.Strrpos(tt.param.str, tt.param.find, tt.param.offset)
		if err == nil {
			t.Errorf("expect error, but got nil")
		}
		if !errors.Is(err, PHP.ErrSttrposStringNotFound) {
			t.Errorf("expect error nil, but got %s", err)
		}
	})

	t.Run("It should return error if string not found", func(t *testing.T) {
		tt := &test{
			param: param{
				str:  "I love Go, I love Go too!",
				find: "hahaha",
			},
		}

		_, err := PHP.Strrpos(tt.param.str, tt.param.find, tt.param.offset)
		if err == nil {
			t.Errorf("expect error, but got nil")
		}
		if !errors.Is(err, PHP.ErrSttrposStringNotFound) {
			t.Errorf("expect error nil, but got %s", err)
		}
	})
}
