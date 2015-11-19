package gamelevel

import tl "github.com/JoelOtter/termloop"
import "math/rand"

type Beast interface {
	tl.Drawable
	tl.DynamicPhysical
	Animate()
	Move()
}

type BasicBeast struct {
	*Pattern
	runes []rune
	cell tl.Cell
	prevX, prevY int
}

func NewBasicBeast(x, y int, color tl.Attr, runes []rune) (*BasicBeast) {
	return &BasicBeast{NewPattern(x, y, 1, 1, color, false, runes[0]),
		runes,
		tl.Cell{color, tl.ColorBlack, runes[0]},
		x, y}
}

func (beast *BasicBeast) Animate() {
	ch := beast.runes[0]
	if rand.Intn(10) > 7 {
		ch = beast.runes[rand.Intn(len(beast.runes))]
	}
	beast.cell.Ch = ch
	beast.SetCell(0, 0, &beast.cell)
}

func (beast *BasicBeast) Move() {
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

func (beast *BasicBeast) Collide(collision tl.Physical) {
	beast.SetPosition(beast.prevX, beast.prevY)
}


type Ghost struct {
	*BasicBeast
}

func NewGhost(x, y int) (*Ghost) {
	return &Ghost{NewBasicBeast(x, y, tl.ColorYellow, []rune{'o', 'O'})}
}

