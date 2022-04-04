package main

import (
	"fmt"
)

func Implode[T string | int](array []T, sep ...string) string {
	var result string
	for i := 0; i < len(array); i++ {
		if sep == nil {
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
