package pehape

import "strings"

// The LTrim() function removes whitespace or other predefined characters from the left side of a string.
// Parameters
//   - str => specifies the string to check
//   - charLists (optional) => Specifies which characters to remove from the string. if omitted,
//     ordinary whitespace, tab, new line, vertical tab, and carriage return are removed
//
// Return
// - res => modified string
func LTrim(str string, charLists ...string) (res string) {
	if len(charLists) == 0 {
		return strings.TrimLeft(str, " \t\n\x0B\r")
	}

	return strings.TrimLeft(str, strings.Join(charLists, ""))
}
