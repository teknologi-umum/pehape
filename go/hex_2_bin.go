package pehape

import (
	"encoding/hex"
)

// The hex2bin() function converts a string of hexadecimal values to ASCII characters.
// Parameter
// -encoded => hexadecimal value to be converted
// Return
// - str => string of ASCII characters
// - error => error if encoded string contain non hexadecimal character
func Hex2Bin(encoded string) (str string, err error) {
	s, err := hex.DecodeString(encoded)
	return string(s), err
}
