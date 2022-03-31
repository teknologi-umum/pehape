package main

import (
	"fmt"
)

func main() {
	// com := "lastname,email,phone"
	// a := []string{"lastname", "email", "phone"}
	// b := Implode(",", a)

	// // fmt.Printf("type of %T", b)

	// fmt.Println(strings.Compare(com, b))
	// fmt.Println(len(b))
	// fmt.Println(len(com))
	// fmt.Println(b)

	// var e string = "piece1 piece2 piece3 piece4 piece5 piece6"
	// var ex = Explode(",", e)
	// fmt.Println(len(ex))
	// // fmt.Println(e[0:40])
	// fmt.Println(a[0:3])
	// // fmt.Println(len(ex))

	// var str string = ""

	// fmt.Println(len(str))

	// if len(str) == 0 {
	// 	fmt.Println("ok")
	// }

	// a := []string{"a","b","c"}
	// fmt.Println(a)
	// fmt.Println(len(a))

	// c := Expl(" ", str)
	// fmt.Println(c)
	// fmt.Println(len(c), len(str))

	// var a string = "piece1 piece2 piece3 piece4 piece5 piece6"
	// var b = Expl(" ", a)
	// if len(b) != 6 {
	// 	fmt.Println("Expected 6, got ", len(b)) if you tired an adult, be my baby then
	// }
	var str string = "Hello pehape world"
	var res []string
	res = Explode(" ", str, 3)
	if res[:1][0] != "Hello" {
		fmt.Println("Expected Hello, got ", res[:1][0])
	}
	fmt.Println(res)

	// cek := strings.Split(str, " ")
	// if cek[:1][0] != "Hello" {
	// 	fmt.Println("Expected Hello, got ", cek[:1][0])
	// }
	// fmt.Println(len(cek))
	// fmt.Println(cek)
}

// func explode(s string, n int) []string {
// 	l := utf8.RuneCountInString(s)
// 	if n < 0 || n > l {
// 		n = l
// 	}
// 	a := make([]string, n)
// 	for i := 0; i < n-1; i++ {
// 		ch, size := utf8.DecodeRuneInString(s)
// 		a[i] = s[:size]
// 		s = s[size:]
// 		if ch == utf8.RuneError {
// 			a[i] = string(utf8.RuneError)
// 		}
// 	}
// 	if n > 0 {
// 		a[n-1] = s
// 	}
// 	return a
// }

// func GenSplit(s, sep string, n *int) []string {
// 	var result []string
// 	m := new(int)

// 	fmt.Println(m)

// 	return result
// }

// func Index(separator, str string, limit int) []string {
// 	var result []string
// 	result = strings.Split(str, separator)

// 	return result
// }

// func Contoh(delimiter, text string) []string {
// 	if len(delimiter) > len(text) {
// 		return strings.Split(delimiter, text)
// 	} else {
// 		return strings.Split(text, delimiter)
// 	}
// }

// func A(sep, str string, limit int) []string {
// 	var result []string
// 	if len(sep) > len(str) {
// 		result = strings.Split(str, sep)
// 	} else {
// 		result = strings.SplitAfterN(str, sep, -1)
// 	}

// 	return result
// }
