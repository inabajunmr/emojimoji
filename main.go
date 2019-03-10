package main

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"io"
	"os"

	"github.com/inabajunmr/emosh"
	gocolor "github.com/inabajunmr/gocolors"
)

func main() {

	input := "GOOD/nBYE"
	fcolor := gocolor.Of(gocolor.Red, 255)
	bcolor := gocolor.Of(gocolor.Yellow, 255)

	emoji, _ := emosh.GenerateEmoji(input, fcolor, bcolor)

	buf := &bytes.Buffer{}
	err := jpeg.Encode(buf, emoji, nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	_, err = io.Copy(os.Stdout, buf)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
