package pehape

import (
	"errors"
)

var (
	ErrSttrposInvalidOffset  = errors.New("offset is greather than length of string")
	ErrSttrposStringNotFound = errors.New("string is not found")
)

// The Strrpos() function finds the position of the last occurrence of a string inside another string.
// Parameters
//   - str => The string to search in.
//   - find => Specifies the string to find
//   - offset =>  If zero or positive, the search is performed left to right skipping the first offset bytes of the str.
//     If negative, the search is performed right to left skipping the last offset bytes of the str and searching for the first occurrence of find.
func Strrpos(str, find string, offset ...int) (int, error) {
	var o int = 0
	if len(offset) > 0 {
		o = offset[0]
		if o >= len(str) || (len(str)+o < 0) {
			return 0, ErrSttrposInvalidOffset
		}
	}
	if o < 0 {
		// read string from right
		for i := len(str) + o; i >= len(find); i-- {
			left := (i - len(find)) + 1
			right := i + 1
			if str[left:right] == find {
				return left, nil
			}
		}
		return 0, ErrSttrposStringNotFound
	}

	// read string from right
	for i := len(str); i >= len(find)+o; i-- {
		left := i - len(find)
		right := i
		if str[left:right] == find {
			return left, nil
		}
	}
	return 0, ErrSttrposStringNotFound
}
