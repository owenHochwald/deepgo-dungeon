package utils

import "fmt"

type TileType int

const (
	Void TileType = iota
	Floor
	Door
)

func PrintGrid(tiles [][]TileType) {
	if tiles == nil {
		fmt.Println("Grid is nil")
		return
	}

	tileSymbols := map[TileType]string{
		Void:  ".",
		Floor: "#",
		Door:  "D",
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
					if tiles[y][x] == Floor {
						tiles[y][x] = Door
					} else {
						tiles[y][x] = Floor
					}
				}
			}
		}
	}

	return tiles
}
