package main

import (
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
	t.Run("CanImplodeArrayOfInt", func(t *testing.T) {
		var array = []int{1, 2, 3, 4, 5}
		if res := Implode(array, " "); res != "1 2 3 4 5" {
			t.Errorf("Expected 1 2 3 4 5, got %s", res)
		}
		if res := Implode(array, ":. "); res != "1:. 2:. 3:. 4:. 5" {
			t.Errorf("Expected 1:. 2:. 3:. 4:. 5, got %s", res)
		}
	})
	t.Run("CanImplodeEMptyArray", func(t *testing.T) {
		var arr = []string{}
		if res := Implode(arr, " "); res != "" {
			t.Errorf("Expected empty string, got %s", res)
		}
	})

}