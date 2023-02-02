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
			s := printEmojiArgsAsString(cCtx.Args().Slice())
			fmt.Println(s)
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func printEmojiArgsAsString(args []string) string {
	s := ""
	for index, arg := range args {
		if index == 0 {
			s = s + emoji.Sprintf(":%s:", arg)
		} else {
			s = s + emoji.Sprintf(" :%s:", arg)
		}
	}
	return s
}
