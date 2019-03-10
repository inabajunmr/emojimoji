package main

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"io"
	"log"
	"os"

	"github.com/inabajunmr/emosh/emosh"
	gocolor "github.com/inabajunmr/gocolors"
	cli "gopkg.in/urfave/cli.v1"
	"gopkg.in/urfave/cli.v1/altsrc"
)

func initApp() (*cli.App, error) {
	app := cli.NewApp()
	app.Name = "emosh"
	app.Description = "emoji from text"
	flags := []cli.Flag{
		altsrc.NewStringFlag(cli.StringFlag{Name: "text"}),
		altsrc.NewStringFlag(cli.StringFlag{Name: "bc"}),
		altsrc.NewStringFlag(cli.StringFlag{Name: "fc"}),
	}
	app.Flags = flags
	app.Action = func(c *cli.Context) error {
		fcolor, _ := gocolor.ValueOf(c.String("fc"), 255)
		bcolor, _ := gocolor.ValueOf(c.String("bc"), 255)
		text := c.String("text")

		emoji, _ := emosh.GenerateEmoji(text, bcolor, fcolor)
		buf := &bytes.Buffer{}
		jpeg.Encode(buf, emoji, nil)
		_, err := io.Copy(os.Stdout, buf)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		return nil
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	return app, nil
}

func mainInternal() int {
	app, err := initApp()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	err = app.Run(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	return 0
}

func main() {
	os.Exit(mainInternal())
}
