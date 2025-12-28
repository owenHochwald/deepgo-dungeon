package utils

import (
	"os"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type TileSet struct {
	Void   []*ebiten.Image
	Room   []*ebiten.Image
	Wall   []*ebiten.Image
	Door   []*ebiten.Image
	Sprite []*ebiten.Image
}

func LoadTileSet(basePath string) (*TileSet, error) {
	ts := &TileSet{}
	var err error

	loadDir := func(dir string) ([]*ebiten.Image, error) {
		var imgs []*ebiten.Image
		path := filepath.Join(basePath, dir)
		files, err := os.ReadDir(path)
		if err != nil {
			return nil, err
		}
		for _, f := range files {
			if filepath.Ext(f.Name()) == ".png" {
				img, _, err := ebitenutil.NewImageFromFile(filepath.Join(path, f.Name()))
				if err != nil {
					return nil, err
				}
				imgs = append(imgs, img)
			}
		}
		return imgs, nil
	}

	if ts.Void, err = loadDir("void"); err != nil {
		return nil, err
	}
	if ts.Room, err = loadDir("room"); err != nil {
		return nil, err
	}
	if ts.Wall, err = loadDir("wall"); err != nil {
		return nil, err
	}
	if ts.Door, err = loadDir("door"); err != nil {
		return nil, err
	}
	if ts.Sprite, err = loadDir("sprite"); err != nil {
		return nil, err
	}

	return ts, nil
}

func DrawDungeon(screen *ebiten.Image, grid [][]TileType, ts *TileSet) {
	height := len(grid)
	if height == 0 {
		return
	}
	width := len(grid[0])

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// Deterministic "random" index based on coordinates
			idx := x*31 + y*17

			var img *ebiten.Image
			switch grid[y][x] {
			case Void:
				img = ts.Void[idx%len(ts.Void)]
			case Floor, Hallway:
				img = ts.Room[idx%len(ts.Room)]
			case Wall:
				img = ts.Wall[idx%len(ts.Wall)]
			case Door:
				img = ts.Door[idx%len(ts.Door)]
			}

			if img != nil {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(x*TileSize), float64(y*TileSize))
				screen.DrawImage(img, op)
			}

			// Decoration Sprites on Room/Hallway w/ 10% chance
			if (grid[y][x] == Floor || grid[y][x] == Hallway) && (idx%10 == 0) {
				sImg := ts.Sprite[(idx/10)%len(ts.Sprite)]
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(x*TileSize), float64(y*TileSize))
				screen.DrawImage(sImg, op)
			}
		}
	}
}
