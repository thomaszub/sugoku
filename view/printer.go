package view

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/thomaszub/sugoku/game"
)

type Position struct {
	Row uint8
	Col uint8
}

type Printer struct {
	Color lipgloss.Color
}

func (p *Printer) PrintBoard(board game.Board, pos Position) string {
	rows := []string{}
	for row := uint8(0); row <= 2; row++ {
		row := printRow(board, row, pos)
		rows = append(rows, row)
	}
	s := lipgloss.JoinVertical(lipgloss.Left, rows...)
	return lipgloss.NewStyle().Margin(2).Render(s)
}

func printRow(board game.Board, row uint8, pos Position) string {
	blocks := []string{}
	for col := uint8(0); col <= 2; col++ {
		block := printBlock(board, row, col, pos)
		blocks = append(blocks, block)
	}
	return lipgloss.JoinHorizontal(lipgloss.Top, blocks...)
}

func printBlock(board game.Board, rowStart, colStart uint8, pos Position) string {
	rows := []string{}
	for row := rowStart * 3; row < rowStart*3+3; row++ {
		cols := []string{}
		for col := colStart * 3; col < colStart*3+3; col++ {
			val := board[row][col]
			if row == pos.Row && col == pos.Col {
				cols = append(cols, PositionStyle.Render(fmt.Sprintf("%d", val)))
			} else {
				cols = append(cols, fmt.Sprintf("%d", val))
			}
		}
		col := strings.Join(cols, " ")
		rows = append(rows, col)
	}
	return BorderBlockStyle[rowStart][colStart].Render(strings.Join(rows, "\n"))
}
