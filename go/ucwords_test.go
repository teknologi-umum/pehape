package pehape_test

import (
	"testing"

	PHP "github.com/teknologi-umum/pehape/go"
)

func TestUcwords(t *testing.T) {
	/*
	 * https://www.php.net/manual/en/function.ucwords.php
	 */
	var foo = "hello world!"
	var bar = "HELLO WORLD!"

	if res := PHP.Ucwords(foo); res != "Hello World!" {
		t.Errorf("Expected Hello World!, got %s", res)
	}

	if res := PHP.Ucwords(bar); res != "Hello World!" {
		t.Errorf("Expected Hello World!, got %s", res)
	}

	if res := PHP.Ucwords(""); res != "" {
		t.Errorf("Expected empty string, got %s", res)
	}
}
