package pehape_test

import (
	"testing"

	PHP "github.com/teknologi-umum/pehape/go"
)

func TestStrrev(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"", ""},
		{"a", "a"},
		{"ab", "ba"},
		{"abc", "cba"},
		{"abcd", "dcba"},
		{"abcde", "edcba"},
		{"abcdef", "fedcba"},
		{"abcdefg", "gfedcba"},
		{"abcdefgh", "hgfedcba"},
		{"abcdefghi", "ihgfedcba"},
		{"abcdefghij", "jihgfedcba"},
		{"abcdefghijk", "kjihgfedcba"},
		{"abcdefghijkl", "lkjihgfedcba"},
		{"abcdefghijklm", "mlkjihgfedcba"},
		{"abcdefghijklmn", "nmlkjihgfedcba"},
		{"abcdefghijklmno", "onmlkjihgfedcba"},
		{"abcdefghijklmnop", "ponmlkjihgfedcba"},
		{"abcdefghijklmnopq", "qponmlkjihgfedcba"},
		{"abcdefghijklmnopqr", "rqponmlkjihgfedcba"},
		{"abcdefghijklmnopqrs", "srqponmlkjihgfedcba"},
		{"abcdefghijklmnopqrst", "tsrqponmlkjihgfedcba"},
		{"abcdefghijklmnopqrstu", "utsrqponmlkjihgfedcba"},
		{"abcdefghijklmnopqrstuv", "vutsrqponmlkjihgfedcba"},
		{"abcdefghijklmnopqrstuvw", "wvutsrqponmlkjihgfedcba"},
		{"abcdefghijklmnopqrstuvwx", "xwvutsrqponmlkjihgfedcba"},
		{"abcdefghijklmnopqrstuvwxy", "yxwvutsrqponmlkjihgfedcba"},
		{"abcdefghijklmnopqrstuvwxyz", "zyxwvutsrqponmlkjihgfedcba"},
		{"Hello World!", "!dlroW olleH"},
		{"kasur rusak", "kasur rusak"},
	}
	for _, test := range tests {
		if got := PHP.Strrev(test.input); got != test.want {
			t.Errorf("Strrev(%q) = %q, want %q", test.input, got, test.want)
		}
	}
}
