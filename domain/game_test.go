package domain

import (
	"github.com/google/uuid"
	"testing"
)

func TestGame_MoveNegativePreventMoveOnEmptyGame(t *testing.T) {
	t.Parallel()
	t.Log("games cant start until they have 2 players")
	game := NewGame()
	game.AddPlayer(&Player{PlayerID: uuid.New().String()})
	moved := game.Move(game.Players[1], 1)
	if moved {
		t.Error("should not have moved with only 1 player")
	}
}

func TestGame_MovePositivePreventMoveOnEmptyGame(t *testing.T) {
	t.Parallel()
	t.Log("games cant start until they have 2 players")
	game := NewGame()
	game.AddPlayer(&Player{PlayerID: uuid.New().String()})
	game.AddPlayer(&Player{PlayerID: uuid.New().String()})
	moved := game.Move(game.Players[1], 1)
	if !moved {
		t.Error("should have moved with 2 players")
	}
}
