package gamelevel

import tl "github.com/JoelOtter/termloop"

type Door struct {
	*Pattern
}

func NewDoor(x, y, w, h int, color tl.Attr) (*Door) {
	return &Door{NewPattern(x, y, w, h, color, true, 'X')}
}

func (door *Door) Collide(collision tl.Physical) {
	if _, ok := collision.(*Player); ok {
		TheGameState.Player.OpenDoor(door)
	}
}
