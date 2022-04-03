package main

import (
	"fmt"
	"testing"
)

func exp(s string, opts ...string) string {
	var option string
	if len(opts) > 0 {
		option = opts[0]
	}

	return s + option
}

//function to test implode function
func TestImplode(t *testing.T) {
	t.Run("CanImplodeArrayOfString", func(t *testing.T) {
		var array = []string{"Hello", "world"}

		if res := Implode(array, " "); res != "Hello world" {
			t.Errorf("Expected Hello world, got %s", res)
		}
		if res := Implode(array, "  "); res != "Hello  world" {
			t.Errorf("Expected Hello world, got %s", res)
		}
		if res := Implode(array, ""); res != "Helloworld" {
			t.Errorf("Expected Hello world, got %s", res)
		}
		if res := Implode(array); res != "Helloworld" {
			t.Errorf("Expected Hello world, got %s", res)
		}
		if res := Implode(array, " wkwkwk "); res != "Hello wkwkwk world" {
			t.Errorf("Expected Hello world, got %s", res)
		}
		if res := Implode(array, "\n"); res != "Hello\nworld" {
			t.Errorf("Expected Hello world, got %s", res)
		}
	})
	firstInts, secondInts := splitAnySlice([]int{0, 1, 2, 3})
	fmt.Println(firstInts, secondInts)
	// prints [0 1] [2 3]

	firstStrings, secondStrings := splitAnySlice([]string{"zero", "one", "two", "three"})
	fmt.Println(firstStrings, secondStrings)

	var arr []string = []string{"Hello", "world"}
	fmt.Println(ImpGen(arr, " "))
	var arr2 []int = []int{0, 1, 2, 3}
	fmt.Println(ImpGen(arr2, " "))
}
