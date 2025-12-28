package utils

type TileType int

const (
	Void TileType = iota
	Floor
	Door
)

func GenerateGrid(width, height int, Rooms []*Rect, Hallways []Rect) [][]TileType {
	if width <= 0 || height <= 0 {
		return nil
	}

	// init grid with void tiles
	tiles := make([][]TileType, height)
	for i := range tiles {
		tiles[i] = make([]TileType, width)
	}

	for _, room := range Rooms {
		for y := room.Y; y < room.Y+room.H; y++ {
			for x := room.X; x < room.X+room.W; x++ {
				tiles[y][x] = Floor // Fills the entire rectangle
			}
		}
	}

	for _, hallway := range Hallways {
		for y := hallway.Y; y < hallway.Y+hallway.H; y++ {
			for x := hallway.X; x < hallway.X+hallway.W; x++ {
				tiles[y][x] = Floor
			}
		}
	}

	return tiles

}
