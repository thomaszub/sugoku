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
