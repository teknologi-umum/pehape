package pehape

import (
	"errors"
	"regexp"
)

var (
	ErrStrchrInvalidParameterType = errors.New("invalid parameter type")
	ErrStrchrStringNotFound       = errors.New("string not found")
)

// The Strchr() function searches for the first occurrence of a string inside another string.
// Parameter
//   - str => string to search
//   - search => Specifies the string to search for. If this parameter is a number,
//     it will search for the character matching the ASCII value of the number
//   - beforeSearch => Optional. A boolean value whose default is "false".
//     If set to "true", it returns the part of the string before the first occurrence of the search parameter.
//
// Returns
// - string => Returns the rest of the string (from the matching point)
// - error => either string not found, or the parameter search is invalid
func Strchr(str string, search interface{}, beforeSearch ...bool) (string, error) {
	var searchedCharacter string

	if val, ok := search.(int); ok {
		searchedCharacter = string(rune(val))
	} else {
		if val, ok := search.(string); ok {
			searchedCharacter = val
		} else {
			return "", ErrStrchrInvalidParameterType
		}
	}

	// search using regexp
	rgx := regexp.MustCompile(searchedCharacter)
	loc := rgx.FindStringIndex(str)
	if loc == nil {
		// if string not found
		return "", ErrStrchrStringNotFound
	}

	if len(beforeSearch) == 0 || !beforeSearch[0] {
		// if beforeSearch parameters is not defined or if the first index of beforeSearch is FALSE
		return str[loc[0]:], nil
	}
	// if beforeSearch parameters is defined and if the first index of beforeSearch is TRUE
	return str[:loc[0]], nil
}
