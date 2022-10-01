package pehape

import "testing"

func TestGetCapacity(t *testing.T) {
	tests := []struct {
		name     string
		capacity int
		mod      int
		lenS     int
		leng     int
	}{
		{"get mod 0", 10, 0, 100, 10},
		{"get mod 1", 12, 1, 100, 9},
		{"get mod 0", 1, 0, 10, 10},
		{"get mod 1", 2, 1, 11, 10},
		{"get mod 1", 2, 1, 11, 10},
		{"get mod 2", 2, 2, 12, 10},
		{"get mod 3", 2, 3, 13, 10},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			capacityAct, modAct := getCapacity(tt.lenS, tt.leng)
			if capacityAct != tt.capacity || modAct != tt.mod {
				t.Errorf("(%d, %d): expected (%d, %d), actual (%d, %d)", tt.lenS, tt.mod, tt.capacity, tt.mod, capacityAct, modAct)
			}
		})
	}
}

func TestStrSplit(t *testing.T) {
	tests := []struct {
		name     string
		expected []string
		s        string
		length   int
	}{
		{"with no pad \"hello\"", []string{"h", "e", "l", "l", "o"}, "hello", 0},
		{"with pad 3 \"hello\"", []string{"hel", "lo"}, "hello", 3},
		{"with pad 2 \"hello\"", []string{"he", "ll", "o"}, "hello", 2},
		{"with pad 4 \"hello\"", []string{"hell", "o"}, "hello", 4},
		{"with no pad \"Hello Friend\"", []string{"H", "e", "l", "l", "o", " ", "F", "r", "i", "e", "n", "d"}, "Hello Friend", 0},
		{"with pad 5 \"Hello Friend\"", []string{"Hello", " Frie", "nd"}, "hello", 5},
		{"with no data and pad", []string{}, "", 0},
		{"with no data with pad", []string{}, "", 5},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			var actual []string
			if tt.length == 0 {
				actual = StrSplit(tt.s)
			} else {
				actual = StrSplit(tt.s, tt.length)
			}
			if len(actual) != len(tt.expected) {
				for i, v := range actual {
					if v != tt.expected[i] {
						t.Errorf("(%s): expected %s, actual %s", tt.s, tt.expected[i], v)
					}
				}
			}
		})
	}
}
