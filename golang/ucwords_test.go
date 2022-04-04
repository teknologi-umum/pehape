package main

import "testing"

func TestUcwords(t *testing.T) {
	/*
	 * https://www.php.net/manual/en/function.ucwords.php
	 */
	t.Run("CanUcwords", func(t *testing.T) {
		var str = "kata siapa SSA asAS"
		if res := Ucwords(str); res != "Kata Siapa Ssa Asas" {
			t.Errorf("Expected Kata Siapa Ssa Asas, got %s", res)
		}
		if res := Ucwords(""); res != "" {
			t.Errorf("Expected empty string, got %s", res)
		}
		if res := Ucwords("hello world"); res != "Hello World" {
			t.Errorf("Expected Hello World, got %s", res)
		}
		if res := Ucwords("HELLO WORDL"); res != "Hello Wordl" {
			t.Errorf("Expected Hello Wordl, got %s", res)
		}
	})
}
