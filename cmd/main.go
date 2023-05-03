package main

import (
	"log"
	"os"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"

	rating "local.packages/rating"
)

const (
	screenWidth  = 320
	screenHeight = 240
)

type Game struct {
	rating *rating.Rating
	once   sync.Once
}

func (g *Game) Update() error {
	g.once.Do(func() {
		g.rating = rating.NewRating(StarImage, 15, 20, 10)
		g.rating.SetValue(5.7)
	})

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.rating.SetValue(9.8)
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.rating.SetValue(3.2)
	}
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		g.rating.SetValue(5.7)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.rating.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	os.Exit(run())
}

func run() int {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("ebiten game")
	err := ebiten.RunGame(&Game{})
	if err != nil {
		log.Println("failed to run game!")
		return 1
	}

	return 0
}
