package pehape

import "unicode"

// The Lcfirst() function converts the first character of a string to lowercase.
// Parameter
// - str => string to be converted
// Return
// - res => converted string
func Lcfirst(str string) string {
	if len(str) == 0 {
		return ""
	}
	char := string(unicode.ToLower(rune(str[0])))
	return char + str[len(char):]
}
