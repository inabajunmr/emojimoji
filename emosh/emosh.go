package emosh

import (
	"fmt"
	"image"
	"image/color"
	"io/ioutil"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/draw"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
	"golang.org/x/xerrors"
)

const (
	imagePath   = "./noto-emoji/png/128"
	imageHeight = 128
)

func GenerateEmoji(text string, bcolor color.RGBA, fcolor color.RGBA) (*image.RGBA, error) {

	ft, err := loadFont()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil, xerrors.Errorf("Unexpected error. Can not load font file.", err)
	}

	lines := strings.Split(text, "/n")

	max := 0
	for _, line := range lines {
		c := utf8.RuneCountInString(line)
		if c > max {
			max = c
		}
	}

	fontSize := imageHeight / len(lines)
	imageWidth := max * fontSize

	opt := truetype.Options{
		Size:              float64(fontSize),
		DPI:               0,
		Hinting:           0,
		GlyphCacheEntries: 0,
		SubPixelsX:        0,
		SubPixelsY:        0,
	}

	img := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))
	draw.Draw(img, img.Bounds(), &image.Uniform{fcolor}, image.ZP, draw.Src)

	face := truetype.NewFace(ft, &opt)

	dr := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(bcolor),
		Face: face,
		Dot:  fixed.Point26_6{},
	}

	margin := fontSize - 10/len(lines)
	for _, line := range lines {

		dr.Dot.X = 0
		dr.Dot.Y = fixed.I(margin)

		dr.DrawString(string(line))
		margin = margin + fontSize
	}

	rctSrc := img.Bounds()
	imgDst := image.NewRGBA(image.Rect(0, 0, rctSrc.Dy(), rctSrc.Dy()))
	draw.CatmullRom.Scale(imgDst, imgDst.Bounds(), img, rctSrc, draw.Over, nil)

	return imgDst, nil
}

func loadFont() (font *truetype.Font, err error) {
	f, err := Assets.Open("/html/index.html")
	if err != nil {
		fmt.Println(err)
		return nil, xerrors.Errorf("Unexpected error. Can not load font file.", err)
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return nil, xerrors.Errorf("Unexpected error. Can not load font file.", err)
	}

	ft, err := truetype.Parse(b)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil, xerrors.Errorf("Unexpected error. Can not parse font file.", err)
	}

	return ft, err
}
