package utils

import (
	"os"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type TileSet struct {
	Void       []*ebiten.Image
	Room       []*ebiten.Image
	Wall       []*ebiten.Image
	Door       []*ebiten.Image
	Sprite     []*ebiten.Image
	WallSprite []*ebiten.Image
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
	if ts.WallSprite, err = loadDir("wall-sprite"); err != nil {
		return nil, err
	}

	return ts, nil
}

func DrawDungeon(screen *ebiten.Image, visualGrid [][]VisualTile, ts *TileSet) {
	for y, row := range visualGrid {
		for x, tile := range row {
			var img *ebiten.Image

			switch tile.Type {
			case Void:
				img = ts.Void[tile.Variant]
			case Floor, Hallway:
				img = ts.Room[tile.Variant]
			case Wall:
				img = ts.Wall[tile.Variant]
			case Door:
				img = ts.Door[tile.Variant]
			}

			if img != nil {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(x*TileSize), float64(y*TileSize))
				screen.DrawImage(img, op)
			}

			// Draw sprite decorator on top
			if tile.Type == Floor && tile.Decorator != -1 {
				sImg := ts.Sprite[tile.Decorator]
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(x*TileSize), float64(y*TileSize))
				screen.DrawImage(sImg, op)
			}

			// Draw wall sprite decorator on top
			if tile.Type == Wall && tile.Decorator != -1 {
				sImg := ts.WallSprite[tile.Decorator]
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(x*TileSize), float64(y*TileSize))
				screen.DrawImage(sImg, op)
			}
		}
	}
}
