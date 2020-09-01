package main

import (
	"flag"
	"fmt"
	"image/png"
	"os"

	"github.com/kshiva1126/goybox/toys/mkimg"
)

func main() {
	var (
		width      = flag.Int("w", 200, "Set the width of the image.")
		height     = flag.Int("h", 200, "Set the height of the image.")
		imageColor = flag.String("c", "red", "Set the color of the image.")
		filename   = flag.String("n", "sampleImage.png", "Set the file name of the image.")
		text       = flag.String("t", "", "Set the text to be added to the image.")
		fontColor  = flag.String("fc", "white", "Set the color of the font.")
		fontsize   = flag.Int("fs", 24, "Set the font size.")
	)
	flag.Parse()

	c, err := mkimg.NewCreator(mkimg.Params{
		Width:          width,
		Height:         height,
		ImageColorname: imageColor,
		Text:           text,
		Fontsize:       fontsize,
		FontColorname:  fontColor,
	})
	if err != nil {
		fmt.Println(err)
		flag.Usage()
		os.Exit(2)
	}
	img := c.Create()
	file, err := os.Create(*filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	if err = png.Encode(file, img); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
