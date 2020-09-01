package mkimg_test

import (
	"image/color"
	"testing"

	"github.com/kshiva1126/goybox/toys/mkimg"
)

func Test_Create(t *testing.T) {
	type Input struct {
		width, height, fontsize             int
		imageColorname, fontColorname, text string
	}

	type Expected struct {
		height, width int
		color         color.RGBA
	}

	tests := []struct {
		name     string
		input    Input
		expected Expected
	}{
		{
			name: "1",
			input: Input{
				width:          50,
				height:         100,
				imageColorname: "red",
				text:           "",
				fontsize:       0,
				fontColorname:  "",
			},
			expected: Expected{
				width:  50,
				height: 100,
				color:  color.RGBA{0xff, 0x00, 0x00, 0xff},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			c, err := mkimg.NewCreator(mkimg.Params{
				Width:          &tt.input.width,
				Height:         &tt.input.height,
				ImageColorname: &tt.input.imageColorname,
				Text:           &tt.input.text,
				Fontsize:       &tt.input.fontsize,
				FontColorname:  &tt.input.fontColorname,
			})
			if err != nil {
				t.Errorf("Expected return of nil, but got %v", err)
			}
			img := c.Create()
			if img.Bounds().Dx() != tt.expected.width {
				t.Errorf("Expected return of %v, but got %v",
					tt.expected.width, img.Bounds().Dx())
			}
			if img.Bounds().Dy() != tt.expected.height {
				t.Errorf("Expected return of %v, but got %v",
					tt.expected.height, img.Bounds().Dy())
			}
			if img.At(0, 0) != tt.expected.color {
				t.Errorf("Expected return of %v, but got %v",
					tt.expected.color, img.At(0, 0))
			}
		})
	}
}
