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
	defer g.Unlock()
	if len(g.Players) >= 2 {
		return false
	}
	for i := 1; i <= 2; i++ {
		_, found := g.Players[i]
		if !found {
			player.Lock()
			g.Players[i] = player
			player.Number = i
			player.GameID = g.GameID
			player.Unlock()
			return true
		}
	}
	return false
}

func (g *Game) RemovePlayer(player *Player) {
	g.Lock()
	player.Lock()

	delete(g.Players, player.Number)
	player.Number = 0
	player.GameID = ""

	g.Unlock()
	player.Unlock()

}

func (g *Game) Move(player *Player, location int) bool {
	g.Lock()
	defer g.Unlock()
	if len(g.Players) != 2 {
		return false
	}
	return g.Board.Move(player, location)

}

func (g *Game) Render() *GameRendered {
	g.Lock()
	defer g.Unlock()
	result := &GameRendered{
		Board:       g.Board.Render(),
		Players:     make(map[string]string),
		PlayersTurn: g.Board.PlayerNumberTurn,
		CanMove:     g.Board.CheckForAvailableMoves(),
	}
	_, result.WinnerNumber = g.Board.CheckForWinner()

	for number, player := range g.Players {
		result.Players[strconv.Itoa(number)] = player.Name
	}
	return result
}

type GameRendered struct {
	Board        [][]int           `json:"board"`
	Players      map[string]string `json:"players"`
	PlayersTurn  int               `json:"playersTurn"`
	WinnerNumber int               `json:"winnerNumber"`
	CanMove      bool              `json:"canMove"`
}
