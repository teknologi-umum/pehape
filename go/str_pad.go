package pehape

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

const (
	STR_PAD_RIGHT int = iota
	STR_PAD_LEFT
	STR_PAD_BOTH
)

var (
	ErrTooManyArgs           = errors.New("too many args")
	ErrPadTypeInvalid        = errors.New("invalid pad type")
	ErrPadInvalid            = errors.New("invalid pad")
	ErrPadMustNotEmptyString = errors.New("pad must be a non-empty string")
)

// StrPad Pad a string to a certain length with another string
// This function returns the string string padded on the left, the right,
// or both sides to the specified padding length.
// If the optional argument pad_string is not supplied, the string is padded with spaces,
// otherwise it is padded with characters from pad_string up to the limit.
// Optional params:
// - pad_string => string or int or float64
// - pad_type => STR_PAD_LEFT, STR_PAD_RIGHT, STR_PAD_BOTH
// php doc: https://www.php.net/manual/en/function.str-pad.php
func StrPad(s string, length int, args ...any) (string, error) {
	// first check args length
	if len(args) > 2 {
		return "", ErrTooManyArgs
	}

	// check if length < len(s)
	if length < len(s) {
		return s, nil
	}

	padString := " "

	// set pad
	if len(args) > 0 {
		switch v := args[0].(type) {
		case string, float64, int:
			padString = fmt.Sprintf("%v", v)
		case bool:
			if v {
				padString = "1"
			} else {
				return "", ErrPadInvalid
			}
		default:
			return "", ErrPadInvalid
		}

		// validate pad only when args[0] not empty
		if padString == "" {
			return "", ErrPadMustNotEmptyString
		}
	}

	// set pad type
	var padType int
	if len(args) == 2 {
		switch v := args[1].(type) {
		case int:
			padType = v
		case float64:
			padType = int(v)
		case string:
			// if its string float64, it will error (DEPRECATED)
			i, err := strconv.Atoi(v)
			if err != nil {
				return "", ErrPadTypeInvalid
			}
			padType = i
		default:
			return "", ErrPadTypeInvalid
		}
	}

	// validate pad type
	if padType < STR_PAD_RIGHT || padType > STR_PAD_BOTH {
		return "", ErrPadTypeInvalid
	}

	// calculate repeat count
	leftRightRepeatCount := int(math.Ceil(float64(length-len(s)) / float64(len(padString))))

	switch padType {
	case STR_PAD_LEFT:
		return strings.Repeat(padString, leftRightRepeatCount)[:length-len(s)] + s, nil
	case STR_PAD_BOTH:
		lenRightFloat64 := float64(length-len(s)) / 2

		lenRight := int(math.Ceil(lenRightFloat64))
		lenLeft := length - len(s) - lenRight

		// +1 to avoid index out of range
		return strings.Repeat(padString, leftRightRepeatCount/2+1)[:lenLeft] + s + strings.Repeat(padString, leftRightRepeatCount/2+1)[:lenRight], nil
	default:
		return s + strings.Repeat(padString, leftRightRepeatCount)[:length-len(s)], nil
	}
}
