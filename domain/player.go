package domain

import "sync"

type Player struct {
	sync.Mutex
	PlayerID string
	Secret   string
	GameID   string
	Name     string
	Number   int
}

func NewPlayer(playerID, secret, name string) *Player {
	return &Player{
		PlayerID: playerID,
		Secret:   secret,
		Name:     name,
	}
}
