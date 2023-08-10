package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/thomaszub/sugoku/view"
)

type model struct {
	board    Board
	position Position
}

func initialModel() model {
	board := baseBoard
	return model{
		board:    board,
		position: Position{4, 4},
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "left", "h":
			if m.position.Col > 0 {
				m.position.Col--
			}
		case "up", "k":
			if m.position.Row > 0 {
				m.position.Row--
			}
		case "down", "j":
			if m.position.Row < 8 {
				m.position.Row++
			}
		case "right", "l":
			if m.position.Col < 8 {
				m.position.Col++
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	s := "sugoku\n"

	s += printBoard(m.board, m.position)

	s += "\n\nPress q to quit.\n"
	return s
}

func printBoard(board Board, pos Position) string {
	rows := []string{}
	for row := uint8(0); row <= 2; row++ {
		row := printRow(board, row, pos)
		rows = append(rows, row)
	}
	s := lipgloss.JoinVertical(lipgloss.Left, rows...)
	return lipgloss.NewStyle().Margin(2).Render(s)
}

func printRow(board Board, row uint8, pos Position) string {
	blocks := []string{}
	for col := uint8(0); col <= 2; col++ {
		block := printBlock(board, row, col, pos)
		blocks = append(blocks, block)
	}
	return lipgloss.JoinHorizontal(lipgloss.Top, blocks...)
}

func printBlock(board Board, rowStart, colStart uint8, pos Position) string {
	rows := []string{}
	for row := rowStart * 3; row < rowStart*3+3; row++ {
		cols := []string{}
		for col := colStart * 3; col < colStart*3+3; col++ {
			val := board[row][col]
			if row == pos.Row && col == pos.Col {
				cols = append(cols, view.PositionStyle.Render(fmt.Sprintf("%d", val)))
			} else {
				cols = append(cols, fmt.Sprintf("%d", val))
			}
		}
		col := strings.Join(cols, " ")
		rows = append(rows, col)
	}
	return view.BorderBlockStyle[rowStart][colStart].Render(strings.Join(rows, "\n"))
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("sugoku crashed for some reason: %v", err)
		os.Exit(1)
	}
}
