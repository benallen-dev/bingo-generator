package bingo

import (
	"fmt"
	"math/rand"
)

// GenerateBingoCards generates bingo cards for each round of bingo
func Generate(assetDir string, WINNERS [][]int, CARDS_PER_ROUND int) (output []Round) {

	rounds := readInput(assetDir)
	output = make([]Round, len(rounds))

	for roundIdx, round := range rounds {

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

		// Generate (N - winners) cards to fill out the rest of the cards
		// These will eventually win but we expect them to not make it that far
		// before 4 bingos are called

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

		output[roundIdx] = Round{
			WinningCards: winningCards,
			BackupCards:  backupCards,
		}
	}

	return output
}
