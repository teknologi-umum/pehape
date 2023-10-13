package pehape

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	ErrLengthNotValid            = errors.New("length is not valid")
	ErrLengthMustGreaterThanZero = errors.New("length must be greater than zero")

	ErrSeparatorNotValid = errors.New("separator is not valid")
)

// ChunkSplit - Split a string into smaller chunks
// php doc: https://www.php.net/manual/en/function.chunk-split
func ChunkSplit(str string, args ...any) (result string, err error) {
	if len(args) > 2 {
		return "", ErrTooManyArgs
	}

	length := 76
	separator := "\r\n"

	switch len(args) {
	case 1:
		length, err = getChunkSplitLength(args[0])
		if err != nil {
			return "", err
		}
	case 2:
		length, err = getChunkSplitLength(args[0])
		if err != nil {
			return "", err
		}
		separator, err = getChunkSplitSeparator(args[1])
		if err != nil {
			return "", err
		}
	}

	if len(str) < length {
		return str, nil
	}

	var sb strings.Builder

	var index int
	for {
		if len(str) <= index {
			sb.WriteString(str[index:])
			return sb.String(), nil
		}

		if index+length > len(str) {
			length = len(str) - index
		}

		sb.WriteString(str[index : index+length])
		sb.WriteString(separator)

		index += length

	}
}

func getChunkSplitLength(arg any) (int, error) {
	switch v := arg.(type) {
	case int:
		return v, nil
	case float64: // DEPRECATED
		return 0, ErrLengthNotValid
	case string:
		i, err := strconv.Atoi(v)
		if err != nil {
			return 0, ErrLengthNotValid
		}
		return i, nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, ErrLengthMustGreaterThanZero
	}
	return 0, ErrLengthNotValid
}

func getChunkSplitSeparator(arg any) (string, error) {
	switch v := arg.(type) {
	case string:
		return v, nil
	case bool:
		if v {
			return "1", nil
		}
		return "", nil
	case float64, int:
		return fmt.Sprintf("%v", v), nil
	}
	return "", ErrSeparatorNotValid
}
