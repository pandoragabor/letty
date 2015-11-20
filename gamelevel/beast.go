package gamelevel

import tl "github.com/JoelOtter/termloop"
import "math/rand"
import (
	"time"
)

type Enemy interface {
	tl.Drawable
	tl.DynamicPhysical
	Animate()
	Move()
}

type Beast struct {
	*Pattern
	runes []rune
	cell tl.Cell
	prevX, prevY int
	lastMillis, speed int64
	blocked bool
	active bool
}

func NewBeast(x, y int, color tl.Attr, runes []rune, speed int64) (*Beast) {
	return &Beast{NewPattern(x, y, 1, 1, color, false, runes[0]),
		runes,
		tl.Cell{color, tl.ColorBlack, runes[0]},
		x, y,
		0, speed,
		false,
		true}
}

func (beast *Beast) Animate() {
	if !beast.active {
		return
	}

	ch := beast.runes[0]
	if rand.Intn(10) > 7 {
		ch = beast.runes[rand.Intn(len(beast.runes))]
	}
	beast.cell.Ch = ch
	beast.SetCell(0, 0, &beast.cell)
}

func (beast *Beast) Move() {
	if !beast.active {
		return
	}

	if !beast.blocked {
		// if wasn't blocked last, time limit on the creature's speed
		millis := time.Now().UnixNano() / 1000000
		if millis - beast.lastMillis <= beast.speed {
			return
		}
		beast.lastMillis = millis
	}

	// aim for the player
	toX := TheGameState.Player.prevX
	toY := TheGameState.Player.prevY
	beast.prevX, beast.prevY = beast.Position()
	x := beast.prevX
	y := beast.prevY

	dx := 0
	dy := 0
	if x < toX {
		dx = 1
	} else if x > toX {
		dx = -1
	}
	if y < toY {
		dy = 1
	} else if y > toY {
		dy = -1
	}

	if beast.blocked {
		// if we were blocked last time and can go either x or y, choose one
		if dx != 0 && dy != 0 {
			if rand.Intn(2) == 0 {
				x += dx
			} else {
				y += dy
			}
		}
	} else {
		// we were not blocked last time
		x += dx
		y += dy
	}
	beast.SetPosition(x, y)
	beast.blocked = false
}

func (beast *Beast) Collide(collision tl.Physical) {
	if _, ok := collision.(*Bullet); ok {
		TheGameState.Level.RemoveEntity(beast)
		beast.active = false
	} else {
		beast.SetPosition(beast.prevX, beast.prevY)
		beast.blocked = true
	}
}

