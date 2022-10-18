package pehape

import "testing"

func TestFunction(t *testing.T) {
	tests := []struct {
		name     string
		expected string
		err      error
		num      float64
		opts     []string
	}{
		{"1000", "1000", nil, 1000.12, nil},
		{"error param > 4", "", errMust4Params, 1000.12, []string{"3", "4", "5", "5"}},
		{"param 2 not int", "", errParam2MustInt, 1000.12, []string{"f", "4", "5"}},
		{"1,000.12", "1,000.12", nil, 1000.12, []string{"2"}},
		{"1,000.12 php magic", "1,000.12", nil, 1000.12, []string{"2.4"}},
		{"1.000,12", "1.000,12", nil, 1000.12, []string{"2", ",", "."}},
		{"1 000.12", "1 000.12", nil, 1000.12, []string{"2", ".", " "}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual, err := NumberFormat(tt.num, tt.opts...)
			if actual != tt.expected && err != tt.err {
				t.Errorf("(%v, %v): expected (%s, %v), actual (%s, %v)", tt.num, tt.opts, tt.expected, tt.err, actual, err)
			}
		})
	}
}
