package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/owenHochwald/deepgo-dungeon/internal/utils"
)

const (
	screenWidth  = 320 * 4
	screenHeight = 240 * 4
)

type Game struct {
	TreeRoot *utils.Node
	Rooms    []*utils.Node
	Hallways []utils.Rect
}

func NewGame() *Game {
	n := utils.NewNode(0, 0, screenWidth, screenHeight, 0)
	n.Split(10, 3)

	rooms := n.GetLeaves()

	var hallways []utils.Rect
	n.CreateHallways(&hallways)

	return &Game{
		TreeRoot: n,
		Rooms:    rooms,
		Hallways: hallways,
	}
}

func (g *Game) Update() error {
	return nil
}
func (g *Game) Draw(screen *ebiten.Image) {
	// draw hallways
	for _, hallway := range g.Hallways {
		vector.FillRect(
			screen,
			float32(hallway.X), float32(hallway.Y),
			float32(hallway.W), float32(hallway.H),
			color.RGBA{R: 150, G: 150, B: 150, A: 255}, false)
	}

	// draw rooms
	for _, node := range g.Rooms {
		vector.StrokeRect(
			screen,
			float32(node.Container.X), float32(node.Container.Y),
			float32(node.Container.W), float32(node.Container.H),
			1, color.RGBA{R: 50, G: 50, B: 50, A: 255}, false,
		)

		vector.FillRect(
			screen,
			float32(node.Room.X), float32(node.Room.Y),
			float32(node.Room.W), float32(node.Room.H),
			color.White, false,
		)
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
