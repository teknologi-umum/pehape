package pehape

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrPad(t *testing.T) {
	tests := []struct {
		name           string
		expectedString string
		expectedError  error
		s              string
		length         int
		args           []any
	}{
		{
			name:           `no args return as is`,
			expectedString: "1000",
			expectedError:  nil,
			s:              "1000",
			length:         1,
		},
		{
			name:           `no args`,
			expectedString: "1000     ",
			expectedError:  nil,
			s:              "1000",
			length:         9,
		},
		{
			name:           `args only pad return as is`,
			expectedString: "1000",
			expectedError:  nil,
			s:              "1000",
			length:         1,
			args:           []any{"1"},
		},
		{
			name:           `args only pad 1`,
			expectedString: "100011111",
			expectedError:  nil,
			s:              "1000",
			length:         9,
			args:           []any{"1"},
		},
		{
			name:           `args only pad 2`,
			expectedString: "100012121",
			expectedError:  nil,
			s:              "1000",
			length:         9,
			args:           []any{"12"},
		},
		{
			name:           `args only pad 3`,
			expectedString: "100012312",
			expectedError:  nil,
			s:              "1000",
			length:         9,
			args:           []any{"123"},
		},
		{
			name:           `args only pad 4`,
			expectedString: "100012341",
			expectedError:  nil,
			s:              "1000",
			length:         9,
			args:           []any{"1234"},
		},
		{
			name:           `args only pad 5`,
			expectedString: "100012345",
			expectedError:  nil,
			s:              "1000",
			length:         9,
			args:           []any{"12345"},
		},
		{
			name:           `args only with float`,
			expectedString: "10001.51.",
			expectedError:  nil,
			s:              "1000",
			length:         9,
			args:           []any{1.5},
		},
		{
			name:           `args with left pad return as is`,
			expectedString: "1000",
			expectedError:  nil,
			s:              "1000",
			length:         1,
			args:           []any{"_", STR_PAD_LEFT},
		},
		{
			name:           `args with left pad type 1`,
			expectedString: "_____1000",
			expectedError:  nil,
			s:              "1000",
			length:         9,
			args:           []any{"_", STR_PAD_LEFT},
		},
		{
			name:           `args with left pad type 2`,
			expectedString: "121211000",
			expectedError:  nil,
			s:              "1000",
			length:         9,
			args:           []any{"12", STR_PAD_LEFT},
		},
		{
			name:           `args with left pad type 3`,
			expectedString: "123121000",
			expectedError:  nil,
			s:              "1000",
			length:         9,
			args:           []any{"123", STR_PAD_LEFT},
		},
		{
			name:           `args with left pad type 4`,
			expectedString: "123411000",
			expectedError:  nil,
			s:              "1000",
			length:         9,
			args:           []any{"1234", STR_PAD_LEFT},
		},
		{
			name:           `args with left pad type 5`,
			expectedString: "123451000",
			expectedError:  nil,
			s:              "1000",
			length:         9,
			args:           []any{"12345", STR_PAD_LEFT},
		},
		{
			name:           `args with left pad float param`,
			expectedString: "1.51.1000",
			expectedError:  nil,
			s:              "1000",
			length:         9,
			args:           []any{1.5, STR_PAD_LEFT},
		},
		{
			name:           `args with both pad type return as is`,
			expectedString: "1000",
			expectedError:  nil,
			s:              "1000",
			length:         1,
			args:           []any{"1", STR_PAD_BOTH},
		},
		{
			name:           `args with both pad type 1`,
			expectedString: "111000111",
			expectedError:  nil,
			s:              "1000",
			length:         9,
			args:           []any{"1", STR_PAD_BOTH},
		},
		{
			name:           `args with both pad type 2`,
			expectedString: "121000121",
			expectedError:  nil,
			s:              "1000",
			length:         9,
			args:           []any{"12", STR_PAD_BOTH},
		},
		{
			name:           `args with both pad type 3`,
			expectedString: "121000123",
			expectedError:  nil,
			s:              "1000",
			length:         9,
			args:           []any{"123", STR_PAD_BOTH},
		},
		{
			name:           `args with both pad type 4`,
			expectedString: "121000123",
			expectedError:  nil,
			s:              "1000",
			length:         9,
			args:           []any{"1234", STR_PAD_BOTH},
		},
		{
			name:           `args with both pad type 5`,
			expectedString: "121000123",
			expectedError:  nil,
			s:              "1000",
			length:         9,
			args:           []any{"12345", STR_PAD_BOTH},
		},
		{
			name:           `args with both pad type 5 int param`,
			expectedString: "121000123",
			expectedError:  nil,
			s:              "1000",
			length:         9,
			args:           []any{12345, STR_PAD_BOTH},
		},
		{
			name:           `args with both pad with float param`,
			expectedString: "12100012.",
			expectedError:  nil,
			s:              "1000",
			length:         9,
			args:           []any{12.5, STR_PAD_BOTH},
		},
		{
			name:           `args with both pad with float param 1`,
			expectedString: "121000121",
			expectedError:  nil,
			s:              "1000",
			length:         9,
			args:           []any{12.0, STR_PAD_BOTH},
		},
		{
			name:           `args with both pad with float param 2`,
			expectedString: "1.51.5110001.51.51.",
			expectedError:  nil,
			s:              "1000",
			length:         19,
			args:           []any{1.5, STR_PAD_BOTH},
		},
		{
			name:           `args with both pad with float param 2 and float pad type`,
			expectedString: "1.51.5110001.51.51.",
			expectedError:  nil,
			s:              "1000",
			length:         19,
			args:           []any{1.5, 2.0},
		},
		{
			name:           `should error too many args`,
			expectedString: "",
			expectedError:  ErrTooManyArgs,
			s:              "1000",
			length:         19,
			args:           []any{1.5, STR_PAD_BOTH, "oke"},
		},
		{
			name:           `should error pad type 1`,
			expectedString: "",
			expectedError:  ErrPadTypeInvalid,
			s:              "1000",
			length:         19,
			args:           []any{1.5, -1},
		},
		{
			name:           `should error pad type 2`,
			expectedString: "",
			expectedError:  ErrPadTypeInvalid,
			s:              "1000",
			length:         19,
			args:           []any{1.5, 10},
		},
		{
			name:           `should error pad type 3`,
			expectedString: "",
			expectedError:  ErrPadTypeInvalid,
			s:              "1000",
			length:         19,
			args:           []any{1.5, func() {}},
		},
		{
			name:           `should error pad must not empty string`,
			expectedString: "",
			expectedError:  ErrPadMustNotEmptyString,
			s:              "1000",
			length:         19,
			args:           []any{"", 10},
		},
		{
			name:           `should error pad invalid`,
			expectedString: "",
			expectedError:  ErrPadInvalid,
			s:              "1000",
			length:         19,
			args:           []any{func() {}, 10},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual, err := StrPad(tt.s, tt.length, tt.args...)
			assert.Equal(t, tt.expectedString, actual)
			assert.Equal(t, tt.expectedError, err)

			if tt.length > len(tt.s) && tt.expectedError == nil {
				assert.Equal(t, tt.length, len(actual))
			}
		})
	}
}
