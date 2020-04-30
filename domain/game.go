package domain

import (
	"github.com/google/uuid"
	"strconv"
	"sync"
)

type Game struct {
	sync.Mutex
	GameID  string
	Name    string
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

func (g *Game) Move(player *Player, location int) bool {
	if len(g.Players) != 2 {
		return false
	}
	return g.Board.Move(player, location)

}

func (g *Game) Render() *GameRendered {
	result := &GameRendered{
		Board:       g.Board.Render(),
		Players:     make(map[string]string),
		PlayersTurn: g.Board.PlayerNumberTurn,
	}
	for number, player := range g.Players {
		result.Players[strconv.Itoa(number)] = player.Name
	}
	return result
}

type GameRendered struct {
	Board       [][]int           `json:"board"`
	Players     map[string]string `json:"players"`
	PlayersTurn int               `json:"playersTurn"`
}
