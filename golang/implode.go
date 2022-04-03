package main

import "fmt"

func Imp(separator string, array []string) string {
	var result string

	for i := 0; i < len(array); i++ {
		if i == 0 {
			result = array[i]
		} else {
			result = result + separator + array[i]
		}
	}
	return result
}

func Implode(array []string, separator ...string) string {
	var result string

	for i := 0; i < len(array); i++ {
		if separator == nil {
			result = result + array[i]
		} else {
			if i == 0 {
				result = array[i]
				fmt.Println("0 : ", result)
			} else {
				result = result + separator[0] + array[i]
				fmt.Println("else :", result)
			}
		}
	}
	return result
}

func splitAnySlice[T any](s []T) ([]T, []T) {
	mid := len(s) / 2
	return s[:mid], s[mid:]
}

func ImpGen[T any](array []T, separator string) string {
	// var result string

	return fmt.Sprintf("%T", array)

	// fmt.Println("array : ", array)
}
