package gamelevel

import tl "github.com/JoelOtter/termloop"

type Door struct {
	*Pattern
}

func NewDoor(x, y, w, h int, color tl.Attr) (*Door) {
	return &Door{NewPattern(x, y, w, h, color, true, 'X')}
}
