package main

import (
	"os"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	//can be improved: hijack stdout check emojis
	os.Args = []string{"emojiterm", "tropical_drink cup_with_straw"}
	main()
}

func TestArgsToEmojis(t *testing.T) {
	args := []string{"beer", "sushi"}
	expectBeerAndSushiEmoji := strings.TrimSpace(printEmojiArgsAsString(args))
	beerAndSushi := "🍺  🍣"

	if strings.Compare(beerAndSushi, expectBeerAndSushiEmoji) != 0 {
		t.Fatalf("ExpectedbeerAndSushiEmoji expected %v, got %v", beerAndSushi, expectBeerAndSushiEmoji)
	}
}
