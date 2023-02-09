package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kyokomi/emoji/v2"
	"github.com/urfave/cli/v2"

	"golang.design/x/clipboard"
)

func main() {
	useClipboard := false

	app := &cli.App{
		Name:    "emojiterm - text to emoji converter - Use underscore instead of space in emoji names",
		Usage:   "emojiterm \"<space separeted list of tokens>\" | Example: emojiterm \"sun wine_glass beer)\"\nNames are based on this file: https://raw.githubusercontent.com/iamcal/emoji-data/master/emoji.json",
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
				if err := writeToClipboard(s); err != nil {
					fmt.Println("Clipboard not supported")
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

func writeToClipboard(s string) error {
	if err := clipboard.Init(); err != nil {
		return err
	}
	clipboard.Write(clipboard.FmtText, []byte(s))
	return nil
}
