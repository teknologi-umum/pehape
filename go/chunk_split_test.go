package pehape

import "testing"

func TestChunkSplit(t *testing.T) {
	tests := []struct {
		name           string
		given          string
		args           []any
		expectedString string
		expectedError  error
	}{
		{"should return as is 1", "hello", nil, "hello", nil},
		{"should return as is 2", "hello", []any{10}, "hello", nil},
		{"should return as is 3", "hello", []any{10, "oke"}, "hello", nil},
		{"should return chunk 1", "hello", []any{2}, "he\r\nll\r\no\r\n", nil},
		{"should return chunk 2", "hello", []any{2, "oke"}, "heokellokeooke", nil},
		{"should return error len args > 2", "hello", []any{2, "oke", "random"}, "", ErrTooManyArgs},
		{"should return error length not valid because float64", "hello", []any{2.3}, "", ErrLengthNotValid},
		{"should return success with length string", "hello", []any{"2"}, "he\r\nll\r\no\r\n", nil},
		{"should return error because length string float64", "hello", []any{"2.5"}, "", ErrLengthNotValid},
		{"should return success with length bool true", "hello", []any{true}, "h\r\ne\r\nl\r\nl\r\no\r\n", nil},
		{"should return error because string length bool false", "hello", []any{false}, "", ErrLengthMustGreaterThanZero},
		{"should return success with separator bool true", "hello", []any{2, true}, "he1ll1o1", nil},
		{"should return success with separator bool false", "hello", []any{2, false}, "hello", nil},
		{"should return success with separator int", "hello", []any{2, 1}, "he1ll1o1", nil},
		{"should return success with separator float64", "hello", []any{2, 1.5}, "he1.5ll1.5o1.5", nil},
		{"should return error because separator not valid", "hello", []any{2, func() {}}, "", ErrSeparatorNotValid},
		{"should return error because length not valid", "hello", []any{func() {}, func() {}}, "", ErrLengthNotValid},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual, err := ChunkSplit(tt.given, tt.args...)
			if actual != tt.expectedString {
				t.Errorf("(%s): expected %s, actual %s", tt.given, tt.expectedString, actual)
			}

			if err != tt.expectedError {
				t.Errorf("(%s): expected %s, actual %s", tt.given, tt.expectedError, err)
			}
		})
	}
}
