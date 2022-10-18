package pehape

import (
	"errors"
	"strconv"
	"strings"
	"sync"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var (
	errMust4Params   = errors.New("expects at most 4 parameters")
	errParam2MustInt = errors.New("expects parameter 2 to be int")
)

var (
	messagePrinterNumberFormat *message.Printer
	onceNumberFormat           sync.Once
)

func init() {
	onceNumberFormat.Do(func() {
		messagePrinterNumberFormat = message.NewPrinter(language.English)
	})
}

// NumberFormat is for formats a number with grouped thousands and optionally decimal digits.
//
// num : The number being formatted.
// option 1 : Sets the number of decimal digits. If 0, the decimal_separator is omitted from the return value. Default = 0
// option 2 : Sets the separator for the decimal point. Default = "."
// option 3 : Sets the thousands separator. Default = ","
// The behaviour of this function is like https://www.php.net/manual/en/function.number-format.php
func NumberFormat(num float64, options ...string) (string, error) {
	if len(options) > 3 {
		return "", errMust4Params
	}
	strFormat := "%.0"

	if len(options) > 0 {
		fixInteger := getIntegerFromString(options[0])
		_, err := strconv.Atoi(fixInteger)
		if err != nil {
			return "", errParam2MustInt
		}

		strFormat = strFormat[:len(strFormat)-1] + fixInteger
	}

	result := messagePrinterNumberFormat.Sprintf(strFormat+"f", num)

	if len(options) >= 2 {
		result = strings.ReplaceAll(result, ".", options[1])
	}

	if len(options) == 3 {
		result = strings.ReplaceAll(result, ",", options[2])
	}

	return result, nil
}

// php magic
func getIntegerFromString(s string) string {
	for i, v := range s {
		if v < 48 || v > 57 {
			return s[:i]
		}
	}
	return s
}
