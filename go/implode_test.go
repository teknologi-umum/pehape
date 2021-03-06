package pehape_test

import (
	"testing"

	PHP "github.com/teknologi-umum/pehape/go"
)

//function to test implode function
func TestImplode(t *testing.T) {
	t.Run("Can Implode Array Of String", func(t *testing.T) {
		var array = []string{"Hello", "world"}

		if res := PHP.Implode(array, " "); res != "Hello world" {
			t.Errorf("Expected Hello world, got %s", res)
		}
		if res := PHP.Implode(array, "  "); res != "Hello  world" {
			t.Errorf("Expected Hello world, got %s", res)
		}
		if res := PHP.Implode(array, ""); res != "Helloworld" {
			t.Errorf("Expected Hello world, got %s", res)
		}
		if res := PHP.Implode(array); res != "Helloworld" {
			t.Errorf("Expected Hello world, got %s", res)
		}
		if res := PHP.Implode(array, " wkwkwk "); res != "Hello wkwkwk world" {
			t.Errorf("Expected Hello world, got %s", res)
		}
		if res := PHP.Implode(array, "\n"); res != "Hello\nworld" {
			t.Errorf("Expected Hello world, got %s", res)
		}
	})
	t.Run("Can Implode Array Of Int", func(t *testing.T) {
		var array = []int{1, 2, 3, 4, 5}
		if res := PHP.Implode(array, " "); res != "1 2 3 4 5" {
			t.Errorf("Expected 1 2 3 4 5, got %s", res)
		}
		if res := PHP.Implode(array, ":. "); res != "1:. 2:. 3:. 4:. 5" {
			t.Errorf("Expected 1:. 2:. 3:. 4:. 5, got %s", res)
		}
	})
	t.Run("Can Implode EmptyArray", func(t *testing.T) {
		var arr = []string{}
		if res := PHP.Implode(arr, " "); res != "" {
			t.Errorf("Expected empty string, got %s", res)
		}
	})

}
