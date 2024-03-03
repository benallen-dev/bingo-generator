package bingo

type Round struct {
	WinningCards []*Card
	BackupCards  []*Card
}

func (r *Round) Display() (output string) {

	output = ""

	output += "\033[32m" // Green color for winning cards
	output += "Winning cards ===========================\n"

	for _, card := range r.WinningCards {
		output += card.OneLine() + "\n"
	}

	output += "\n"
	output += "\033[33m" // Yellow color for backup cards
	output += "Backup cards ============================\n"

	for _, card := range r.BackupCards {
		output += card.OneLine() + "\n"
	}
	output += "\n"
	output += "\033[0m" // Reset color

	return output
}
