package pehape

import (
	"errors"
	"fmt"
	"math"
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

func StrPad(s string, length int, args ...any) (string, error) {
	// first check args length
	if len(args) > 2 {
		return "", ErrTooManyArgs
	}

	// check if length < len(s)
	if length < len(s) {
		return s, nil
	}

	pad := " "

	// set pad
	if len(args) > 0 {
		switch args[0].(type) {
		case string, float64, int:
			pad = fmt.Sprintf("%v", args[0])
		default:
			return "", ErrPadInvalid
		}

		// validate pad only when args[0] not empty
		if pad == "" {
			return "", ErrPadMustNotEmptyString
		}
	}

	// set pad type
	var padType int
	if len(args) == 2 {
		switch args[1].(type) {
		case int:
			padType = args[1].(int)
		case float64:
			padType = int(args[1].(float64))
		default:
			return "", ErrPadTypeInvalid
		}
	}

	// validate pad type
	if padType < STR_PAD_RIGHT || padType > STR_PAD_BOTH {
		return "", ErrPadTypeInvalid
	}

	// calculate repeat count
	leftRightRepeatCount := int(math.Ceil(float64(length-len(s)) / float64(len(pad))))

	switch padType {
	case STR_PAD_LEFT:
		return strings.Repeat(pad, leftRightRepeatCount)[:length-len(s)] + s, nil
	case STR_PAD_BOTH:
		lenRightFloat64 := float64(length-len(s)) / 2

		lenRight := int(math.Ceil(lenRightFloat64))
		lenLeft := length - len(s) - lenRight

		// +1 to avoid index out of range
		return strings.Repeat(pad, leftRightRepeatCount/2+1)[:lenLeft] + s + strings.Repeat(pad, leftRightRepeatCount/2+1)[:lenRight], nil
	default:
		return s + strings.Repeat(pad, leftRightRepeatCount)[:length-len(s)], nil
	}
}
