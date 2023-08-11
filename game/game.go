package game

type Game struct {
	initialBoard Board
	currentBoard Board
}

func NewGame() Game {
	initialBoard := RandomizeBoard(NewBaseBoard(), 256)
	currentBoard := removeCells(copyBoard(initialBoard), 40)
	return Game{
		initialBoard: initialBoard,
		currentBoard: currentBoard,
	}
}

func (g Game) CurrentBoard() Board {
	return g.currentBoard
}

func (g Game) SetNumber(row, col, val uint8) {
	cell := g.currentBoard[row][col]
	if cell.Initial {
		return
	}
	g.currentBoard[row][col].Value = val
}

func (g Game) DeleteNumber(row, col uint8) {
	cell := g.currentBoard[row][col]
	if cell.Initial {
		return
	}
	g.currentBoard[row][col].Value = 0
}
