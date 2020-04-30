package domain

import (
	"errors"
	"github.com/sirupsen/logrus"
	"strconv"
	"sync"
)

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

func (i *instance) List() []map[string]string {
	i.Lock()
	defer i.Unlock()
	result := make([]map[string]string, len(i.gamesByID))
	if len(i.gamesByID) == 0 {
		return result
	}
	gameRow := 0
	for _, game := range i.gamesByID {
		result[gameRow] = make(map[string]string)
		result[gameRow]["ID"] = game.GameID
		result[gameRow]["NAME"] = game.Name
		result[gameRow]["PLAYERS"] = strconv.Itoa(len(game.Players))
		gameRow++
	}
	return result
}

func (i *instance) New(name string) (game *Game) {
	i.Lock()
	defer i.Unlock()
	game = NewGame()
	game.Name = name
	i.gamesByID[game.GameID] = game
	return
}

func (i *instance) AddPlayer(player *Player) {
	i.Lock()
	defer i.Unlock()
	i.playersByID[player.PlayerID] = player
}

func (i *instance) JoinGame(gameID string, player *Player) (bool, string) {
	i.Lock()
	defer i.Unlock()
	game, found := LocalInstance.gamesByID[gameID]
	if !found {
		return false, ""
	}
	return game.AddPlayer(player), game.GameID
}

func (i *instance) GetPlayer(playerID, secret string) (*Player, bool) {
	i.Lock()
	defer i.Unlock()
	player, found := i.playersByID[playerID]
	if !found {
		logrus.Errorln("no player with ID", playerID)
		return nil, false
	}
	if player.Secret != secret {
		logrus.Errorln("no player with secret", secret)
		return nil, false
	}
	return player, true
}

func (i *instance) Move(player *Player, location int) bool {
	i.Lock()
	defer i.Unlock()
	game, found := i.gamesByID[player.GameID]
	if !found {
		return false
	}
	return game.Move(player, location)
}

func (i *instance) RenderGame(gameID string) (*GameRendered, error) {
	game, found := i.gamesByID[gameID]
	if !found {
		return nil, errors.New("game not found")
	}
	return game.Render(), nil
}
