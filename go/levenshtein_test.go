package pehape_test

import (
	"testing"

	PHP "github.com/teknologi-umum/pehape/go"
)

func TestLevenshtein(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"First arg empty", args{"", "abcde"}, 5},
		{"Second arg empty", args{"abcde", ""}, 5},
		{"Same args", args{"abcde", "abcde"}, 0},
		{"ab/aa", args{"ab", "aa"}, 1},
		{"ab/ba", args{"ab", "ba"}, 2},
		{"ab/aaa", args{"ab", "aaa"}, 2},
		{"bbb/a", args{"bbb", "a"}, 3},
		{"kitten/sitting", args{"kitten", "sitting"}, 3},
		{"distance/difference", args{"distance", "difference"}, 5},
		{"a cat/an abct", args{"a cat", "an abct"}, 4},
		{"こにんち/こんにちは", args{"こにんち", "こんにちは"}, 3}, // "Hello" in Japanese
		{"🙂😄🙂😄/😄🙂😄🙂", args{"🙂😄🙂😄", "😄🙂😄🙂"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PHP.Levenshtein(tt.args.str1, tt.args.str2); got != tt.want {
				t.Errorf("levenshteinDistance() = %v, want %v", got, tt.want)
			}
		})
	}
	t.Run("Work with custom cost", func(t *testing.T) {
		if res := PHP.Levenshtein("aa", "ab", 1, 2, 3); res != 2 {
			t.Errorf("expected 2,got %d", res)
		}
		if res := PHP.Levenshtein("aa", "aab", 1, 2, 3); res != 1 {
			t.Errorf("expected 1,got %d", res)
		}
		if res := PHP.Levenshtein("ab", "a", 1, 2, 3); res != 3 {
			t.Errorf("expected 3,got %d", res)
		}
	})
}
