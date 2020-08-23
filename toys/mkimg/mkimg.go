package mkimg

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strings"

	"golang.org/x/image/colornames"
)

var (
	x = 0
	y = 0
)

func CreateImage(height, width *int, colorname, filename *string) (*os.File, error) {
	colorRGBA, err := getColorRGBA(*colorname)
	if err != nil {
		return nil, err
	}

	img := image.NewRGBA(image.Rect(x, y, *width, *height))

	for i := img.Rect.Min.Y; i < img.Rect.Max.Y; i++ {
		for j := img.Rect.Min.X; j < img.Rect.Max.X; j++ {
			img.Set(j, i, colorRGBA)
		}
	}

	if !strings.HasSuffix(*filename, ".png") {
		extension := ".png"
		*filename += extension
	}
	file, err := os.Create(*filename)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err = png.Encode(file, img); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return file, nil
}

func getColorRGBA(colorname string) (color.Color, error) {
	colorname = strings.ToLower(string(rune(colorname[0]))) + colorname[1:]
	colorNameMap := colornames.Map
	v, ok := colorNameMap[colorname]
	if !ok {
		return nil, fmt.Errorf("invalid value \"%v\" for flag -c", colorname)
	}

	return v, nil
}
