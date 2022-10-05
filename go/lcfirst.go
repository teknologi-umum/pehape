package pehape

import (
	"unicode"
	"unicode/utf8"
)

// The Lcfirst() function converts the first character of a string to lowercase.
// Parameter
// - str => string to be converted
// Return
// - res => converted string
func Lcfirst(str string) string {
	if str == "" {
		return str
	}

	r, size := utf8.DecodeRuneInString(str)
	// convert first char to lowercase
	l := unicode.ToLower(r)
	return string(l) + str[size:]
}
