package pehape

import (
	"fmt"
)

//  The Implode function joins array/slice element with a string.
//  Parameter :
//  - array => The array/slice to join to a string.
//  - separator => Specifies what to put between the array elements.
//  Return :
//  - The string with all array/slice elements joined.
func Implode[T any](array []T, sep ...string) string {
	var result string
	for i := 0; i < len(array); i++ {
		if sep == nil || len(sep) == 0 {
			result = result + fmt.Sprint(array[i])
		} else {
			if i == 0 {
				result = fmt.Sprint(array[i])
			} else {
				result = result + sep[0] + fmt.Sprint(array[i])
			}
		}
	}
	return result
}
