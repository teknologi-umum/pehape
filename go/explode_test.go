package pehape_test

import (
	"strings"
	"testing"

	PHP "github.com/teknologi-umum/pehape/go"
)

//function to test explode function
func TestExplode(t *testing.T) {

	/*
	 * from
	 * https://www.PHP.net/manual/en/function.explode.PHP#refsect1-function.explode-examples
	 */

	t.Run("Can Explode Space Separated String By Space", func(t *testing.T) {
		str := "Hello pehape world"
		
		//nil
		if res, _ := PHP.Explode(" ", str); res[0] != []string{"Hello", "pehape", "world"}[0] {
			t.Errorf("Expected Hello, got %s", res[0])
		}
		if res, _ := PHP.Explode(" ", str); len(res) != 3 {
			t.Errorf("Expected 3, got %d", len(res))
		}

		//0
		if res, _ := PHP.Explode(" ", str, 0); res[:1][0] != "Hello" {
			t.Errorf("Expected Hello, got %s", res[0])
		}

		if res, _ := PHP.Explode(" ", str, 0); len(res) != 1 {
			t.Errorf("Expected 1, got %d", len(res))
		}

		//1
		if res, _ := PHP.Explode(" ", str, 1); res[:1][0] != "Hello" {
			t.Errorf("Expected 1, got %d", len(res))
		}
		if res, _ := PHP.Explode(" ", str, 1); len(res) != 1 {
			t.Errorf("Expected 1, got %d", len(res))
		}

		//5
		if res, _ := PHP.Explode(" ", str, 5); res[:1][0] != "Hello" {
			t.Errorf("Expected 1, got %d", len(res))
		}
		if res, _ := PHP.Explode(" ", str, 5); len(res) != 3 {
			t.Errorf("Expected 5, got %d", len(res))
		}

		//-1
		if res, _ := PHP.Explode(" ", str, -1); res[:1][0] != []string{"Hello", "pehape"}[0] {
			t.Errorf("Expected 1, got %d", len(res))
		}
		if res, _ := PHP.Explode(" ", str, -1); len(res) != 2 {
			t.Errorf("Expected 5, got %d", len(res))
		}

		//-3
		if res, _ := PHP.Explode(" ", str, -3); strings.Compare(res[:1][0], "") != 1 {
			t.Errorf("Expected nil, got %d", len(res))
		}
		if res, _ := PHP.Explode(" ", str, -3); len(res) != 0 {
			t.Errorf("Expected 0, got %d", len(res))
		}

		//-5
		if res, _ := PHP.Explode(" ", str, 5); res[:1][0] != []string{"Hello", "pehape", "world"}[0] {
			t.Errorf("Expected 1, got %d", len(res))
		}
		if res, _ := PHP.Explode(" ", str, 5); len(res) != 3 {
			t.Errorf("Expected 5, got %d", len(res))
		}
	})

	t.Run("Can Explode Double Space Separated String By Space", func(t *testing.T) {
		var str string = "Hello  pehape  world"
		if res, _ := PHP.Explode(" ", str); res[0] != []string{"Hello", "", "pehape", "", "world"}[0] {
			t.Errorf("Expected Hello, got %s", res[0])
		}
	})

	t.Run("Can Explode Space Separated String By Double Space", func(t *testing.T) {
		var str string = "Hello pehape world"
		if res, _ := PHP.Explode("  ", str); res[0] != str {
			t.Errorf("Expected Hello, got %s", res[0])
		}
	})

	t.Run("Explode Space Separated String By Empty String Throws Exception", func(t *testing.T) {
		var str string = "Hello pehape world"
		_, err := PHP.Explode("", str)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if err.Error() != "Separator cannot be empty" {
			t.Errorf("Expected error, got %s", err.Error())
		}
	})

}
