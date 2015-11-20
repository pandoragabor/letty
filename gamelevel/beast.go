package gamelevel

import tl "github.com/JoelOtter/termloop"
import "math/rand"
import (
	"time"
	"math"
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

func (beast *Beast) Draw(s *tl.Screen) {
	beast.Move()
	beast.Pattern.Draw(s)
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
	toX, toY := TheGameState.Player.Position()
	beast.prevX, beast.prevY = beast.Position()

	dx, dy := 0, 0
	if beast.prevX < toX {
		dx = 1
	} else if beast.prevX > toX {
		dx = -1
	}
	if beast.prevY < toY {
		dy = 1
	} else if beast.prevY > toY {
		dy = -1
	}

	if beast.blocked {
		// if we were blocked last time and can go either x or y, choose one
		if dx != 0 && dy != 0 {
			if rand.Intn(2) == 0 {
				dx = 0
			} else {
				dy = 0
			}
		}
	} else {
		// we were not blocked last time
		distance := math.Sqrt(float64(
			(beast.prevX - toX) * (beast.prevX - toX) +
			(beast.prevY - toY) * (beast.prevY - toY)))
		if distance > 2 {
			// if far away, add some wiggle to the movement
			if dx == 0 {
				dx = rand.Intn(2) - 1
			} else if dy == 0 {
				dy = rand.Intn(2) - 1
			}
		}
	}
	if dx != 0 || dy != 0 {
		beast.SetPosition(beast.prevX + dx, beast.prevY + dy)
	}
	beast.blocked = false
}

func (beast *Beast) Kill() {
	TheGameState.Level.RemoveEntity(beast)
	beast.active = false
}

func (beast *Beast) Collide(collision tl.Physical) {
	beast.SetPosition(beast.prevX, beast.prevY)
	beast.blocked = true
}
