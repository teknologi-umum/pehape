package main

import (
	"errors"
	"strings"
)

func Explode(separator string, str string, limit ...int) []string {
	if separator == "" {
		return []string{errors.New("Separator cannot be empty").Error()}
	}

	var result []string
	result = strings.Split(str, separator)

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
