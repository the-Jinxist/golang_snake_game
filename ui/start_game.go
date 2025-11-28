package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type StartGameModel struct {
	choices []string // items on the to-do list
	cursor  int
}

func InitalModel() StartGameModel {
	return StartGameModel{
		choices: []string{"Start Game", "Exit"},
		cursor:  0,
	}
}

func (m StartGameModel) Init() tea.Cmd {
	return nil
}

func (m StartGameModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "A", "w":
			nextCursor := m.cursor - 1
			if nextCursor < 0 {
				nextCursor = 0
			}

			m.cursor = nextCursor
		case "B", "s":
			nextCursor := m.cursor + 1
			if nextCursor > 1 {
				nextCursor = 1
			}

			m.cursor = nextCursor
		case "enter", " ":
			if m.cursor == 1 {
				return m, tea.Quit
			}
		}

	}
	return m, nil
}

func (m StartGameModel) View() string {

	var style = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7D56F4")).
		Align(lipgloss.Center).
		Padding(5).
		Border(lipgloss.BlockBorder(), true, true, true, true).
		Width(80).
		Height(5)

	title := "Super Snake"
	options := fmt.Sprint(
		"\n\n\n",
	)

	for index, value := range m.choices {

		prefix := ""
		if index == m.cursor {
			prefix = "> "
		}

		options += fmt.Sprintf("\n%s%s", prefix, value)
	}
	return style.Render(title + options)
}
