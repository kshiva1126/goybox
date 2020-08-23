package mkimg_test

import (
	"image"
	_ "image/png"
	"os"
	"testing"

	"github.com/kshiva1126/goybox/toys/mkimg"
)

func Test_CreateImage(t *testing.T) {
	type Input struct {
		height, width       int
		colorname, filename string
	}

	type Expected struct {
		height, width int
	}
	tests := []struct {
		name     string
		input    Input
		expected Expected
	}{
		{
			name: "1",
			input: Input{
				height:    100,
				width:     100,
				colorname: "red",
				filename:  "test.png",
			},
			expected: Expected{
				height: 100,
				width:  100,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			_, err := mkimg.CreateImage(&tt.input.height,
				&tt.input.width, &tt.input.colorname, &tt.input.filename)
			if err != nil {
				t.Errorf("Expected return of nil, but got %v", err)
			}

			reader, err := os.Open(tt.input.filename)
			if err != nil {
				t.Errorf("Expected return of nil, but got %v", err)
			}

			config, _, err := image.DecodeConfig(reader)
			if err != nil {
				t.Errorf("Expected return of nil, but got %v", err)
			}

			if config.Height != tt.expected.height {
				t.Errorf("Expected return of %v, but got %v", tt.expected.height, config.Height)
			}

			if err := os.Remove(reader.Name()); err != nil {
				t.Errorf("Expected return of nil, but got %v", err)
			}
		})
	}
}
