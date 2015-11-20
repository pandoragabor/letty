package gamelevel

import (
	tl "github.com/JoelOtter/termloop"
)

type Bullet struct {
	*Pattern
	dx, dy int
	active bool
}

func NewBullet(x, y, dx, dy int) (*Bullet) {
	var ch rune = '-'
	if dy != 0 {
		ch = '|'
	}
	return &Bullet{NewPattern(x, y, 1, 1, tl.ColorWhite, false, ch), dx, dy, true}
}

func (bullet *Bullet) Move() {
	if bullet.active {
		x, y := bullet.Position()
		bullet.SetPosition(x + bullet.dx, y + bullet.dy)
	}
}

func (bullet *Bullet) Collide(collision tl.Physical) {
	TheGameState.Level.RemoveEntity(bullet)
	bullet.active = false
}