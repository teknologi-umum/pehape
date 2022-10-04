package pehape

import (
	"errors"
)

var (
	ErrStrrposInvalidOffset  = errors.New("offset is greather than length of string")
	ErrStrrposStringNotFound = errors.New("string is not found")
)

// The Strrpos() function finds the position of the last occurrence of a string inside another string.
// Parameters
//   - haystack => The string to search in.
//   - needle => Specifies the string to find
//   - offset =>  If zero or positive, the search is performed left to right skipping the first offset bytes of the haystack.
//     If negative, the search is performed right to left skipping the last offset bytes of the haystack and searching for the first occurrence of needle.
func Strrpos(haystack, needle string, offset ...int) (int, error) {
	var o int = 0
	if len(offset) > 0 {
		o = offset[0]
		if o >= len(haystack) || (len(haystack)+o < 0) {
			return 0, ErrStrrposInvalidOffset
		}
	}
	if o < 0 {
		// read string from right
		for i := len(haystack) + o; i >= len(needle); i-- {
			left := (i - len(needle)) + 1
			right := i + 1
			if haystack[left:right] == needle {
				return left, nil
			}
		}
		return 0, ErrStrrposStringNotFound
	}

	// read string from right
	for i := len(haystack); i >= len(needle)+o; i-- {
		left := i - len(needle)
		right := i
		if haystack[left:right] == needle {
			return left, nil
		}
	}
	return 0, ErrStrrposStringNotFound
}
