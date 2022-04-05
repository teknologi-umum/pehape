package pehape

import (
	"errors"
	"strings"
)

/*
 The Explode function breaks a string into an array.
 Parameter :
 - separator => Specifies where to break the string.
 - str => The string to split.
 - limit => Specifies the number of array elements to return.
 If separator is empty then the function returns an error with error message `Separator cannot be empty`.
 If limit is greater than length of element after split, then return all elements.
 If limit is less than minus length of element after split, then return all elements.
 if limit is empty or null then return all elements. If limit is
 greater than 0, then an array/slice with a maximum of limit elements is returned.
 If limit is less than 0, then an array/slice with a maximum of limit elements is returned.
 If the limit is 0, return the result as an array/slice with a single element.

 Return :
 - The array/slice with all elements.
*/

func Explode(separator string, str string, limit ...int) ([]string, error) {
	if separator == "" {
		return nil, errors.New("Separator cannot be empty")
	}

	var result = strings.Split(str, separator)

	if limit == nil {
		return result, nil
	} else {
		limited := limit[0]
		if limited > len(result) || limited < -len(result) {
			return result, nil
		} else if limited == 0 {
			return result[:1], nil
		} else if limited < 0 {
			return result[:len(result)+limited], nil
		} else if limited > 0 {
			return result[:limited], nil
		} else {
			return result, nil
		}
	}

}
