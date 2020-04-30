package domain

import "sync"

type Player struct {
	sync.Mutex
	PlayerID string
	GameID   string
	Number   int
}
