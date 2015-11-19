package gamelevel

import tl "github.com/JoelOtter/termloop"

type Key struct {
	*Pattern
}

func NewKey(x, y int, color tl.Attr) (*Key) {
	pattern := NewPattern(x, y, 1, 1, color, false, '⟟')
	return &Key{pattern}
}

func (key *Key) Collide(collision tl.Physical) {
	if _, ok := collision.(*Player); ok {
		TheGameState.Player.AddKey(key)
	}
}

