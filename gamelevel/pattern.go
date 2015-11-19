package gamelevel

import tl "github.com/JoelOtter/termloop"

type Pattern struct {
	*tl.Entity
	Color tl.Attr
}

func NewPattern(x, y, w, h int, color tl.Attr, inverse bool, ch rune) (*Pattern) {
	canvas := tl.NewCanvas(w, h)
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			if inverse {
				canvas[i][j].Fg = tl.ColorBlack
				canvas[i][j].Bg = color
			} else {
				canvas[i][j].Fg = color
			}
			canvas[i][j].Ch = ch
		}
	}
	pattern := Pattern{
		tl.NewEntityFromCanvas(x, y, canvas),
		color,
	}
	return &pattern
}

func (pattern *Pattern) Tick(event tl.Event) {
}
