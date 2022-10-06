package pehape

import (
	"testing"
)

func TestGetCapacity(t *testing.T) {
	tests := []struct {
		name     string
		capacity int
		mod      int
		lenS     int
		leng     int
	}{
		{"get mod 0", 10, 0, 100, 10},
		{"get mod 1", 12, 1, 100, 9},
		{"get mod 0", 1, 0, 10, 10},
		{"get mod 1", 2, 1, 11, 10},
		{"get mod 1", 2, 1, 11, 10},
		{"get mod 2", 2, 2, 12, 10},
		{"get mod 3", 2, 3, 13, 10},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			capacityAct, modAct := getCapacity(tt.lenS, tt.leng)
			if capacityAct != tt.capacity || modAct != tt.mod {
				t.Errorf("(%d, %d): expected (%d, %d), actual (%d, %d)", tt.lenS, tt.mod, tt.capacity, tt.mod, capacityAct, modAct)
			}
		})
	}
}

func TestStrSplit(t *testing.T) {
	tests := []struct {
		name        string
		expected    []string
		errExpected error
		s           string
		length      []int
	}{
		{"with no pad \"hello\"", []string{"h", "e", "l", "l", "o"}, nil, "hello", nil},
		{"with pad 3 \"hello\"", []string{"hel", "lo"}, nil, "hello", []int{3}},
		{"with pad 2 \"hello\"", []string{"he", "ll", "o"}, nil, "hello", []int{2}},
		{"with pad 4 \"hello\"", []string{"hell", "o"}, nil, "hello", []int{4}},
		{"with no pad \"Hello Friend\"", []string{"H", "e", "l", "l", "o", " ", "F", "r", "i", "e", "n", "d"}, nil, "Hello Friend", nil},
		{"with pad 5 \"Hello Friend\"", []string{"Hello", " Frie", "nd"}, nil, "Hello Friend", []int{5}},
		{"with no data and pad", []string{""}, nil, "", nil},
		{"with length < 1", []string{}, errMustGreaterThanZero, "", []int{0}},
		{"with 3 params", []string{}, errMust2Params, "", []int{0, 4}},
		{"with no data with pad", []string{""}, nil, "", []int{5}},
		{"with length > len(s)", []string{"hello"}, nil, "hello", []int{500}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			var actual []string
			var err error
			if tt.length == nil {
				actual, err = StrSplit(tt.s)
			} else {
				actual, err = StrSplit(tt.s, tt.length...)
			}

			if tt.errExpected == nil {

				if len(actual) != len(tt.expected) {
					t.Fatalf("length not same, len actual : %d, len expected : %d", len(actual), len(tt.expected))
				}

				for i, v := range tt.expected {
					if v != actual[i] {
						t.Errorf("(%s): expected %s, actual %s", tt.s, v, actual[i])
					}
				}

			} else {
				if err != tt.errExpected {
					t.Errorf("(%s): expected (%v), actual (%v)", tt.s, tt.errExpected, err)
				}
			}
		})
	}
}
