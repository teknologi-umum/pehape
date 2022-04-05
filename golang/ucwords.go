package pehape

import "strings"

/*
 * Ucwords function converts the first character of each word in a string to Uppercase.
 * Parameter :
 * - str => The string to convert.
 * Return :
 * - The string with the first character of each word converted to Uppercase.
 */
func Ucwords(str string) string {
	return strings.Title(strings.ToLower(str))
}
