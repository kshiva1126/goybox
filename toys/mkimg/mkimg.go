package mkimg

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"os"
	"strings"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
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
	addLabel(img, 40, *height/2, "私はGoを書きます。")

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

func addLabel(img *image.RGBA, x, y int, label string) {
	b, err := ioutil.ReadFile("../../font.ttf")
	if err != nil {
		panic(err)
	}
	tt, err := truetype.Parse(b)
	if err != nil {
		panic(err)
	}

	opt := truetype.Options{Size: 10}
	col := color.RGBA{200, 100, 0, 255}
	point := fixed.Point26_6{fixed.Int26_6(x * 64), fixed.Int26_6(y * 64)}

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: truetype.NewFace(tt, &opt),
		Dot:  point,
	}
	d.DrawString(label)
}
