package game

import (
	"fmt"
	"math/rand"
)

type switchRowsRandomizer struct {
	firstRow  int
	secondRow int
}

func (r switchRowsRandomizer) randomize(board Board) Board {
	firstRow := board[r.firstRow]
	secondRow := board[r.secondRow]
	board[r.secondRow] = firstRow
	board[r.firstRow] = secondRow
	return board
}

type switchColsRandomizer struct {
	firstCol  int
	secondCol int
}

func (r switchColsRandomizer) randomize(board Board) Board {
	for i := range board {
		oldSecondCol := board[i][r.secondCol]
		board[i][r.secondCol] = board[i][r.firstCol]
		board[i][r.firstCol] = oldSecondCol
	}
	return board
}

type switchNumberRandomizer struct {
	firstNumber  uint8
	secondNumber uint8
}

func (r switchNumberRandomizer) randomize(board Board) Board {
	for i := range board {
		for j := range board[i] {
			switch board[i][j] {
			case r.firstNumber:
				board[i][j] = r.secondNumber
			case r.secondNumber:
				board[i][j] = r.firstNumber
			}
		}
	}
	return board
}

func applyRandomRandomizer(board Board) Board {
	id := rand.Intn(3)
	switch id {
	case 0:
		block := rand.Intn(3)
		return switchRowsRandomizer{
			firstRow:  3*block + rand.Intn(3),
			secondRow: 3*block + rand.Intn(3),
		}.randomize(board)
	case 1:
		block := rand.Intn(3)
		return switchColsRandomizer{
			firstCol:  3*block + rand.Intn(3),
			secondCol: 3*block + rand.Intn(3),
		}.randomize(board)
	case 2:
		return switchNumberRandomizer{
			firstNumber:  uint8(1 + rand.Intn(9)),
			secondNumber: uint8(1 + rand.Intn(9)),
		}.randomize(board)
	}
	panic(fmt.Sprintf("Random number %d not mapped to randomizer. This is a bug", id))
}
