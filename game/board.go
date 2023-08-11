package game

type Cell struct {
	Value   uint8
	Initial bool
}

type Board = [][]Cell

func NewBaseBoard() Board {
	board := make([][]Cell, len(baseBoard))
	for i, row := range baseBoard {
		board[i] = make([]Cell, len(row))
		for j, cell := range row {
			board[i][j] = Cell{
				Value:   cell,
				Initial: true,
			}
		}
	}
	return board
}

func copyBoard(board Board) Board {
	newBoard := make([][]Cell, len(board))
	for i := range board {
		newBoard[i] = make([]Cell, len(board[i]))
		copy(newBoard[i], board[i])
	}
	return newBoard
}

var baseBoard = [][]uint8{
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
