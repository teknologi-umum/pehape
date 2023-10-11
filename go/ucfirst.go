package pehape

import (
	"unicode"
	"unicode/utf8"
)

// The Ucfirst() function converts the first character of a string to uppercase.
// Parameter
// - str => string to be converted
// Return
// - res => converted string
func Ucfirst(str string) string {
	if str == "" {
		return ""
	}

	r, size := utf8.DecodeRuneInString(str)

	u := unicode.ToUpper(r)
	return string(u) + str[size:]
}
