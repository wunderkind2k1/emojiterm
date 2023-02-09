package main

import (
	"fmt"
	"log"
	"os"

	"github.com/atotto/clipboard"
	"github.com/kyokomi/emoji/v2"
	"github.com/urfave/cli/v2"
)

func main() {
	useClipboard := false

	app := &cli.App{
		Name: "emojiterm - text to emoji converter - Use underscore instead of space in emoji names",
		Usage: "emojiterm \"<space separeted list of tokens>\" | Example: emojiterm \"sun wine_glass beer)\"\n" +
			"Names are based on this file: https://raw.githubusercontent.com/iamcal/emoji-data/master/emoji.json",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:               "clipboard",
				Aliases:            []string{"c"},
				Usage:              "write emojis to clipboard",
				DisableDefaultText: true,
				Action: func(cCtx *cli.Context, b bool) error {
					useClipboard = b

					return nil
				},
			},
		},
		Action: func(cCtx *cli.Context) error {
			s := printEmojiArgsAsString(cCtx.Args().Slice())
			fmt.Println(s)

			if useClipboard {
				if err := clipboard.WriteAll(s); err != nil {
					fmt.Printf("Clipboard not supported: %v\n", err)
				}
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func printEmojiArgsAsString(args []string) string {
	emojisAsString := ""

	for index, arg := range args {
		if index == 0 {
			emojisAsString += emoji.Sprintf(":%s:", arg)
		} else {
			emojisAsString += emoji.Sprintf(" :%s:", arg)
		}
	}

	return emojisAsString
}
