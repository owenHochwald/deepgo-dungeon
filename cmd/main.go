package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
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
	Grid     [][]utils.VisualTile
	TileSet  *utils.TileSet
}

func NewGame() *Game {
	ts, err := utils.LoadTileSet("public/tiles")
	if err != nil {
		log.Fatal(err)
	}

	n := utils.NewNode(0, 0, screenWidth, screenHeight, 0)
	n.Split(minRoomSize, maxRecursionDepth)

	rooms := n.GetLeaves()

	var hallways []utils.Rect
	n.CreateHallways(&hallways)

	tiledRooms := utils.CreateTiledRooms(rooms)
	tiledHallways := utils.CreateTiledHallways(hallways)

	tiles := utils.GenerateGrid(gridWidth, gridHeight, tiledRooms, tiledHallways)
	visualTiles := utils.GenerateVisualGrid(tiles, ts)

	return &Game{
		TreeRoot: n,
		Rooms:    rooms,
		Hallways: hallways,
		Grid:     visualTiles,
		TileSet:  ts,
	}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	utils.DrawDungeon(screen, g.Grid, g.TileSet)
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
