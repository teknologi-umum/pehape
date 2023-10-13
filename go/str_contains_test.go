package pehape_test

import (
	"testing"

	PHP "github.com/teknologi-umum/pehape/go"
)

func TestStrContains(t *testing.T) {
	testCases := []struct {
		name     string
		haystack string
		needle   string
		expected bool
	}{
		{
			name:     "needle is in the middle of the haystack",
			haystack: "hello world",
			needle:   "world",
			expected: true,
		},
		{
			name:     "needle is at the beginning of the haystack",
			haystack: "hello world",
			needle:   "hello",
			expected: true,
		},
		{
			name:     "needle is at the end of the haystack",
			haystack: "hello world",
			needle:   "world",
			expected: true,
		},
		{
			name:     "needle is not in the haystack",
			haystack: "hello world",
			needle:   "goodbye",
			expected: false,
		},
		{
			name:     "needle is an empty string",
			haystack: "hello world",
			needle:   "",
			expected: true,
		},
		{
			name:     "haystack is an empty string",
			haystack: "",
			needle:   "world",
			expected: false,
		},
		{
			name:     "both haystack and needle are empty strings",
			haystack: "",
			needle:   "",
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := PHP.StrContains(tc.haystack, tc.needle)
			if result != tc.expected {
				t.Errorf("expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
