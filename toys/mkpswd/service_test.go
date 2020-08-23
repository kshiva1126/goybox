package mkpswd_test

import (
	"testing"

	"github.com/kshiva1126/goybox/toys/mkpswd"
)

func Test_AllowValue(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected string
	}{
		{
			"1",
			[]string{"l", "u", "n", "s", "c"},
			"",
		},
		{
			"2",
			[]string{"l", "u", "n", "s", "c", "a"},
			"invalid \"a\" for flag -char",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			err := mkpswd.AllowValue(tt.input)
			if err != nil {
				if err.Error() != tt.expected {
					t.Errorf("Expected return of %v, but got %v", tt.expected, err)
				}
			}
		})
	}
}
