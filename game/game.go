package game

type Game struct {
	initialBoard Board
	currentBoard Board
}

func NewGame() Game {
	initialBoard := RandomizeBoard(NewBaseBoard(), 256)
	currentBoard := copyBoard(initialBoard)
	return Game{
		initialBoard: initialBoard,
		currentBoard: currentBoard,
	}
}

func (g Game) CurrentBoard() Board {
	return g.currentBoard
}
