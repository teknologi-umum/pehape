package main

import (
	"fmt"
	"strings"
)

func main() {
	com := "lastname,email,phone"
	a := []string{"lastname", "email", "phone"}
	b := Implode(",", a)

	// fmt.Printf("type of %T", b)

	fmt.Println(strings.Compare(com, b))
	fmt.Println(len(b))
	fmt.Println(len(com))
	fmt.Println(b)

	var e string = "piece1 piece2 piece3 piece4 piece5 piece6"
	var ex = Explode(",", e)
	fmt.Println(len(ex))
	// fmt.Println(e[0:40])
	fmt.Println(a[0:3])
	// fmt.Println(len(ex))

	var str string = ""

	fmt.Println(len(str))

	if len(str) == 0 {
		fmt.Println("ok")
	}

}
