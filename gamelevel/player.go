package gamelevel

import (
	tl "github.com/JoelOtter/termloop"
	"time"
)

type Player struct {
    entity *tl.Entity
	prevX  int
	prevY  int
	keys   map[tl.Attr]int
	lastFireMillis int64
}

func NewPlayer(sx, sy int) (*Player) {
	player := Player{
		entity: tl.NewEntity(sx, sy, 1, 1),
		keys: make(map[tl.Attr]int, 100),
	}
	player.entity.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: '옷'})
	return &player
}

func (player *Player) AddKey(key *Key) {
	color := key.Color
	if keys, ok := player.keys[color]; ok {
		player.keys[color] = keys + 1
	} else {
		player.keys[color] = 1
	}
	TheGameState.Level.RemoveEntity(key)
}

func (player *Player) OpenDoor(door *Door) bool {
	color := door.Color
	if keys, ok := player.keys[color]; ok && keys > 0 {
		TheGameState.Level.RemoveEntity(door)
		player.keys[color] = keys - 1
		return true
	}
	return false
}

func (player *Player) Draw(screen *tl.Screen) {
	screenWidth, screenHeight := screen.Size()
	x, y := player.entity.Position()
	TheGameState.Level.SetOffset(screenWidth/2-x, screenHeight/2-y)
	player.entity.Draw(screen)
}

func (player *Player) Fire(dx, dy int) {
	millis := time.Now().UnixNano() / 1000000
	if millis - player.lastFireMillis <= 300 {
		return
	}
	player.lastFireMillis = millis

	x, y := player.entity.Position()
	bullet := NewBullet(x + dx, y + dy, dx, dy)
	TheGameState.Level.AddEntity(bullet)
	TheGameState.Bullets = append(TheGameState.Bullets, bullet)
}

func (player *Player) Tick(event tl.Event) {
	if event.Type == tl.EventKey { // Is it a keyboard event?
		player.prevX, player.prevY = player.entity.Position()
		switch event.Key { // If so, switch on the pressed key.
		case tl.KeyArrowRight:
			player.entity.SetPosition(player.prevX+1, player.prevY)
			break
		case tl.KeyArrowLeft:
			player.entity.SetPosition(player.prevX-1, player.prevY)
			break
		case tl.KeyArrowUp:
			player.entity.SetPosition(player.prevX, player.prevY-1)
			break
		case tl.KeyArrowDown:
			player.entity.SetPosition(player.prevX, player.prevY+1)
			break
     	}
		switch event.Ch {
			case 'a': player.Fire(-1, 0)
			case 'd': player.Fire(1, 0)
			case 'w': player.Fire(0, -1)
			case 's': player.Fire(0, 1)
		}
	}
}

func (player *Player) Size() (int, int) {
    return player.entity.Size()
}

func (player *Player) Position() (int, int) {
    return player.entity.Position()
}

func (player *Player) Collide(collision tl.Physical) {
	var blocked bool
	switch collision.(type) {
	case *Wall:
		blocked = true
	case *Door:
		blocked = !player.OpenDoor(collision.(*Door))
	case *Key:
		player.AddKey(collision.(*Key))
		blocked = false
	}
	if blocked {
		player.entity.SetPosition(player.prevX, player.prevY)
	}
}