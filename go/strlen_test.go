package pehape_test

import (
	"testing"

	PHP "github.com/teknologi-umum/pehape/go"
)

func TestStrlen(t *testing.T) {
	// Test empty string
	if PHP.Strlen("") != 0 {
		t.Errorf("Strlen(\"\") = %d; want 0", PHP.Strlen(""))
	}

	// Test string with one character
	if PHP.Strlen("a") != 1 {
		t.Errorf("Strlen(\"a\") = %d; want 1", PHP.Strlen("a"))
	}

	// Test string with multiple characters
	if PHP.Strlen("hello world") != 11 {
		t.Errorf("Strlen(\"hello world\") = %d; want 11", PHP.Strlen("hello world"))
	}

	// Test string with special characters
	if PHP.Strlen("こんにちは") != 15 {
		t.Errorf("Strlen(\"こんにちは\") = %d; want 15", PHP.Strlen("こんにちは"))
	}

	// Test emoji
	if PHP.Strlen("🤗🥰") != 8 {
		t.Errorf("Strlen(\"🤗🥰\") = %d; want 8", PHP.Strlen("🤗🥰"))
	}
}
