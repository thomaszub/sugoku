package view

import "github.com/charmbracelet/lipgloss"

func positionStyle(foregroundColor, backgroundColor lipgloss.Color) lipgloss.Style {
	return lipgloss.NewStyle().
		Bold(true).
		Foreground(foregroundColor).
		Background(backgroundColor)
}

func guessedCellStyle(foregroundColor lipgloss.Color) lipgloss.Style {
	return lipgloss.NewStyle().
		Bold(true).
		Foreground(foregroundColor)
}

func borderBlockStyle(color lipgloss.Color) [][]lipgloss.Style {
	border := lipgloss.NewStyle().
		BorderStyle(lipgloss.OuterHalfBlockBorder()).
		BorderForeground(color)

	return [][]lipgloss.Style{
		{
			border.Copy().Border(lipgloss.OuterHalfBlockBorder(), true, true, false, true),
			border.Copy().Border(lipgloss.OuterHalfBlockBorder(), true, true, false, true),
			border.Copy().Border(lipgloss.OuterHalfBlockBorder(), true, true, false, true),
		},
		{
			border.Copy(),
			border.Copy(),
			border.Copy(),
		},
		{
			border.Copy().Border(lipgloss.OuterHalfBlockBorder(), false, true, true, true),
			border.Copy().Border(lipgloss.OuterHalfBlockBorder(), false, true, true, true),
			border.Copy().Border(lipgloss.OuterHalfBlockBorder(), false, true, true, true),
		},
	}
}
