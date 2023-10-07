package pehape

import (
	"errors"
	"strings"
)

var (
	ErrStrReplaceInvalidParameter = errors.New("must be given parameter correctly")
)

// The StrReplace() function replaces some characters with some other characters in a string.
// but it has a bit different than PHP does:
// 1. it not allowed if parameters has nested slice.
// 2. parameters can only be string or []string, so if you must convert it first into string like PHP does.
// 3. you can get a count of the number of replacements in return value.
// Parameters
// - find => value to find.
// - replace => value to replace the value in 'find'.
// - str => value to be searched.
// Returns
// - interface{} => string or an slice with replaced values
// - int => count of the number of replacements
// - error => errors if parameter is invalid
func StrReplace(find, replace, str interface{}) (interface{}, int, error) {
	switch replaceValue := replace.(type) {
	case string:
		var count int

		switch strValue := str.(type) {
		case string:
			switch findValue := find.(type) {
			case string: // ('replace' == string) && ('str' == string) && ('find' == string)
				if findValue == "" {
					return strValue, count, nil
				}

				if c := strings.Count(strValue, findValue); c > 0 {
					count += c
					strValue = strings.ReplaceAll(strValue, findValue, replaceValue)
				}
				return strValue, count, nil
			case []string: // ('replace' == string) && ('str' == string) && ('find' == slice)
				for i := 0; i < len(findValue); i++ {
					if findValue[i] == "" {
						continue
					}

					if c := strings.Count(strValue, findValue[i]); c > 0 {
						strValue = strings.ReplaceAll(strValue, findValue[i], replaceValue)
						count += c
					}
				}
				return strValue, count, nil
			}
		case []string:
			var count int

			switch findValue := find.(type) {
			case string: // ('replace' == string) && ('str' == slice) && ('find' == string)
				if findValue == "" {
					return strValue, count, nil
				}

				// check every value of 'str' that matches 'find', then replace it
				for i := 0; i < len(strValue); i++ {
					if c := strings.Count(strValue[i], findValue); c > 0 {
						strValue[i] = strings.ReplaceAll(strValue[i], findValue, replaceValue)
						count += c
					}
				}
				return strValue, count, nil
			case []string: // ('replace' == string) && ('str' == slice) && ('find' == slice)
				// check every value of 'str' that matches every value of 'find', then replace it
				for i := 0; i < len(findValue); i++ {
					if findValue[i] == "" {
						continue
					}

					for j := 0; j < len(strValue); j++ {
						if c := strings.Count(strValue[j], findValue[i]); c > 0 {
							strValue[j] = strings.ReplaceAll(strValue[j], findValue[i], replaceValue)
							count += c
						}
					}
				}
				return strValue, count, nil
			}
		}
	case []string:
		switch findValue := find.(type) {
		case []string:
			switch strValue := str.(type) {
			case string: // ('replace' == slice) && ('str' == string) && ('find' == slice)
				var count int

				for i := 0; i < len(findValue); i++ {
					if findValue[i] == "" {
						continue
					}

					if c := strings.Count(strValue, findValue[i]); c > 0 {
						if i < len(replaceValue) {
							strValue = strings.ReplaceAll(strValue, findValue[i], replaceValue[i])
						} else {
							strValue = strings.ReplaceAll(strValue, findValue[i], "")
						}

						count += c
					}
				}
				return strValue, count, nil
			case []string: // ('replace' == slice) && ('str' == slice) && ('find' == slice)
				var count int

				// check every value of 'str' that matches every value of 'find', then replace it
				for i := 0; i < len(findValue); i++ {
					if findValue[i] == "" {
						continue
					}

					for j := 0; j < len(strValue); j++ {
						if c := strings.Count(strValue[j], findValue[i]); c > 0 {
							// if 'charsReplace' out of index then replace with empty string
							if i < len(replaceValue) {
								strValue[j] = strings.ReplaceAll(strValue[j], findValue[i], replaceValue[i])
							} else {
								strValue[j] = strings.ReplaceAll(strValue[j], findValue[i], "")
							}
							count += c
						}
					}
				}
				return strValue, count, nil
			}
		}
	}
	return nil, 0, ErrStrReplaceInvalidParameter
}
