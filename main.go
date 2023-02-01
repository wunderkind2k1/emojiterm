package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kyokomi/emoji/v2"
	"github.com/urfave/cli/v2"
)

func main() {
	//TODO: create a solution for space separated emojis like :wine glass:
	app := &cli.App{
		Name:  "emojiterm",
		Usage: "convert text to emojis and output them on stdout",
		Action: func(cCtx *cli.Context) error {
			var s string
			for _, arg := range cCtx.Args().Slice() {
				s = s + emoji.Sprintf(" :%s:", arg)
			}
			//emoji.Println(fmt.Sprintf(":%s:", cCtx.Args().Slice()[]))
			//fmt.Printf("Hello %s\n", cCtx.Args().Get(0))
			fmt.Println(s)
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
