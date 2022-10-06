package pehape

import (
	"errors"
	"reflect"
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
	vFind := reflect.ValueOf(find)
	vReplace := reflect.ValueOf(replace)
	vStr := reflect.ValueOf(str)

	switch vReplace.Kind() {
	case reflect.String:

		var charsReplace = replace.(string)
		var count int

		// check if given 'str' is a slice
		switch vStr.Kind() {
		case reflect.String:

			var strCopy = str.(string)

			switch vFind.Kind() {
			case reflect.String: // ('replace' == string) && ('str' == string) && ('find' == string)
				if c := strings.Count(strCopy, find.(string)); c > 0 {
					count += c
					strCopy = strings.ReplaceAll(strCopy, find.(string), charsReplace)
				}
				return strCopy, count, nil

			case reflect.Slice: // ('replace' == string) && ('str' == string) && ('find' == slice)
				var charsFind []string
				// only slice of string are allowed for 'find'
				charsFind, ok := find.([]string)
				if !ok {
					return nil, 0, ErrStrReplaceInvalidParameter
				}

				if len(charsFind) > 0 && len(charsFind[0]) != 0 {
					for i := 0; i < len(charsFind); i++ {
						if c := strings.Count(strCopy, charsFind[i]); c > 0 {
							strCopy = strings.ReplaceAll(strCopy, charsFind[i], charsReplace)
							count += c
						}
					}
				}
				return strCopy, count, nil
			}

		case reflect.Slice: // ('replace' == string) && ('str' == slice)

			var charsStr = []string{}
			var count int

			// only slice of string are allowed for 'str'
			charsStr, ok := str.([]string)
			if !ok {
				return nil, 0, ErrStrReplaceInvalidParameter
			}

			switch vFind.Kind() {
			case reflect.String: // ('replace' == string) && ('str' == slice) && ('find' = string)
				if len(find.(string)) == 0 {
					return charsStr, 0, nil
				}

				// check every value of 'charsStr' that matches 'find', then replace it
				for i := 0; i < len(charsStr); i++ {
					if c := strings.Count(charsStr[i], find.(string)); c > 0 {
						charsStr[i] = strings.ReplaceAll(charsStr[i], find.(string), replace.(string))
						count += c
					}
				}
				return charsStr, count, nil

			case reflect.Slice: // ('replace' == string) && ('str' == slice) && ('find' == slice)
				// only slice of string are allowed for 'find'
				var charsFind []string

				charsFind, ok := find.([]string)
				if !ok {
					return nil, 0, ErrStrReplaceInvalidParameter
				}

				// check every value of 'str' that matches every value of 'find', then replace it
				for i := 0; i < len(charsFind); i++ {
					if charsFind[i] == "" {
						continue
					}
					for j := 0; j < len(charsStr); j++ {
						if c := strings.Count(charsStr[j], charsFind[i]); c > 0 {
							charsStr[i] = strings.ReplaceAll(charsStr[j], charsFind[i], replace.(string))
							count += c
						}
					}
				}
				return charsStr, count, nil
			}
		}
	case reflect.Slice: // ('replace' == slice)
		// the 'find' also must be a slice, otherwise return an error
		var charsReplace []string
		// only slice of string are allowed for 'replace'
		charsReplace, ok := replace.([]string)
		if !ok {
			return nil, 0, ErrStrReplaceInvalidParameter
		}

		switch vFind.Kind() {
		case reflect.Slice: // ('replace' == slice) && ('find' == slice)

			var charsFind []string

			// only slice of string are allowed for 'find'
			charsFind, ok := find.([]string)
			if !ok {
				return nil, 0, ErrStrReplaceInvalidParameter
			}

			switch vStr.Kind() {
			case reflect.Slice: // ('replace' == slice) && ('find' == slice) && ('str' == slice)
				var charsStr []string
				var count int

				// only slice of string are allowed for 'str'
				charsStr, ok := str.([]string)
				if !ok {
					return nil, 0, ErrStrReplaceInvalidParameter
				}

				// check every value of 'str' that matches every value of 'find', then replace it
				for i := 0; i < len(charsFind); i++ {
					if charsFind[i] == "" {
						continue
					}
					for j := 0; j < len(charsStr); j++ {
						if c := strings.Count(charsStr[j], charsFind[i]); c > 0 {
							// if 'charsReplace' out of index then replace with empty string
							if i < len(charsReplace) {
								charsStr[j] = strings.ReplaceAll(charsStr[j], charsFind[i], charsReplace[i])
							} else {
								charsStr[j] = strings.ReplaceAll(charsStr[j], charsFind[i], "")
							}
							count += c
						}
					}
				}
				return charsStr, count, nil
			case reflect.String: // ('replace' == slice) && ('find' == slice) && ('str' == string)
				var strCopy = str.(string)
				var count int
				// check every value of 'charsFind' that matches 'str', then replace it
				for i := 0; i < len(charsFind); i++ {
					if charsFind[i] == "" {
						continue
					}
					if c := strings.Count(strCopy, charsFind[i]); c > 0 {
						if i < len(charsReplace) {
							strCopy = strings.ReplaceAll(strCopy, charsFind[i], charsReplace[i])
						} else {
							strCopy = strings.ReplaceAll(strCopy, charsFind[i], "")
						}
						count += c
					}
				}
				return strCopy, count, nil
			}
		}
	}
	return nil, 0, ErrStrReplaceInvalidParameter
}
