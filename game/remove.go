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
	val := board[row][col]
	if val != 0 {
		board[row][col] = 0
		return board
	}
	return removeCell(board)
}
