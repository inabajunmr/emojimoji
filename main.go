package main

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png"
	"io"
	"os"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/gobold"
	"golang.org/x/image/math/fixed"
)

const (
	exitSuccess = 0
	exitFailure = 1
)

const (
	imagePath     = "./noto-emoji/png/128"
	fontSize      = 64  // point
	imageWidth    = 640 // pixel
	imageHeight   = 120 // pixel
	textTopMargin = 80  // fixed.I
)

func main() {
	// TrueType フォントの読み込み
	ft, err := truetype.Parse(gobold.TTF)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(exitFailure)
	}

	opt := truetype.Options{
		Size:              fontSize,
		DPI:               0,
		Hinting:           0,
		GlyphCacheEntries: 0,
		SubPixelsX:        0,
		SubPixelsY:        0,
	}

	img := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))
	face := truetype.NewFace(ft, &opt)

	dr := &font.Drawer{
		Dst:  img,
		Src:  image.White,
		Face: face,
		Dot:  fixed.Point26_6{},
	}

	text := "Hello, world! 👋"

	// TODO 改行の数をカウントし、初期位置を決める
	// 画像の高さ / 改行の数でいいはず

	// 描画の初期位置
	dr.Dot.X = 0
	dr.Dot.Y = fixed.I(textTopMargin) // TODO

	// 一文字ずつ描画していく
	for _, r := range text {
		// 初期位置をリセット
		// Dot.Yを既存+テキストの高さにする
		dr.DrawString(string(r))
	}

	// JPEG に変換して stdout に出力
	buf := &bytes.Buffer{}
	err = jpeg.Encode(buf, img, nil)
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
