package bingo

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func readInput(assetPath string) (output [][]Song) {

	// Read files from the assets 
	dirEntries, err := os.ReadDir(assetPath)
	if err != nil {
		log.Fatal(err)
	}

	output = make([][]Song, len(dirEntries))

	// Print the files
	for idx, dirEntry := range dirEntries {

		// Open the file
		file, err := os.Open(assetPath + "/" + dirEntry.Name())
		if err != nil {
			log.Fatal(err)
		}

		// Read the file
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			lineparts := strings.Split(line, " - ")

			if len(lineparts) != 2 {
				log.Printf("Invalid line in %s: %s", dirEntry.Name(), line)
				continue
			}

			output[idx] = append(output[idx], Song{
				PlayPos: len(output[idx]) + 1,
				Title:  lineparts[0],
				Artist: lineparts[1],
			})
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}

	return output
}
