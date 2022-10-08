package pehape

import "bytes"

// Chr is a function that returns a character from a number.
// Parameter :
// - codepoint => it can be in decimal, octal, or hex values.
// Return :
// - single-byte string
func Chr(codepoint int) string {
	var result bytes.Buffer
	result.WriteByte(byte(codepoint))

	return result.String()
}
