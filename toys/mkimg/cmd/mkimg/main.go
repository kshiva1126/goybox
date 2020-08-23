package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kshiva1126/goybox/toys/mkimg"
)

func main() {
	var (
		height    = flag.Int("h", 100, "Assign the image's height")
		width     = flag.Int("w", 100, "Assign the image's width")
		colorname = flag.String("c", "red", "Colorize the image")
		filename  = flag.String("n", "sampleImage.png", "Name the image")
	)
	flag.Parse()

	_, err := mkimg.CreateImage(height, width, colorname, filename)
	if err != nil {
		fmt.Println(err)
		flag.Usage()
		os.Exit(2)
	}
}
