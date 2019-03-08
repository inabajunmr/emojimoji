package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/draw"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

const (
	exitSuccess = 0
	exitFailure = 1
)

const (
	imagePath   = "./noto-emoji/png/128"
	imageHeight = 128
)

// TODO error handling
// TODO text and bg Color
func generate(text string) *image.RGBA {
	b, err := ioutil.ReadFile("./GenShinGothic-Bold.ttf")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ft, err := truetype.Parse(b)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(exitFailure)
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

	blue := color.RGBA{0, 0, 255, 255} // TODO color
	draw.Draw(img, img.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src)

	face := truetype.NewFace(ft, &opt)

	dr := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(color.RGBA{255, 255, 255, 255}),
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

	return imgDst
}

func main() {

	input := "かたい/n意思"

	buf := &bytes.Buffer{}
	err := jpeg.Encode(buf, generate(input), nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(exitFailure)
	}

	_, err = io.Copy(os.Stdout, buf)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(exitFailure)
	}
}
