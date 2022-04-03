package main

import (
	"strings"
	"testing"
)

//function to test explode function
func TestExplode(t *testing.T) {

	/*
	 * from
	 * https://www.php.net/manual/en/function.explode.php#refsect1-function.explode-examples
	 */

	var str string = "piece1 piece2 piece3 piece4 piece5 piece6"
	six := Explode(" ", str, 6)
	if len(six) != 6 {
		t.Errorf("Expected 1, got %d", len(six))
	}

	if six[:1][0] != "piece1" {
		t.Errorf("Expected piece1, got %s", six[0])
	}

	str = "Hello pehape world"

	//nil
	if res := Explode(" ", str); res[0] != []string{"Hello", "pehape", "world"}[0] {
		t.Errorf("Expected Hello, got %s", res[0])
	}
	if res := Explode(" ", str); len(res) != 3 {
		t.Errorf("Expected 3, got %d", len(res))
	}

	//0
	if res := Explode(" ", str, 0); res[:1][0] != "Hello" {
		t.Errorf("Expected Hello, got %s", res[0])
	}

	if res := Explode(" ", str, 0); len(res) != 1 {
		t.Errorf("Expected 1, got %d", len(res))
	}

	//1
	if res := Explode(" ", str, 1); res[:1][0] != "Hello" {
		t.Errorf("Expected 1, got %d", len(res))
	}
	if res := Explode(" ", str, 1); len(res) != 1 {
		t.Errorf("Expected 1, got %d", len(res))
	}

	//5
	if res := Explode(" ", str, 5); res[:1][0] != "Hello" {
		t.Errorf("Expected 1, got %d", len(res))
	}
	if res := Explode(" ", str, 5); len(res) != 3 {
		t.Errorf("Expected 5, got %d", len(res))
	}

	//-1
	if res := Explode(" ", str, -1); res[:1][0] != []string{"Hello", "pehape"}[0] {
		t.Errorf("Expected 1, got %d", len(res))
	}
	if res := Explode(" ", str, -1); len(res) != 2 {
		t.Errorf("Expected 5, got %d", len(res))
	}

	//-3
	if res := Explode(" ", str, -3); strings.Compare(res[:1][0], "") != 1 {
		t.Errorf("Expected nil, got %d", len(res))
	}
	if res := Explode(" ", str, -3); len(res) != 0 {
		t.Errorf("Expected 0, got %d", len(res))
	}

	//-5
	if res := Explode(" ", str, 5); res[:1][0] != []string{"Hello", "pehape", "world"}[0] {
		t.Errorf("Expected 1, got %d", len(res))
	}
	if res := Explode(" ", str, 5); len(res) != 3 {
		t.Errorf("Expected 5, got %d", len(res))
	}

}
