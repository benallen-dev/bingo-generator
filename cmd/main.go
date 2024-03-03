package main

import (
	"fmt"

	"github.com/benallen-dev/bingo-generator/pkg/bingo"
)

// This isn't exacty the best thing I've ever written but oh well it works

const (
	CARDS_PER_ROUND int = 9
)

var WINNERS = [][]int{
	{12, 15, 18, 21, 25, 26},
	{12, 15, 18, 21, 25, 26},
	{12, 15, 18, 21, 25, 26},
	{12, 15, 18, 21, 25, 26},
}

func main() {
	fmt.Println("Welcome to the Bingo Generator!")
	fmt.Println()

	// Generate the bingo cards
	rounds := bingo.Generate(WINNERS, CARDS_PER_ROUND)

	for roundIdx, round := range rounds {
		fmt.Println("Round", roundIdx+1)
		fmt.Print(round.Display())
	}
}
