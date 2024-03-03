package bingo

import "fmt"

type Song struct {
	PlayPos int
	Title   string
	Artist  string
}

func (s Song) String() string {
	return fmt.Sprintf("%d:\t\t%s - %s", s.PlayPos, s.Title, s.Artist)
}
