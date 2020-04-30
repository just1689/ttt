package domain

import (
	"encoding/json"
	"github.com/google/uuid"
	"testing"
)

func TestNewBoard(t *testing.T) {
	t.Parallel()
	t.Log("a new board should be initialized to 0 for each addressable tile")
	board := NewBoard()
	for x := board.FirstRow; x <= board.LastRow; x++ {
		for y := board.FirstRow; y <= board.LastRow; y++ {
			value := board.Tiles[x][y]
			if value != 0 {
				t.Error("found a tile not set to 0 at", x, y)
			}
		}
	}
}

func TestBoard_Reset(t *testing.T) {
	t.Parallel()
	t.Log("resetting a board should set all addressable tiles to 0")
	board := NewBoard()
	for x := board.FirstRow; x <= board.LastRow; x++ {
		for y := board.FirstRow; y <= board.LastRow; y++ {
			board.Tiles[x][y] = 1
		}
	}
	board.Reset()
	for x := board.FirstRow; x <= board.LastRow; x++ {
		for y := board.FirstRow; y <= board.LastRow; y++ {
			value := board.Tiles[x][y]
			if value != 0 {
				t.Error("found a tile not set to 0 at", x, y)
			}
		}
	}

}

func TestBoard_Render(t *testing.T) {
	t.Parallel()
	game := NewGame()
	r := game.Board.Render()
	b, _ := json.Marshal(r)
	t.Log(string(b))
}

func TestBoard_MoveNegativeTurn(t *testing.T) {
	t.Log("you should not be able to move if its not your turn")
	game := NewGame()
	game.AddPlayer(&Player{PlayerID: uuid.New().String()})
	game.AddPlayer(&Player{PlayerID: uuid.New().String()})
	moved := game.Board.Move(game.Players[2], 1)
	if moved {
		t.Error("player 2 should not be able to move when its not their turn")
	}
}
func TestBoard_MovePositiveTurn(t *testing.T) {
	t.Log("you should not be able to move if its not your turn")
	game := NewGame()
	game.AddPlayer(&Player{PlayerID: uuid.New().String()})
	game.AddPlayer(&Player{PlayerID: uuid.New().String()})
	moved := game.Board.Move(game.Players[1], 1)
	if !moved {
		t.Error("player 1 should be able to move when its not their turn")
	}
}

func TestBoard_MovePositiveTurnsTwoPlayers(t *testing.T) {
	t.Log("you should not be able to move if its not your turn")
	game := NewGame()
	game.AddPlayer(&Player{PlayerID: uuid.New().String()})
	game.AddPlayer(&Player{PlayerID: uuid.New().String()})
	moved := game.Board.Move(game.Players[1], 1)
	if !moved {
		t.Error("player 1 should be able to move when its not their turn")
	}
	moved = game.Board.Move(game.Players[2], 2)
	if !moved {
		t.Error("player 2 should be able to move when its not their turn")
	}
}

func TestBoard_MoveNegativeTurnsTwoPlayers(t *testing.T) {
	t.Log("you should not be able to move if its not your turn")
	game := NewGame()
	game.AddPlayer(&Player{PlayerID: uuid.New().String()})
	game.AddPlayer(&Player{PlayerID: uuid.New().String()})
	moved := game.Board.Move(game.Players[1], 1)
	if !moved {
		t.Error("player 1 should be able to move when its not their turn")
	}
	moved = game.Board.Move(game.Players[1], 2)
	if moved {
		t.Error("player 1 should not be able to twice")
	}
	moved = game.Board.Move(game.Players[2], 2)
	if !moved {
		t.Error("player 2 should be able to twice")
	}
}

func TestBoard_MoveNegativeSameLocation(t *testing.T) {
	t.Log("you should not be able to move to a tile that is already occupied")
	game := NewGame()
	game.AddPlayer(&Player{PlayerID: uuid.New().String()})
	game.AddPlayer(&Player{PlayerID: uuid.New().String()})
	game.Board.Move(game.Players[1], 1)
	moved := game.Board.Move(game.Players[2], 1)
	if moved {
		t.Error("player was able to move to tile already occupied")
	}
}

