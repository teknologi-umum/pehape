package pehape

import (
	"errors"
	"math"
)

var (
	ErrInvalidSegmentLength = errors.New("the length of each segment must be greater than zero")
)

// StrSplit splits a string into an array.
// Parameters
// - str => string to be splitted
// - length => (optional) only first index is used. Specifies the length of each array element. Default is 1.
// Returns
// - res => array of splitted string
// - err => error if length less than 1
func StrSplit(str string, segLength ...int) (res []string, err error) {
	var charLen int = 1
	var strLen int = len(str)
	if len(segLength) > 0 {
		if segLength[0] < 1 {
			err = ErrInvalidSegmentLength
			return
		}
		charLen = segLength[0]
	}

	for i := 1; i <= int(math.Ceil(float64(strLen)/float64(charLen))); i++ {
		if i*charLen > strLen {
			res = append(res, str[(charLen*i-charLen):])
		} else {
			res = append(res, str[(charLen*i-charLen):i*charLen])
		}
	}
	return
}
