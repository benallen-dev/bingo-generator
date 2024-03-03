package bingo

import (
	"fmt"
	"errors"
	"math/rand"
)

type Card struct {
	Songs[] Song
}

func (c *Card) String() string {
	output := fmt.Sprintf("Card (wins at %d):\n", c.WinsAt())
	for _, song := range c.Songs {
		output += song.String() + "\n"
	}
	return output
}

func (c *Card) OneLine() string {
	return fmt.Sprintf("Card (wins at %d)", c.WinsAt())
}

func (c *Card) Display() string {
	foo := []int{0, 1, 2, 3, 4, 5, 6, 7, 8 }
	rand.Shuffle(len(foo), func(i, j int) { foo[i], foo[j] = foo[j], foo[i] })

	output := fmt.Sprintf("Card (wins at %d):\n", c.WinsAt())
	for _, i := range foo {
		output += fmt.Sprintf("\t%s\n", c.Songs[i].String())
	}
	output += "\n"

	return output
}

func NewCard() *Card {
	return &Card{
		Songs: []Song{},
	}
}

func (c *Card) AddSong(s Song) error {
	if len(c.Songs) >= 9 {
		return errors.New("Card is full")
	}

	c.Songs = append(c.Songs, s)

	return nil
}

func (c *Card) RemoveSong(s Song) error {
	for i, song := range c.Songs {
		if song == s {
			c.Songs = append(c.Songs[:i], c.Songs[i+1:]...)
			return nil
		}
	}
	return errors.New("Song not found")
}

func (c *Card) WinsAt() int {
	highestPosition := 0

	for _, song := range c.Songs {
		if song.PlayPos > highestPosition {
			highestPosition = song.PlayPos
		}
	}

	return highestPosition
}

func (c *Card) Contains(s Song) bool {
	for _, song := range c.Songs {
		if song.PlayPos == s.PlayPos {
			return true
		}
	}
	return false
}

func (c *Card) Complete() bool {
	return len(c.Songs) == 9
}

func (c *Card) isIdentical(other *Card) bool {
	if len(c.Songs) != len(other.Songs) {
		return false
	}

	for i, song := range c.Songs {
		if song != other.Songs[i] {
			return false
		}
	}

	return true
}
