package domain

import (
	"github.com/google/uuid"
	"sync"
)

type Game struct {
	sync.Mutex
	GameID  string
	Players map[int]*Player
	Board   *Board
}

func NewGame() *Game {
	return &Game{
		GameID:  uuid.New().String(),
		Players: make(map[int]*Player),
		Board:   NewBoard(),
	}
}

func (g *Game) AddPlayer(player *Player) bool {
	g.Lock()
	player.Lock()
	defer g.Unlock()
	defer player.Unlock()
	if len(g.Players) >= 2 {
		return false
	}
	for i := 1; i <= 2; i++ {
		_, found := g.Players[i]
		if !found {
			g.Players[i] = player
			player.Number = i
			player.GameID = g.GameID
			return true
		}
	}
	return false
}

func (g *Game) RemovePlayer(player *Player) {
	g.Lock()
	player.Lock()
	defer g.Unlock()
	defer player.Unlock()

	delete(g.Players, player.Number)
	player.Number = 0
	player.GameID = ""
}
