package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kyokomi/emoji/v2"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "emojiterm - text to emoji converter - Use underscore instead of space in emoji names",
		Usage: "emojiterm \"<space separeted list of tokens>\" | Example: emojiterm \"sun wine_glass beer)\"\nNames are based on this file: https://raw.githubusercontent.com/iamcal/emoji-data/master/emoji.json",
		Action: func(cCtx *cli.Context) error {
			var s string
			for _, arg := range cCtx.Args().Slice() {
				s = s + emoji.Sprintf(" :%s:", arg)
			}
			fmt.Println(s)
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
