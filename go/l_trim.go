package pehape

import (
	"fmt"
	"regexp"
	"strings"
)

// The LTrim() function removes whitespace or other predefined characters from the left side of a string.
// Parameters
//   - str => specifies the string to check
//   - chars (optional) => Specifies which characters to remove from the string. if omitted,
//     ordinary whitespace, tab, new line, vertical tab, and carriage return are removed
//
// Return
// - res => modified string
// - err => error if given range is invalid
func LTrim(str string, chars ...string) (res string, err error) {
	if len(chars) == 0 {
		return strings.TrimLeft(str, " \t\n\x0B\r"), nil
	}
	s := strings.Join(chars, "")
	var charLists string
	// reverse string to read regex from right side (like php doing)
	rStr := rev(s)
	// use regex to get range pattern
	r := regexp.MustCompile(`.\.{2}.`)
	// concate non-range char
	charLists = strings.Join(r.Split(rStr, -1), "")
	// range char pattern
	rChar := r.FindAll([]byte(rStr), -1)
	// generates correct characters between ranges
	for _, val := range rChar {
		if val[3] > val[0] {
			return "", fmt.Errorf("invalid range")
		}

		for i := val[3]; i <= val[0]; i++ {
			charLists += string(i)
		}
	}
	return strings.TrimLeft(str, charLists), nil
}

// todo: use pehape.Strrev
func rev(str string) (res string) {
	for i := len(str) - 1; i >= 0; i-- {
		res += string(str[i])
	}
	return
}
