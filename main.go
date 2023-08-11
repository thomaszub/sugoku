package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/thomaszub/sugoku/game"
	"github.com/thomaszub/sugoku/view"
)

type model struct {
	game     game.Game
	position view.Position
	printer  view.Printer
}

func initialModel() model {
	game := game.NewGame()
	return model{
		game:     game,
		position: view.Position{Row: 4, Col: 4},
		printer:  view.NewPrinter("#FAFAFA", "#7D56F4"),
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

	s += m.printer.PrintBoard(m.game.CurrentBoard(), m.position)

	s += "\n\nPress q to quit.\n"
	return s
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("sugoku crashed for some reason: %v", err)
		os.Exit(1)
	}
}
