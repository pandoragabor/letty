package main

import tl "github.com/JoelOtter/termloop"
import "github.com/pandoragabor/letty/gamelevel"

func main() {
	g := tl.NewGame()
	g.Screen().SetFps(30)

	gamelevel.TheGameState.StartLevel(1, g.Screen())

	g.Screen().AddEntity(tl.NewFpsText(0, 0, tl.ColorRed, tl.ColorDefault, 0.5))
	g.Start()
}
