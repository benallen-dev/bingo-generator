package image

import (
	"image"
	"image/color"
	"image/png"

	"golang.org/x/image/font"
	"golang.org/x/image/font/inconsolata"
	"golang.org/x/image/math/fixed"

	"os"

	"github.com/benallen-dev/bingo-generator/pkg/bingo"
)

// DrawCard draws a bingo card
func DrawCard(card *bingo.Card, filename string) {
	// Create a new image
	img := image.NewRGBA(image.Rect(0, 0, 400, 400))

	// Draw the card
	drawCard(img, card)

	// Save the image
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	png.Encode(f, img)
}

func drawCard(img *image.RGBA, card *bingo.Card) {
	// Draw the card background
	drawCardBackground(img)

	// Draw the songs
	drawSongs(img, card)
}

func drawCardBackground(img *image.RGBA) {
	// Draw the background
	for x := 0; x < 400; x++ {
		for y := 0; y < 400; y++ {
			img.Set(x, y, color.RGBA{0, 0, 0, 255})
		}
	}
}

func drawSongs(img *image.RGBA, card *bingo.Card) {
	// Draw the songs
	for i, song := range card.Songs {
		drawText(img, 10, 10+i*40, song.String())
	}
}

func drawText(img *image.RGBA, x, y int, text string) {
	// Draw the text
	col := color.RGBA{255, 255, 255, 255}
	point := fixed.Point26_6{fixed.Int26_6(x * 64), fixed.Int26_6(y * 64)}
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: inconsolata.Regular8x16,
		Dot:  point,
	}
	d.DrawString(text)
}


