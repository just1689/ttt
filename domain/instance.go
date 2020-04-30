package domain

import "sync"

var LocalInstance = NewInstance()

type instance struct {
	sync.Mutex
	playersByID map[string]*Player
	gamesByID   map[string]*Game
}

func NewInstance() *instance {
	return &instance{
		playersByID: make(map[string]*Player),
		gamesByID:   make(map[string]*Game),
	}
}
