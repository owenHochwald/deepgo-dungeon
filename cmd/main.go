package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/owenHochwald/deepgo-dungeon/internal/utils"
)

const (
	screenWidth       = 320 * 4
	screenHeight      = 240 * 4
	minRoomSize       = 4 * utils.TileSize
	maxRecursionDepth = 4
	gridWidth         = screenWidth / utils.TileSize
	gridHeight        = screenHeight / utils.TileSize
)

type Game struct {
	TreeRoot *utils.Node
	Rooms    []*utils.Rect
	Hallways []utils.Rect
	Grid     [][]utils.TileType
}

func NewGame() *Game {
	n := utils.NewNode(0, 0, screenWidth, screenHeight, 0)
	n.Split(minRoomSize, maxRecursionDepth)

	rooms := n.GetLeaves()

	var hallways []utils.Rect
	n.CreateHallways(&hallways)

	tiledRooms := utils.CreateTiledRooms(rooms)
	tiledHallways := utils.CreateTiledHallways(hallways)

	tiles := utils.GenerateGrid(gridWidth, gridHeight, tiledRooms, tiledHallways)
	utils.PrintGrid(tiles)

	return &Game{
		TreeRoot: n,
		Rooms:    rooms,
		Hallways: hallways,
		Grid:     tiles,
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
			float32(node.X), float32(node.Y),
			float32(node.W), float32(node.H),
			1, color.RGBA{R: 50, G: 50, B: 50, A: 255}, false,
		)

		vector.FillRect(
			screen,
			float32(node.X), float32(node.Y),
			float32(node.W), float32(node.H),
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
