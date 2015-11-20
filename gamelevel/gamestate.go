package gamelevel

import "fmt"
import "time"
import tl "github.com/JoelOtter/termloop"

type GameState struct {
	Player *Player
	LevelIndex int
	Level *Level
	Bullets []*Bullet
}

var TheGameState *GameState = new(GameState)

func (gamestate *GameState) StartLevel(levelIndex int, screen *tl.Screen) {
	gamestate.LevelIndex = levelIndex
	gamestate.Level = gamestate.CreateLevelByIndex()

	gamestate.Bullets = make([]*Bullet, 0)

	gamestate.Player = NewPlayer(gamestate.Level.StartX, gamestate.Level.StartY)
	gamestate.Level.AddEntity(gamestate.Player)

	screen.SetLevel(gamestate.Level)

	// animate the beasts
	go func() {
		for {
			for _, enemy := range TheGameState.Level.Enemies {
				enemy.Animate()
			}
			time.Sleep(500 * time.Millisecond)
		}
	}()
}

func (gamestate *GameState) CreateLevelByIndex() (*Level) {
	switch(gamestate.LevelIndex) {
	case 1: return createLevel1()
	}
	panic(fmt.Sprintf("Unknown level %d", gamestate.LevelIndex))
}

