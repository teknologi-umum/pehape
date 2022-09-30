package pehape

import "encoding/hex"

// The bin2hex() function converts a string of ASCII characters to hexadecimal values.
// Parameter:
// - str => string to be converted
// Return:
// - encoded => converted string
func Bin2Hex(str string) (encoded string) {
	return hex.EncodeToString([]byte(str))
}
