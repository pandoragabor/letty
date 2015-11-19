package gamelevel

import (
	tl "github.com/JoelOtter/termloop"
)

type Level struct {
	*tl.BaseLevel
	StartX, StartY int
	Enemies []Enemy
}

func (l *Level) Tick(ev tl.Event) {
	// move the enemies even if there was no event
	for _, e := range l.Enemies {
		e.Move()
	}

	// handle collisions, etc
	l.BaseLevel.Tick(ev)
}

type Wall struct {
	*Pattern
}

func NewWall(x, y, w, h int, color tl.Attr, ch_optional ...rune) (*Wall) {
	ch := '.'
	if len(ch_optional) > 0 {
		ch = ch_optional[0]
	}
	return &Wall{NewPattern(x, y, w, h, color, true, ch)}
}

func createLevelBase(wallColor tl.Attr, wallWidth int, w int, h int) *tl.BaseLevel {
	var level *tl.BaseLevel = tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
		Fg: tl.ColorWhite,
		Ch: ' ',
	})

	level.AddEntity(NewWall(0, 0, w + wallWidth, wallWidth, wallColor))
	level.AddEntity(NewWall(0, h + wallWidth, w + wallWidth, wallWidth, tl.ColorCyan))
	level.AddEntity(NewWall(0, 0, wallWidth, h + wallWidth, tl.ColorCyan))
	level.AddEntity(NewWall(w + wallWidth, 0, wallWidth, h + wallWidth * 2, tl.ColorCyan))

	return level
}

func createLevel1() (*Level) {
	var level *tl.BaseLevel = createLevelBase(tl.ColorCyan, 10, 100, 40)
	level.AddEntity(NewWall(40, 10, 10, 35, tl.ColorCyan))

	key1 := NewKey(15, 45, tl.ColorRed)
	level.AddEntity(key1)

	door1 := NewDoor(50, 38, 60, 1, tl.ColorGreen)
	level.AddEntity(door1)


	key2 := NewKey(108, 48, tl.ColorGreen)
	level.AddEntity(key2)

	door2 := NewDoor(41, 45, 2, 5, tl.ColorRed)
	level.AddEntity(door2)

	enemies := make([]Enemy, 0)
	for x := 50; x < 110; x++ {
		for y := 27; y > 17; y-- {
			ghost := NewGhost(x, y)
			level.AddEntity(ghost)
			enemies = append(enemies, ghost)
		}
	}

	return &Level{level, 20, 20, enemies}
}
