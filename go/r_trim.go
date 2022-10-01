package pehape

import "strings"

// The RTrim() function removes whitespace or other predefined characters from the right side of a string.
// Parameters
//   - str => specifies the string to check
//   - charlists (optional) => Specifies which characters to remove from the string. if omitted,
//     ordinary whitespace, tab, new line, vertical tab, and carriage return are removed
//
// Return
// - res => modified string
func RTrim(str string, charLists ...string) (res string) {
	if len(charLists) == 0 {
		return strings.TrimRight(str, " \t\n\x0B\r")
	}

	return strings.TrimRight(str, strings.Join(charLists, ""))
}
