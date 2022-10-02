package pehape

import "errors"

var (
	ErrStrposInvalidOffset  = errors.New("offset is greather than length of string")
	ErrStrposStringNotFound = errors.New("string not found")
)

// The Strpos() function finds the position of the first occurrence of a string inside another string.
// Parameters
// - str => Specifies the string to search
// - find => String to be find
// - offset => (optional). Specifies where to begin the search. If start is a negative number, it counts from the end of the string.
// Returns
// - int => Returns the position of the first occurrence of a string
// - error => errors if the given offsed is invalid or if string not found
func Strpos(str, find string, offset ...int) (int, error) {
	var o int = 0
	if len(offset) > 0 {
		o = offset[0]
		if o >= len(str) || len(str)+o < 0 {
			return 0, ErrStrposInvalidOffset
		}
	}

	if o < 0 {
		for i := len(str) + o; i <= (len(str) - len(find)); i++ {
			if str[i:i+len(find)] == find {
				return i, nil
			}
		}
		return 0, ErrStrposStringNotFound
	}

	for i := o; i <= (len(str) - len(find)); i++ {
		if str[i:i+len(find)] == find {
			return i, nil
		}
	}
	return 0, ErrStrposStringNotFound
}
