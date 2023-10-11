package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var a = "a"

	var c = "🙂a"

	var b = []rune(a)

	var d = []rune(c)
	e, size := utf8.DecodeRuneInString(c)

	fmt.Println(b)
	fmt.Println(string(b))

	fmt.Println(b[0] - 32)
	fmt.Println(string(b[0] - 32))

	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(size)
}
