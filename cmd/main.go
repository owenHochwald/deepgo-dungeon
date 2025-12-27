package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth  = 320 * 2
	screenHeight = 240 * 2
	tileSize     = 16
)

type Game struct {
	dungeonMap [][]int // 0 -> Wall, 1 -> Floor
}

func NewGame() *Game {
	return &Game{
		dungeonMap: [][]int{
			{0, 0, 0, 0, 1, 0, 0, 1, 0, 0},
			{0, 1, 0, 0, 1, 0, 1, 1, 0, 0},
			{0, 1, 1, 0, 0, 0, 0, 1, 0, 1},
			{0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
			{1, 0, 1, 0, 0, 0, 1, 1, 0, 0},
			{0, 0, 1, 1, 0, 0, 0, 0, 1, 0},
			{0, 1, 0, 0, 0, 1, 1, 0, 0, 0},
			{0, 1, 1, 0, 0, 0, 0, 1, 1, 0},
			{0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
			{1, 0, 0, 1, 0, 0, 1, 0, 0, 0},
		},
	}
}

func (g *Game) Update() error {
	return nil
}
func (g *Game) Draw(screen *ebiten.Image) {
	for y, row := range g.dungeonMap {
		for x, tile := range row {
			if tile == 1 {
				vector.FillRect(
					screen,
					float32(x*tileSize),
					float32(y*tileSize),
					tileSize,
					tileSize,
					color.White,
					false,
				)
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	g := NewGame()

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Retro Dungeon Generator")
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
