package pehape

import (
	"testing"
)

func TestUcwords(t *testing.T) {
	/*
	 * https://www.php.net/manual/en/function.ucwords.php
	 */
	t.Run("Ucwords Can Uppercase All The First Character Of Each Word", func(t *testing.T) {

		var foo = "hello world!"
		var bar = "HELLO WORLD!"

		if res := Ucwords(foo); res != "Hello World!" {
			t.Errorf("Expected Hello World!, got %s", res)
		}

		if res := Ucwords(bar); res != "Hello World!" {
			t.Errorf("Expected Hello World!, got %s", res)
		}

		if res := Ucwords(""); res != "" {
			t.Errorf("Expected empty string, got %s", res)
		}
	})
}
