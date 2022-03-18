package main

import (
	"strings"
	"testing"
)

//function to test implode function
func TestImplode(t *testing.T) {
	var join = []string{"1", "2", "3"}
	var join2 = []string{"a"}
	var join3 = []string{}

	var array = []string{"Hello", "world"}

	a1 := Implode(",", join)
	a2 := Implode(",", join2)
	a3 := Implode(",", join3)
	arr := Implode(" ", array)
	arrWithNoSpace := Implode("", array)
	wkwk := Implode(" wkwkwk ", array)
	enter := Implode("\n", array)

	//check expected result and compare with actual result
	if a1 != "1,2,3" && strings.Compare(a1, "1,2,3") == 1 {
		t.Errorf("a1 is : %s\n", a1)
	}

	if a2 != "a" && strings.Compare(a2, "a") == 1 {
		t.Errorf("a2 is : %s\n", a2)
	}

	if a3 != "" && strings.Compare(a3, "") == 1 {
		t.Errorf("a3 is : %s\n", a3)
	}
	if arr != "Hello world" && strings.Compare(arr, "Hello world") == 1 {
		t.Errorf("arr is : %s\n", arr)
	}

	if arrWithNoSpace != "Helloworld" && strings.Compare(arrWithNoSpace, "Hello world") == 1 {
		t.Errorf("arrWithNoSpace is : %s\n", arrWithNoSpace)
	}

	if wkwk != "Hello wkwkwk world" && strings.Compare(wkwk, "Hello wkwkwk world") == 1 {
		t.Errorf("wkwk is : %s\n", wkwk)
	}
	
	if enter != "Hello\nworld" && strings.Compare(enter, "Hello\nworld") == 1 {
		t.Errorf("enter is : %s\n", enter)
	}
}
