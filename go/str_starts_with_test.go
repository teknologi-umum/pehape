package pehape_test

import (
	"testing"

	PHP "github.com/teknologi-umum/pehape/go"
)

func TestStrStartsWith(t *testing.T) {
	// Test cases
	testCases := []struct {
		haystack string
		needle   string
		expected bool
	}{
		{"hello world", "hello", true},
		{"hello world", "world", false},
		{"hello world", "h", true},
		{"hello world", "hello world", true},
		{"hello world", "hello world!", false},
		{"", "", true},
		{"", "hello", false},
		{"Hello World", "", true},
		{"hello world", "Hello", false},
		{"😅", "🙂", false},
		{"你好，世界", "好", false},
		{"你好，世界", "你好", true},
		{"こんにちは、世界", "こんにちは", true},
	}

	// Iterate over test cases
	for _, tc := range testCases {
		// Run test
		t.Run(tc.haystack+"-"+tc.needle, func(t *testing.T) {
			if got := PHP.StrStartsWith(tc.haystack, tc.needle); got != tc.expected {
				t.Errorf("StrStartsWith(%q, %q) = %v, expected %v", tc.haystack, tc.needle, got, tc.expected)
			}
		})
	}
}
