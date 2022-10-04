package pehape

import (
	"errors"
	"regexp"
	"strings"
)

var (
	ErrTrimInvalidRange = errors.New("invalid character range")
)

// The Trim() function removes whitespace and other predefined characters from both sides of a string.
// Parameters
// - str => string to be checked
// - chars => (optional) Specifies which characters to remove from the string. if omitted,
//   ordinary whitespace, tab, new line, vertical tab, and carriage return are removed
// Returns
// - res => modified string
// - err => errors if given range is invalid

func Trim(str string, chars ...string) (string, error) {
	if len(chars) == 0 {
		return strings.Trim(str, " \t\n\x0B\r"), nil
	}

	var charLists string
	// reverse string to read regex from right side (like php doing)
	reversedString := Strrev(strings.Join(chars, ""))
	// use regex to get range pattern
	r := regexp.MustCompile(`.\.{2}.`)
	// concate non-range char
	charLists = strings.Join(r.Split(reversedString, -1), "")
	// range char pattern
	rangedChar := r.FindAll([]byte(reversedString), -1)
	// generates correct characters between ranges
	for _, val := range rangedChar {
		if val[3] > val[0] {
			return "", ErrTrimInvalidRange
		}

		for i := val[3]; i <= val[0]; i++ {
			charLists += string(i)
		}
	}
	return strings.Trim(str, charLists), nil
}
