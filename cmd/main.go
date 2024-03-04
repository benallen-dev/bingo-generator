package main

import (
	"fmt"
	"os"

	"github.com/benallen-dev/bingo-generator/pkg/bingo"
	"github.com/benallen-dev/bingo-generator/pkg/image"
)

// This isn't exacty the best thing I've ever written but oh well it works

const (
	CARDS_PER_ROUND int = 50
)

var WINNERS = [][]int{
	{12, 15, 23, 28, 29},
	{17, 20, 25, 28, 29},
	{15, 21, 24, 27, 28},
	{14, 18, 26, 30, 31},
}

func main() {
	fmt.Println("Welcome to the Bingo Generator!")
	fmt.Println()

	// Generate the bingo cards
	rounds := bingo.Generate("./assets", WINNERS, CARDS_PER_ROUND)

	// Create output dir if it doesn't exist
	err := os.Mkdir("output", 0755)
	if err != nil {
		fmt.Println("Error creating output directory:", err)
	}

	// Generate the images
	for roundNumber, round := range rounds {
		for cardNumber, card := range round.WinningCards {
			fileName := fmt.Sprintf("output/round-%d-%02d-win-%d.png", roundNumber+1, cardNumber+1, card.WinsAt())
			image.DrawCard(card, fileName, roundNumber+1)
		}

		for cardNumber, card := range round.BackupCards {
			fileName := fmt.Sprintf("output/round-%d-%02d-no-win-%d.png", roundNumber+1, cardNumber+6, card.WinsAt())
			image.DrawCard(card, fileName, roundNumber+1)
		}
	}

}
