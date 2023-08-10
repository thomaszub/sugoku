package game

type Board = [][]uint8

func NewBoard() Board {
	board := copyBoard(baseBoard)
	for i := 0; i < 256; i++ {
		board = applyRandomRandomizer(board)
	}
	return board
}

func copyBoard(board Board) Board {
	newBoard := make([][]uint8, len(board))
	for i := range board {
		newBoard[i] = make([]uint8, len(board[i]))
		copy(newBoard[i], board[i])
	}
	return newBoard
}

var baseBoard = Board{
	{1, 2, 3, 4, 5, 6, 7, 8, 9},
	{4, 5, 6, 7, 8, 9, 1, 2, 3},
	{7, 8, 9, 1, 2, 3, 4, 5, 6},
	{2, 3, 1, 5, 6, 4, 8, 9, 7},
	{5, 6, 4, 8, 9, 7, 2, 3, 1},
	{8, 9, 7, 2, 3, 1, 5, 6, 4},
	{3, 1, 2, 6, 4, 5, 9, 7, 8},
	{6, 4, 5, 9, 7, 8, 3, 1, 2},
	{9, 7, 8, 3, 1, 2, 6, 4, 5},
}
