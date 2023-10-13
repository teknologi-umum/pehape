package pehape

import "strings"

// The StrContains() function checks if a string contains a specified string.
// Parameters
// - haystack => the string to search in.
// - needle => the string to search for.
// Returns
// - bool => true if the string contains the search value, false otherwise.
func StrContains(haystack, needle string) bool {
	return strings.Contains(haystack, needle)
}
