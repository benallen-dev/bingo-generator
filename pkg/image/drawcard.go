package image

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/math/fixed"

	"github.com/benallen-dev/bingo-generator/pkg/bingo"
)

type boundingBox struct {
	x      int
	y      int
	width  int
	height int
}

var (
	BG_COLOR      = color.RGBA{255, 255, 255, 255}
	TEXT_COLOR    = color.RGBA{0, 0, 0, 255}
	WIDTH         = 900
	HEIGHT        = 900
	HEADER_HEIGHT = 100
	BOUNDS        = []boundingBox{
		{20, 120, 260, 260},
		{320, 120, 260, 260},
		{620, 120, 260, 260},
		{20, 420, 260, 260},
		{320, 420, 260, 260},
		{620, 420, 260, 260},
		{20, 720, 260, 260},
		{320, 720, 260, 260},
		{620, 720, 260, 260},
	}

	goFont *truetype.Font
)

func init() {
	var err error
	goFont, err = truetype.Parse(goregular.TTF)
	if err != nil {
		log.Fatal(err)
	}
}

func measureText(text string, fontSize int) int {
	face := truetype.NewFace(goFont, &truetype.Options{
		Size: float64(fontSize),
	})

	return font.MeasureString(face, text).Ceil()

}

func renderText(img *image.RGBA, xPos int, yPos int, color color.RGBA, maxWidth int, text string, fontSize int) {
	face := truetype.NewFace(goFont, &truetype.Options{
		Size: float64(fontSize),
	})

	point := fixed.Point26_6{
		X: fixed.Int26_6(xPos * 64),
		Y: fixed.Int26_6(yPos * 64),
	}
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(color),
		Face: face,
		Dot:  point,
	}

	d.DrawString(text)
}

// DrawCard draws a bingo card
func DrawCard(card *bingo.Card, filename string) {

	imageHeight := HEADER_HEIGHT + HEIGHT
	imageWidth := WIDTH

	// Create a new image
	img := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))

	// Draw the background
	for x := 0; x < imageWidth; x++ {
		for y := 0; y < imageHeight; y++ {
			img.Set(x, y, BG_COLOR)
		}
	}

	// Draw the title
	renderText(img, 20, 60, color.RGBA{0x00, 0x8c, 0xf4, 255}, 900, "DELTA DS1 Muziekbingo                     Ronde X", 40)

	// Horizontal lines
	for i := range WIDTH {
		img.Set(i, HEADER_HEIGHT, TEXT_COLOR)
		img.Set(i, HEADER_HEIGHT+(HEIGHT/3), TEXT_COLOR)
		img.Set(i, HEADER_HEIGHT+(2*HEIGHT/3), TEXT_COLOR)
	}

	// Vertical lines
	for i := HEADER_HEIGHT; i < imageHeight; i++ {
		img.Set(imageWidth/3, i, TEXT_COLOR)
		img.Set(2*imageWidth/3, i, TEXT_COLOR)
	}

	// Draw the songs
	for i, song := range card.Songs {

		xPosition := BOUNDS[i].x
		yPosition := BOUNDS[i].y + BOUNDS[i].height/2
		fontSize := 30

		artistWidth := measureText(song.Artist, fontSize)
		titleWidth := measureText(song.Title, fontSize)

		for artistWidth > BOUNDS[i].width || titleWidth > BOUNDS[i].width {
			fontSize--
			titleWidth = measureText(song.Title, fontSize)
			artistWidth = measureText(song.Artist, fontSize)
		}

		xCenter := xPosition + BOUNDS[i].width/2
		dashWidth := measureText("-", fontSize)
		yPadding := 8

		renderText(img, xCenter-(titleWidth/2), yPosition-(fontSize+yPadding), TEXT_COLOR, BOUNDS[i].width-20, song.Title, fontSize)
		renderText(img, xCenter-(dashWidth/2), yPosition, TEXT_COLOR, BOUNDS[i].width-20, "-", fontSize)
		renderText(img, xCenter-(artistWidth/2), yPosition+(fontSize+yPadding), TEXT_COLOR, BOUNDS[i].width-20, song.Artist, fontSize)

		// get bounding box
		b := BOUNDS[i]

		// draw bounding box for debugging
		for x := b.x; x < b.x+b.width; x++ {
			img.Set(x, b.y, color.RGBA{64, 64, 64, 255})
			img.Set(x, b.y+b.height, color.RGBA{64, 64, 64, 255})
		}
		for y := b.y; y < b.y+b.height; y++ {
			img.Set(b.x, y, color.RGBA{64, 64, 64, 255})
			img.Set(b.x+b.width, y, color.RGBA{64, 64, 64, 255})
		}

	}

	// Save the image
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	png.Encode(f, img)
}
