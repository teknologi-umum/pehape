package pehape_test

import (
	"bytes"
	"fmt"
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
		want := []byte("🐘")
		
		var got bytes.Buffer
		got.WriteString(pehape.Chr(240))
		got.WriteString(pehape.Chr(159))
		got.WriteString(pehape.Chr(144))
		got.WriteString(pehape.Chr(152))

		if !bytes.EqualFold(want, got.Bytes()) {
			t.Errorf("want: %q, got: %q", want, got.Bytes())
		}
	})
}
