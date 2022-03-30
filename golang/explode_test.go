package main

import (
	"testing"
)

//function to test explode function
func TestExplode(t *testing.T) {

	/*
	 * from 
	 * https://www.php.net/manual/en/function.explode.php#refsect1-function.explode-examples	 
	*/

	var str string = "piece1 piece2 piece3 piece4 piece5 piece6"
	six := Exp(" ", str, 0)
	if len(six) != 1 {
		t.Errorf("Expected 1, got %d", len(six))
	}

	str = "Hello pehape world"

	res := Exp(" ", str, 3)
	if len(res) != 3 {
		t.Errorf("Expected 3, got %d", len(res))
	}

}
