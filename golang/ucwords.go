package main

import "strings"

func Ucwords(str string) string {
	return strings.Title(strings.ToLower(str))
}
