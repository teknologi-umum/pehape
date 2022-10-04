package pehape

import "unicode"

// The Lcfirst() function converts the first character of a string to lowercase.
// Parameter
// - str => string to be converted
// Return
// - res => converted string
func Lcfirst(str string) string {
	for _, val := range str {
		char := string(unicode.ToLower(val))
		return char + str[len(char):]
	}
	return ""
}
