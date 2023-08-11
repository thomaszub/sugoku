package game

import "math/rand"

func removeCells(board Board, count int) Board {
	for i := 0; i < count; i++ {
		board = removeCell(board)
	}
	return board
}

func removeCell(board Board) Board {
	row := rand.Intn(9)
	col := rand.Intn(9)
	cell := board[row][col]
	if cell.Value != 0 {
		board[row][col] = Cell{
			Value:   0,
			Initial: false,
		}
		return board
	}
	return removeCell(board)
}
