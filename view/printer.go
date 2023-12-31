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
	positionStyle    lipgloss.Style
	guessedCellStyle lipgloss.Style
	borderBlockStyle [][]lipgloss.Style
}

func NewPrinter(foregroundColor, backgroundColor lipgloss.Color) Printer {
	return Printer{
		positionStyle:    positionStyle(foregroundColor, backgroundColor),
		guessedCellStyle: guessedCellStyle(foregroundColor),
		borderBlockStyle: borderBlockStyle(backgroundColor),
	}
}

func (p *Printer) PrintBoard(board game.Board, pos Position) string {
	rows := []string{}
	for row := uint8(0); row <= 2; row++ {
		row := p.printRow(board, row, pos)
		rows = append(rows, row)
	}
	s := lipgloss.JoinVertical(lipgloss.Left, rows...)
	return lipgloss.NewStyle().Margin(2).Render(s)
}

func (p *Printer) printRow(board game.Board, row uint8, pos Position) string {
	blocks := []string{}
	for col := uint8(0); col <= 2; col++ {
		block := p.printBlock(board, row, col, pos)
		blocks = append(blocks, block)
	}
	return lipgloss.JoinHorizontal(lipgloss.Top, blocks...)
}

func (p *Printer) printBlock(board game.Board, rowStart, colStart uint8, pos Position) string {
	rows := []string{}
	for row := rowStart * 3; row < rowStart*3+3; row++ {
		cols := []string{}
		for col := colStart * 3; col < colStart*3+3; col++ {
			cell := board[row][col]
			cellS := fmt.Sprintf("%d", cell.Value)
			if cell.Value == 0 {
				cellS = " "
			}
			if !cell.Initial {
				cellS = p.guessedCellStyle.Render(cellS)
			}
			if row == pos.Row && col == pos.Col {
				cols = append(cols, p.positionStyle.Render(cellS))
			} else {
				cols = append(cols, cellS)
			}
		}
		col := strings.Join(cols, " ")
		rows = append(rows, col)
	}
	return p.borderBlockStyle[rowStart][colStart].Render(strings.Join(rows, "\n"))
}
