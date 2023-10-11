package pehape

import (
	"fmt"
	"testing"
)

func TestStrShuffle(t *testing.T) {
	tests := []struct {
		name           string
		expectedRandom bool
		given          string
	}{
		{"test 1", true, "abcdefg"},
		{"test 2", true, "There are many variations of passages of Lorem Ipsum available, but the majority have suffered alteration in some form, by injected humour, or randomised words which don't look even slightly believable. If you are going to use a passage of Lorem Ipsum, you need to be sure there isn't anything embarrassing hidden in the middle of text. All the Lorem Ipsum generators on the Internet tend to repeat predefined chunks as necessary, making this the first true generator on the Internet. It uses a dictionary of over 200 Latin words, combined with a handful of model sentence structures, to generate Lorem Ipsum which looks reasonable. The generated Lorem Ipsum is therefore always free from repetition, injected humour, or non-characteristic words etc."},
		{"not random 1", false, "a"},
		{"not random 2", false, ""},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := StrShuffle(tt.given)
			if len(actual) != len(tt.given) {
				t.Errorf("(%s): expected %s, actual %s", tt.given, tt.given, actual)
			}

			if !tt.expectedRandom {
				if actual != tt.given {
					t.Errorf("(%s): expected %s, actual %s", tt.given, tt.given, actual)
				}
			}

			fmt.Printf("actual: %v\n", actual)
		})
	}
}
