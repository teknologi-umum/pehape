package pehape

import (
	"errors"
)

var (
	errMust2Params         = errors.New("expects at most 2 parameters")
	errMustGreaterThanZero = errors.New("The length of each segment must be greater than zero")
)

// StrSplit is to convert a string to an array
// If the optional length parameter is specified, the returned array will be broken down into chunks with each being length in length,
// except the final chunk which may be shorter if the string does not divide evenly.
// The default length is 1, meaning every chunk will be one byte in size.
// If length is less than 1, a nil array and error errMust2Params will be thrown.
// The behaviour of this function is like https://www.php.net/manual/en/function.str-split.php
func StrSplit(s string, length ...int) ([]string, error) {
	if len(length) > 1 {
		return nil, errMust2Params
	}

	chunk := 1
	if len(length) == 1 {
		chunk = length[0]
	}

	if chunk <= 0 {
		return nil, errMustGreaterThanZero
	}

	stringLength := len(s)
	if stringLength == 0 {
		return []string{""}, nil
	}

	if chunk > stringLength {
		return []string{s}, nil
	}

	capacity, mod := getCapacity(stringLength, chunk)

	results := make([]string, 0, capacity)

	to := chunk
	for i := 0; i < stringLength; i += chunk {
		if to <= stringLength {
			results = append(results, s[i:to])
			to += chunk
		}
	}

	if mod > 0 {
		results = append(results, s[stringLength-mod:])
	}
	return results, nil
}

func getCapacity(lenS, n int) (int, int) {
	mod := lenS % n
	if mod == 0 {
		return lenS / n, 0
	}

	return lenS/n + 1, mod
}