func TestBoard_CheckForAvailableMovesPositive(t *testing.T) {
	t.Log("you should not be able to move to a tile that is already occupied")
	game := NewGame()
	game.AddPlayer(&Player{PlayerID: uuid.New().String()})
	game.AddPlayer(&Player{PlayerID: uuid.New().String()})
	lastPlayer := 2
	for location := 1; location < 9; location++ {
		if lastPlayer == 1 {
			lastPlayer = 2
		} else {
			lastPlayer = 1
		}
		game.Board.Move(game.Players[lastPlayer], location)
	}
	available := game.Board.CheckForAvailableMoves()
	if !available {
		t.Error("Board with 8 plays should still have moves available")
	}
}
func TestBoard_CheckForAvailableMovesNegative(t *testing.T) {
	t.Log("you should not be able to move to a tile that is already occupied")
	game := NewGame()
	game.AddPlayer(&Player{PlayerID: uuid.New().String()})
	game.AddPlayer(&Player{PlayerID: uuid.New().String()})
	lastPlayer := 2
	for location := 1; location <= 9; location++ {
		if lastPlayer == 1 {
			lastPlayer = 2
		} else {
			lastPlayer = 1
		}
		game.Board.Move(game.Players[lastPlayer], location)
	}
	available := game.Board.CheckForAvailableMoves()
	if available {
		t.Error("Board with 9 plays should not have moves available")
	}
}

func TestBoard_CheckForWinnerPositiveHorizontal(t *testing.T) {
	t.Log("three in a row of the same number should be a winner")
	game := NewGame()
	game.AddPlayer(&Player{PlayerID: uuid.New().String()})
	game.AddPlayer(&Player{PlayerID: uuid.New().String()})

	game.Board.Tiles[1][1] = 1
	game.Board.Tiles[2][1] = 1
	game.Board.Tiles[3][1] = 1

	winner, number := game.Board.CheckForWinner()
	if !winner {
		t.Error("was expecting winner")
	}
	if number != 1 {
		t.Error("winner should be player 1")
	}
}

func TestBoard_CheckForWinnerNegativeHorizontal(t *testing.T) {
	t.Log("three in a row of different numbers should not yield a winner")
	game := NewGame()
	game.AddPlayer(&Player{PlayerID: uuid.New().String()})
	game.AddPlayer(&Player{PlayerID: uuid.New().String()})

	game.Board.Tiles[1][1] = 1
	game.Board.Tiles[2][1] = 2
	game.Board.Tiles[3][1] = 1

	winner, number := game.Board.CheckForWinner()
	if winner {
		t.Error("was not expecting winner")
	}
	if number != 0 {
		t.Error("winner should be 0")
	}
}

func TestBoard_CheckForWinnerPositiveVertical(t *testing.T) {
	t.Log("three in a row of the same number should be a winner")
	game := NewGame()
	game.AddPlayer(&Player{PlayerID: uuid.New().String()})
	game.AddPlayer(&Player{PlayerID: uuid.New().String()})

	game.Board.Tiles[1][1] = 1
	game.Board.Tiles[1][2] = 1
	game.Board.Tiles[1][3] = 1

	winner, number := game.Board.CheckForWinner()
	if !winner {
		t.Error("was expecting winner")
	}
	if number != 1 {
		t.Error("winner should be player 1")
	}
}

func TestBoard_CheckForWinnerNegativeVertical(t *testing.T) {
	t.Log("three in a row of the same number should be a winner")
	game := NewGame()
	game.AddPlayer(&Player{PlayerID: uuid.New().String()})
	game.AddPlayer(&Player{PlayerID: uuid.New().String()})

	game.Board.Tiles[1][1] = 1
	game.Board.Tiles[1][2] = 2
	game.Board.Tiles[1][3] = 1

	winner, number := game.Board.CheckForWinner()
	if winner {
		t.Error("was not expecting winner")
	}
	if number != 0 {
		t.Error("winner should be 0")
	}
}

func TestBoard_CheckForWinnerPositiveDiagonal(t *testing.T) {
	t.Log("three in diagonally of the same number should be a winner")
	game := NewGame()
	game.AddPlayer(&Player{PlayerID: uuid.New().String()})
	game.AddPlayer(&Player{PlayerID: uuid.New().String()})

	game.Board.Tiles[1][1] = 1
	game.Board.Tiles[2][2] = 1
	game.Board.Tiles[3][3] = 1

	winner, number := game.Board.CheckForWinner()
	if !winner {
		t.Error("was expecting winner")
	}
	if number != 1 {
		t.Error("winner should be player 1")
	}
}

func TestBoard_CheckForWinnerNegativeDiagonal(t *testing.T) {
	t.Log("three in diagonally of the same number should be a winner")
	game := NewGame()
	game.AddPlayer(&Player{PlayerID: uuid.New().String()})
	game.AddPlayer(&Player{PlayerID: uuid.New().String()})

	game.Board.Tiles[1][1] = 1
	game.Board.Tiles[2][2] = 2
	game.Board.Tiles[3][3] = 1

	winner, number := game.Board.CheckForWinner()
	if winner {
		t.Error("was not expecting winner")
	}
	if number != 0 {
		t.Error("winner should be 0")
	}
}
