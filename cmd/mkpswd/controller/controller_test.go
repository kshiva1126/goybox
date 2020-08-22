package controller_test

import (
	"testing"

	"github.com/kshiva1126/goybox/cmd/mkpswd/controller"
)

func Test_CreatePassword(t *testing.T) {
	var nc int = 8
	var np int = 1
	tests := []struct {
		name  string
		input []string
	}{
		{
			"1",
			[]string{"n"},
		},
		{
			"2",
			[]string{"l", "n"},
		},
		{
			"3",
			[]string{"l", "n", "s"},
		},
		{
			"4",
			[]string{"l", "u", "n", "s"},
		},
		{
			"5",
			[]string{"l", "u", "n", "s", "c"},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			outputs, err := controller.CreatePassword(tt.input, &nc, &np)
			if err != nil {
				t.Errorf("Expected return of nil, but got %v", err)
			}
			if len(outputs) != np {
				t.Errorf("Expected length of outputs is %v, but got %v", np, len(outputs))
			}
			for _, output := range outputs {
				if len(output) != nc {
					t.Errorf("Expected length of output is %v, but got %v", nc, len(output))
				}
			}
		})
	}
}
