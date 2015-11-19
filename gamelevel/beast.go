package gamelevel

import tl "github.com/JoelOtter/termloop"
import "math/rand"
import "time"

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
}

func NewBeast(x, y int, color tl.Attr, runes []rune, speed int64) (*Beast) {
	return &Beast{NewPattern(x, y, 1, 1, color, false, runes[0]),
		runes,
		tl.Cell{color, tl.ColorBlack, runes[0]},
		x, y,
		0, speed}
}

func (beast *Beast) Animate() {
	ch := beast.runes[0]
	if rand.Intn(10) > 7 {
		ch = beast.runes[rand.Intn(len(beast.runes))]
	}
	beast.cell.Ch = ch
	beast.SetCell(0, 0, &beast.cell)
}

func (beast *Beast) Move() {
	millis := time.Now().UnixNano() / 1000000
	if millis - beast.lastMillis <= beast.speed {
		return
	}

	beast.lastMillis = millis
	toX := TheGameState.Player.prevX
	toY := TheGameState.Player.prevY
	beast.prevX, beast.prevY = beast.Position()
	x := beast.prevX
	y := beast.prevY
	if x < toX {
		x++
	} else if x > toX {
		x--
	}
	if y < toY {
		y++
	} else if y > toY {
		y--
	}
	beast.SetPosition(x, y)
}

func (beast *Beast) Collide(collision tl.Physical) {
	beast.SetPosition(beast.prevX, beast.prevY)
}

type Ghost struct {
	*Beast
}

func NewGhost(x, y int) (*Ghost) {
	return &Ghost{NewBeast(x, y, tl.ColorYellow, []rune{'o', 'O'}, 100)}
}

