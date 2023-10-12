package pehape

import "strings"

func StrStartsWith(haystack, needle string) bool {
	return strings.HasPrefix(haystack, needle)
}