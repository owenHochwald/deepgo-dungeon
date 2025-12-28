package utils

import "fmt"

type TileType int

const (
	Void TileType = iota
	Floor
	Door
	Wall
	Hallway
)

func PrintGrid(tiles [][]TileType) {
	if tiles == nil {
		fmt.Println("Grid is nil")
		return
	}

	tileSymbols := map[TileType]string{
		Void:    ".",
		Floor:   "#",
		Door:    "D",
		Wall:    "W",
		Hallway: "=",
	}

	for _, row := range tiles {
		for _, tile := range row {
			fmt.Print(tileSymbols[tile] + " ")
		}
		fmt.Println()
	}
}

func GenerateGrid(width, height int, Rooms []*Rect, Hallways []Rect) [][]TileType {
	if width <= 0 || height <= 0 {
		return nil
	}

	tiles := make([][]TileType, height)
	for i := range tiles {
		tiles[i] = make([]TileType, width)
	}

	for _, room := range Rooms {
		for y := room.Y; y < room.Y+room.H; y++ {
			for x := room.X; x < room.X+room.W; x++ {
				if y >= 0 && y < height && x >= 0 && x < width {
					tiles[y][x] = Floor
				}
			}
		}
	}

	for _, hallway := range Hallways {
		for y := hallway.Y; y < hallway.Y+hallway.H; y++ {
			for x := hallway.X; x < hallway.X+hallway.W; x++ {
				if y >= 0 && y < height && x >= 0 && x < width {
					// TODO: Door Generation
					//if tiles[y][x] == Floor {
					//	tiles[y][x] = Door
					//} else {
					tiles[y][x] = Floor
					//}
				}
			}
		}
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if tiles[y][x] == Void {
				if isAdjacentToNavigable(tiles, x, y, width, height) {
					tiles[y][x] = Wall
				}
			}
		}
	}

	return tiles
}

func isAdjacentToNavigable(tiles [][]TileType, x, y, w, h int) bool {
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			nx, ny := x+dx, y+dy
			if nx >= 0 && nx < w && ny >= 0 && ny < h {
				t := tiles[ny][nx]
				if t == Floor || t == Hallway || t == Door {
					return true
				}
			}
		}
	}
	return false
}
