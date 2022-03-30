package main

import (
	"strings"
)

//make explode like php explode
// func Explode(separator string, str string) []string {
// 	var result []string
// 	var temp string
// 	var i int
// 	var j int

// 	for i = 0; i < len(str); i++ {
// 		if str[i] == separator[0] {
// 			for j = 0; j < len(separator); j++ {
// 				if str[i+j] != separator[j] {
// 					break
// 				}
// 			}
// 			if j == len(separator) {
// 				i = i + j
// 				result = append(result, temp)
// 				temp = ""
// 			} else {
// 				temp = temp + string(str[i])
// 			}
// 		} else {
// 			temp = temp + string(str[i])
// 		}
// 	}
// 	result = append(result, temp)
// 	return result
// }

// //make explode function Returns an array of strings, each of which is a substring of string formed by splitting it on boundaries formed by the string separator.
// func Explode(separator string, str string) []string {
// 	var result []string
// 	var temp string
// 	var i int
// 	var j int

// 	for i = 0; i < len(str); i++ {
// 		if str[i] == separator[0] {
// 			for j = 0; j < len(separator); j++ {
// 				if str[i+j] != separator[j] {
// 					break
// 				}
// 			}
// 			if j == len(separator) {
// 				i = i + j
// 				result = append(result, temp)
// 				// temp = ""
// 			} else {
// 				temp = temp + string(str[i])
// 			}
// 		} else {
// 			temp = temp + string(str[i])
// 		}
// 	}
// 	result = append(result, temp)
// 	return result
// }

/// <summary>
/// The Explode function breaks a string into an array.
/// </summary>
/// <param name="separator">Specifies where to break the string.</param>
/// <param name="str">The string to split.</param>
/// <param name="limit">Specifies the number of array elements to return.
/// If limit is greater than 0, then an array with a maximum of limit
/// elements is returned. If limit is less than 0 then an array excluding
/// the last -limit elements is returned. If limit is 0 then an array with
/// one element is returned.</param>
/// <returns>An array of strings.</returns>

//Returns an array of strings, each of which is a substring of string formed by splitting it on boundaries formed by the string separator.
func Explode(separator, str string) []string {
	var result []string
	if len(separator) > len(str) {
		a := strings.Split(str, separator)
		result = append(result, a...)
	} else {
		a := strings.SplitAfterN(str, separator, -1)
		result = append(result, a...)
	}
	return result
}

func Expl(sep, str string, limit int) []string {
	var result []string
	if len(sep) > len(str) {
		result = strings.Split(str, sep)
	} else {
		result = strings.SplitAfterN(str, sep, -1)
	}
	return result
}

func Exp(separator, str string, limit int) []string {
	var result []string
	if len(separator) > len(str) {
		result = strings.Split(str, separator)
	} else {
		result = strings.SplitAfterN(str, separator, -1)
	}

	if limit == 0 {
		return result[:1]
	} else if limit > len(result) || limit < -len(result) {
		// fmt.Println("limit is greater than result")
		return result
	} else if limit < 0 {
		return result[:len(result)+limit]
	} else if limit > 0 {
		return result[:limit]
	} else {
		return result[limit:]
	}

	// switch limit {
	// case 0:
	// 	result = result[:1]
	// case limit < 0:
	// 	result = result[:len(result)-limit]
	// case limit > 0:
	// 	result = result[:limit]
	// case limit > len(result) || limit < +len(result):
	// 	fmt.Println("limit is greater than result")
	// 	return result
	// }

	// if limit == 0 || limit == 1 {
	// 	result = result[:1]
	// } else if limit > len(result) {
	// 	result = result[:limit]
	// } else if limit == -limit {
	// 	result = result[:len(result)-1]
	// } else {
	// 	result = result[:limit]
	// }

	// switch limit {
	// case 0:
	// 	result = result[:1]
	// case -1:
	// 	result = result[:len(result)-1]
	// default:
	// 	result = result[:limit]
	// }
	// return result
}

// func genSplit(s, sep string, sepSave, n int) []string {
// 	if n == 0 {
// 		return nil
// 	}
// 	if sep == "" {
// 		return explode(s, n)
// 	}
// 	// if n < 0 {
// 	// 	n = Count(s, sep) + 1
// 	// }

// 	a := make([]string, n)
// 	n--
// 	i := 0
// 	for i < n {
// 		m := Index(s, sep)
// 		if m < 0 {
// 			break
// 		}
// 		a[i] = s[:m+sepSave]
// 		s = s[m+len(sep):]
// 		i++
// 	}
// 	a[i] = s
// 	return a[:i+1]
// }

// func SplitAfterN(s, sep string, n int) []string {
// 	return genSplit(s, sep, len(sep), n)
// }
