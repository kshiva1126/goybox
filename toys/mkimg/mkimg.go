package mkimg

import (
	"fmt"
	"image"
	"image/color"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font"
)

var (
	x = 0
	y = 0
)

// Creator is the main interface for this package.
type Creator interface {
	Create() *image.RGBA
	SetSize(int, int)
	SetImageColor(string) error
	SetFont() error
	SetFontColor(string) error
	SetText(string)
	SetFontsize(int)
}

// Params is parameters for NewCreator function.
type Params struct {
	Width, Height, Fontsize             *int
	ImageColorname, FontColorname, Text *string
}

// NewCreator returns Creator Interface.
func NewCreator(params Params) (Creator, error) {
	c := &creator{}
	if err := c.SetImageColor(*params.ImageColorname); err != nil {
		return c, err
	}
	c.SetSize(*params.Width, *params.Height)
	if *params.Text != "" {
		if err := c.SetFont(); err != nil {
			return c, err
		}
		c.SetFontColor(*params.FontColorname)
		c.SetText(*params.Text)
		c.SetFontsize(*params.Fontsize)
	}

	return c, nil
}

type creator struct {
	Width, Height, Fontsize int
	Text                    string
	ImageColor, FontColor   color.Color
	Font                    *truetype.Font
}

// Create returns the image
func (c *creator) Create() *image.RGBA {
	img := image.NewRGBA(image.Rect(x, y, c.Width, c.Height))

	for i := img.Rect.Min.Y; i < img.Rect.Max.Y; i++ {
		for j := img.Rect.Min.X; j < img.Rect.Max.X; j++ {
			img.Set(j, i, c.ImageColor)
		}
	}

	if c.Text != "" {
		c.addLabel(img)
	}

	return img
}

// SetImageColor set the color of image.
func (c *creator) SetImageColor(colorname string) error {
	rgba, err := colorRGBA(colorname)
	if err != nil {
		return err
	}
	c.ImageColor = rgba
	return nil
}

// SetFontColor set the font of image.
func (c *creator) SetFontColor(colorname string) error {
	rgba, err := colorRGBA(colorname)
	if err != nil {
		return err
	}
	c.FontColor = rgba
	return nil
}

func colorRGBA(colorname string) (color.Color, error) {
	colornameMap := colornames.Map
	v, ok := colornameMap[colorname]
	if ok == false {
		return nil, fmt.Errorf("invalid value \"%v\"", colorname)
	}

	return v, nil
}

// SetSize set the width and height.
func (c *creator) SetSize(width, height int) {
	c.Width = width
	c.Height = height
}

// SetFont set the font
func (c *creator) SetFont() error {
	tt, err := truetype.Parse(TTF)
	if err != nil {
		return err
	}
	c.Font = tt
	return nil
}

// SetText set the text.
func (c *creator) SetText(text string) {
	c.Text = text
}

// SetFontsize set the fontsize.
func (c *creator) SetFontsize(fontsize int) {
	c.Fontsize = fontsize
}

func (c *creator) addLabel(img *image.RGBA) {
	opt := truetype.Options{Size: float64(c.Fontsize)}
	textWidth := c.calcTextWidth()
	point := freetype.Pt((c.Width-textWidth)/2, (c.Height)/2)
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(c.FontColor),
		Face: truetype.NewFace(c.Font, &opt),
		Dot:  point,
	}
	d.DrawString(c.Text)
}

func (c *creator) calcTextWidth() int {
	var face font.Face
	var textWidth int
	opts := truetype.Options{}
	opts.Size = float64(c.Fontsize)
	face = truetype.NewFace(c.Font, &opts)
	for _, x := range c.Text {
		awidth, ok := face.GlyphAdvance(rune(x))
		if ok == false {
			return textWidth
		}
		textWidth += int(float64(awidth) / 64)
	}
	return textWidth
}
