package top

import (
	"game/constants"
	"game/utils"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	Pos       *utils.Point
	Last      *utils.Point
	Sprite    *utils.Sprite
	Direction constants.Direction
	Moving    bool
}

func CreatePlayer(x, y float64) Player {
	return Player{
		utils.CreatePoint(x, y),
		utils.CreatePoint(x, y),
		utils.CreateSprite("character/down_1.png"),
		constants.Down,
		false,
	}
}

func (p *Player) Update(keys *Keyboard) {
	if !p.Moving {
		dir, move := constants.NoDirection, false

		for _, key := range keys.Keys {
			switch key {
			case ebiten.KeyW:
				dir, move = dir.Combine(constants.Up), true
			case ebiten.KeyS:
				dir, move = dir.Combine(constants.Down), true
			case ebiten.KeyA:
				dir, move = dir.Combine(constants.Left), true
			case ebiten.KeyD:
				dir, move = dir.Combine(constants.Right), true
			}
		}

		if move {
			p.Direction, p.Moving = dir, true
		}
	}

	if p.Moving {
		for _, dir := range p.Direction.Deconstruct() {
			switch dir {
			case constants.Up:
				p.Pos.Y -= constants.PlayerSpeed
			case constants.Down:
				p.Pos.Y += constants.PlayerSpeed
			case constants.Left:
				p.Pos.X -= constants.PlayerSpeed
			case constants.Right:
				p.Pos.X += constants.PlayerSpeed
			}
		}

		dx, dy := p.Pos.X-p.Last.X, p.Pos.Y-p.Last.Y
		if math.Abs(dx) >= 1 || math.Abs(dy) >= 1 {
			p.Pos.X, p.Pos.Y = p.Last.X+utils.Normalize(dx), p.Last.Y+utils.Normalize(dy)
			p.Last.X, p.Last.Y = p.Pos.X, p.Pos.Y
			p.Moving = false
		}
	}
}

func (p *Player) Draw(screen *ebiten.Image, options ebiten.DrawImageOptions) {
	p.Sprite.Draw(screen, options, p.Pos.X, p.Pos.Y)
}
