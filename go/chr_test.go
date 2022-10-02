package pehape_test

import (
	"fmt"
	"strings"
	"testing"

	pehape "github.com/teknologi-umum/pehape/go"
)

func TestChr(t *testing.T) {
	testcases := []struct {
		input     int
		wantOuput string
	}{
		{
			input:     65,
			wantOuput: "A",
		},
		{
			input:     053,
			wantOuput: "+",
		},
		{
			input:     0x52,
			wantOuput: "R",
		},
		{
			input:     -159,
			wantOuput: "a",
		},
		{
			input:     833,
			wantOuput: "A",
		},
	}

	for i, c := range testcases {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			got := pehape.Chr(c.input)

			if c.wantOuput != got {
				t.Errorf("want: %q, got: %q", c.wantOuput, got)
			}
		})
	}

	t.Run("symbol", func(t *testing.T) {
		want := "🐘"

		got := pehape.Chr(240) + pehape.Chr(159) + pehape.Chr(144) + pehape.Chr(152)

		if !strings.EqualFold(want, got) {
			t.Errorf("want: %q, got: %q", want, got)
		}
	})
}
