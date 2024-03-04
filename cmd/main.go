package main

import (
	"fmt"

	"github.com/benallen-dev/bingo-generator/pkg/bingo"
	"github.com/benallen-dev/bingo-generator/pkg/image"
)

// This isn't exacty the best thing I've ever written but oh well it works

const (
	CARDS_PER_ROUND int = 50
)

var WINNERS = [][]int{
	{12, 15, 23, 25, 26, 27 },
	{12, 15, 18, 21, 25, 26},
	{12, 15, 18, 21, 25, 26},
	{12, 15, 18, 21, 25, 26},
}

func main() {
	fmt.Println("Welcome to the Bingo Generator!")
	fmt.Println()

	// Generate the bingo cards
	rounds := bingo.Generate("./assets", WINNERS, CARDS_PER_ROUND)

	for _, round := range rounds {
		fmt.Println(round.Display())
	}

	testCard := rounds[0].WinningCards[0]

	// Generate the images
	fileName := "test-card.png"
	image.DrawCard(testCard, fileName)

}
