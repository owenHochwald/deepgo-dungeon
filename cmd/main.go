package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/owenHochwald/deepgo-dungeon/internal/utils"
)

const (
	screenWidth  = 320 * 2
	screenHeight = 240 * 2
)

type Game struct {
	dungeonMap [][]int // 0 -> Wall, 1 -> Floor
	TreeRoot   *utils.Node
}

func NewGame() *Game {
	n := utils.NewNode(0, 0, screenWidth, screenHeight, 0)
	n.Split(10, 3)

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
		TreeRoot: n,
	}
}

func (g *Game) Update() error {
	return nil
}
func (g *Game) Draw(screen *ebiten.Image) {
	children := g.TreeRoot.GetLeaves()

	for _, node := range children {
		vector.FillRect(
			screen,

			float32(node.Room.X),
			float32(node.Room.Y),
			float32(node.Room.W),
			float32(node.Room.H),
			color.White,
			false)

	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	g := NewGame()

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Retro Dungeon Generator")
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
