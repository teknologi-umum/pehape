package main

import (
	"strings"
)

func Explode(separator, str string, limit int) []string {
	
	var result []string
	result = strings.Split(str, separator)

	if limit == 0 {
		return result[:1]
	} else if limit > len(result) || limit < -len(result) {
		// fmt.Println("limit is greater than result")
		return result
	} else if limit < 0 {
		return result[:len(result)+limit]
	} else if limit > 0 {
		return result[:limit]
	} else {
		return result[limit:]
	}
}