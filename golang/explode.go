package main

import (
	"strings"
)

func Explode(delimiter string, str string, limit ...int) []string {
	var result []string
	result = strings.Split(str, delimiter)

	if limit == nil || limit[0] > len(result) || limit[0] < -len(result) {
		return result
	} else if limit[0] == 0 {
		return result[:1]
	} else if limit[0] < 0 {
		return result[:len(result)+limit[0]]
	} else if limit[0] > 0 {
		return result[:limit[0]]
	} else {
		return result
	}
}
