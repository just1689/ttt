package domain

import "github.com/sirupsen/logrus"

var locationMap = map[int][]int{
	1: {1, 1},
	2: {2, 1},
	3: {3, 1},
	4: {1, 2},
	5: {2, 2},
	6: {3, 2},
	7: {1, 3},
	8: {2, 3},
	9: {3, 3},
}

type Board struct {
	Tiles            map[int]map[int]int
	FirstRow         int
	LastRow          int
	PlayerNumberTurn int
}

func NewBoard() *Board {
	result := &Board{
		Tiles:            make(map[int]map[int]int),
		FirstRow:         1,
		LastRow:          3,
		PlayerNumberTurn: 1,
	}
	result.Reset()
	return result
}

func (b *Board) Reset() {
	for x := b.FirstRow; x <= b.LastRow; x++ {
		b.Tiles[x] = make(map[int]int)
		for y := b.FirstRow; y <= b.LastRow; y++ {
			b.Tiles[x][y] = 0
		}
	}
}

func (b *Board) Move(player *Player, location int) bool {
	if player.Number != b.PlayerNumberTurn {
		return false
	}
	p, found := locationMap[location]
	if !found {
		logrus.Errorln("location not found in map", location)
		return false
	}
	x := p[0]
	y := p[1]
	if b.Tiles[x][y] != 0 {
		logrus.Errorln("tile already occupied")
		return false
	}
	b.Tiles[x][y] = player.Number

	if b.PlayerNumberTurn == 1 {
		b.PlayerNumberTurn = 2
	} else {
		b.PlayerNumberTurn = 1
	}

	return true
}

func (b *Board) CheckForWinner() (hasWinner bool, winnerNumber int) {
	hasWinner, winnerNumber = b.checkForWinnerDiagonal()
	if hasWinner {
		return
	}

	//Check horizontal rows
	for y := 1; y <= 3; y++ {
		hasWinner, winnerNumber = b.checkForWinnerHorizontal(y)
		if hasWinner {
			return
		}
		hasWinner, winnerNumber = b.checkForWinnerVertical(y)
		if hasWinner {
			return
		}

	}
	return
}

func (b *Board) checkForWinnerHorizontal(y int) (hasWinner bool, winnerNumber int) {
	if b.Tiles[1][y] == b.Tiles[2][y] && b.Tiles[2][y] == b.Tiles[3][y] && b.Tiles[3][y] != 0 {
		return true, b.Tiles[1][y]
	}
	return false, 0
}
func (b *Board) checkForWinnerVertical(x int) (hasWinner bool, winnerNumber int) {
	if b.Tiles[x][1] == b.Tiles[x][2] && b.Tiles[x][2] == b.Tiles[x][3] && b.Tiles[x][3] != 0 {
		return true, b.Tiles[x][1]
	}
	return false, 0
}

func (b *Board) checkForWinnerDiagonal() (hasWinner bool, winnerNumber int) {
	if b.Tiles[1][1] == b.Tiles[2][2] && b.Tiles[2][2] == b.Tiles[3][3] && b.Tiles[3][3] != 0 {
		return true, b.Tiles[1][1]
	}
	if b.Tiles[1][3] == b.Tiles[2][2] && b.Tiles[2][2] == b.Tiles[3][1] && b.Tiles[3][1] != 0 {
		return true, b.Tiles[1][3]
	}
	return false, 0
}

func (b *Board) CheckForAvailableMoves() bool {
	for x := 1; x <= b.LastRow; x++ {
		for y := 1; y <= b.LastRow; y++ {
			if b.Tiles[x][y] == 0 {
				return true
			}
		}
	}
	return false
}

func (b *Board) Render() [][]int {
	result := make([][]int, 3)
	for x := 0; x < b.LastRow; x++ {
		result[x] = make([]int, 3)
		for y := 0; y < b.LastRow; y++ {
			result[x][y] = b.Tiles[x+1][y+1]
		}
	}
	return result
}
