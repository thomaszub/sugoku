package view

import "github.com/charmbracelet/lipgloss"

var PositionStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#FAFAFA")).
	Background(lipgloss.Color("#7D56F4"))

var border = lipgloss.NewStyle().
	BorderStyle(lipgloss.OuterHalfBlockBorder()).
	BorderForeground(lipgloss.Color("#7D56F4"))

var BorderBlockStyle = [][]lipgloss.Style{
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
