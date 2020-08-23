package mkimg

import (
	"fmt"
	"image/color"
	"testing"
)

func Test_getColorRGBA(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected color.Color
	}{
		{
			"1",
			"red",
			color.RGBA{0xff, 0x00, 0x00, 0xff},
		},
		{
			"2",
			"blahblah",
			nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			colorRGBA, err := getColorRGBA(tt.input)
			if err != nil {
				if err.Error() != fmt.Sprintf("invalid value \"%v\" for flag -c", tt.input) {
					t.Errorf("Expected return of %v, but got %v", fmt.Sprintf("invalid value \"%v\" for flag -c", tt.input), err)
				}
			}

			if colorRGBA != tt.expected {
				t.Errorf("Expected return of %v, but got %v", tt.expected, colorRGBA)
			}
		})
	}
}
