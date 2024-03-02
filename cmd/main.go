package main

import (
	"fmt"
	"math/rand"
)

// This isn't exacty the best thing I've ever written but oh well it works

const (
	CARDS_PER_ROUND int = 50
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

	rounds := readInput()

	for roundIdx, round := range rounds {
		// fmt.Println()
		// fmt.Println("Round", roundIdx+1)

		// Generate the winning cards
	//	fmt.Println("=========================================")

		winningCards := []*Card{}

		// For each winning number in the round generate a card
		for _, winNumber := range WINNERS[roundIdx] {
			winningSong := round[winNumber-1]

			newCard := NewCard()
			newCard.AddSong(winningSong)

			// While there are still empty spots on the card,
			for newCard.Complete() == false {
				// Add any song that came before and isn't already on the card
				randomSong := round[rand.Intn(winNumber)]
				if newCard.Contains(randomSong) {
					continue
				}
				newCard.AddSong(randomSong)
			}

			winningCards = append(winningCards, newCard)
		}

		// for _, card := range winningCards {
		// 	fmt.Printf("Card: wins at %s\n", round[card.WinsAt()-1].String())
		// }

		// Generate (N - winners) cards to fill out the rest of the cards
		// These will eventually win but we expect them to not make it that far
		// before 4 bingos are called
	//	fmt.Println("=========================================")

		backupCards := []*Card{}

		// Create however many cards are needed to fill out the round
		for len(backupCards) < (CARDS_PER_ROUND - len(winningCards)) {

			newCard := NewCard()

			// Pick a random song to win on between the end of the winning cards
			// and the end of the round. For example, if the last card wins on
			// song 25 and there are 30 songs in the round, we want to pick a
			// number between 26 and 30.

			lastWinNumber := WINNERS[roundIdx][len(WINNERS[roundIdx])-1]
			topIdx := rand.Intn(len(round)-lastWinNumber) + lastWinNumber
			newCard.AddSong(round[topIdx])

			// Fill the card with random songs while there are still empty
			// spots on the card
			for newCard.Complete() == false {
				// Add any song that came before and isn't already on the card
				randomSong := round[rand.Intn(topIdx)]
				if newCard.Contains(randomSong) {
					continue
				}
				newCard.AddSong(randomSong)
			}

			backupCards = append(backupCards, newCard)
		}

		// Check if backup cards have collisions
		// This is O(n^2) but n is small so it's fine
		for i, card := range backupCards {
			for j, otherCard := range backupCards {
				if card.isIdentical(otherCard) && i != j {
					fmt.Println("Duplicate card found!")
					fmt.Println(card.String())
					fmt.Println(otherCard.String())

					// We could remove the duplicate and add a new one but
					// I didn't extract the logic to generate a new backup card
					// into a function so I'm just going to panic for now

					// If this happens we just re-run the program, the odds
					// of this occurring are very low anyway
					panic("Duplicate card found")
				}
			}
		}

		fmt.Println("=========================================")
		fmt.Println("Round", roundIdx+1)

		fmt.Print("\033[32m") // Green color for winning cards
		for _, card := range winningCards {
			fmt.Println(card.Display())
		}
		fmt.Print("\033[33m") // Yellow color for backup cards
		for _, card := range backupCards {
			fmt.Println(card.Display())
		}
		fmt.Print("\033[0m") // Reset color
	}
}
