package utils

import (
	"game/constants"

	"github.com/hajimehoshi/ebiten/v2"
)

const scale float64 = 5

type Sprite struct {
	X         int
	Y         int
	ImagePath string
}

type SpriteArray []*Sprite

func (sprite *Sprite) Draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}

	options.GeoM.Translate(float64(sprite.X*constants.TileSize), float64(sprite.Y*constants.TileSize))
	options.GeoM.Scale(scale, scale)

	screen.DrawImage(Images[sprite.ImagePath], options)
}

func (sprites *SpriteArray) Add(x int, y int, path string) {
	new_sprite := Sprite{x, y, createImage(path)}

	for i := 0; i < len(*sprites); i++ {
		if *(*sprites)[i] == new_sprite {
			return
		}
	}
	*sprites = append(*sprites, &new_sprite)
}