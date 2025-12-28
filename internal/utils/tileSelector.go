package utils

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

func GetTile(sheet *ebiten.Image, id int) *ebiten.Image {
	const tilePerRow = 16
	const tileSize = 16

	sx := (id % tilePerRow) * tileSize
	sy := (id / tilePerRow) * tileSize

	return sheet.SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize)).(*ebiten.Image)
}
