package menu

import (
	"context"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/the-Jinxist/golang_snake_game/internal"
	"github.com/the-Jinxist/golang_snake_game/tui/views"
)

const (
	combinedTitle = `
███████╗██╗   ██╗██████╗ ███████╗██████╗  ███████╗███╗   ██╗ █████╗ ██╗  ██╗███████╗
██╔════╝██║   ██║██╔══██╗██╔════╝██╔══██╗ ██╔════╝████╗  ██║██╔══██╗██║ ██╔╝██╔════╝
███████╗██║   ██║██████╔╝█████╗  ██████╔╝ ███████╗██╔██╗ ██║███████║█████═╝ █████╗  
╚════██║██║   ██║██╔═══╝ ██╔══╝  ██╔══██╗ ╚════██║██║╚██╗██║██╔══██║██╔═██╗ ╚════╝  
███████║╚██████╔╝██║     ███████╗██║  ██║ ███████║██║ ╚████║██║  ██║██║  ██╗███████╗
╚══════╝ ╚═════╝ ╚═╝     ╚══════╝╚═╝  ╚═╝ ╚══════╝╚═╝  ╚═══╝╚═╝  ╚═╝╚═╝  ╚═╝╚══════╝
`
)

// const (
// 	block   = " "
// 	rows    = 30
// 	columns = 20
// )

var _ tea.Model = StartGameModel{}

type StartGameModel struct {
	choices []string // items on the to-do list
	cursor  int
}

func InitalModel() StartGameModel {
	return StartGameModel{
		choices: []string{"Start Game", "Leaderboard", "Exit"},
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
			if nextCursor > 2 {
				nextCursor = 0
			}

			m.cursor = nextCursor
		case "enter", " ":
			fmt.Print("\033[H\033[2J")

			if m.cursor == 1 {
				return m, tea.Batch(views.SwitchModeCmd(views.ModeLeaderboard))
			}

			if m.cursor == 2 {
				return m, tea.Quit
			}

			return m, tea.Batch(views.SwitchModeCmd(views.ModeGame))
		}

	}
	return m, nil
}

func (m StartGameModel) View() string {
	var style = lipgloss.NewStyle().
		Bold(true).
		Align(lipgloss.Left).
		Width(30).
		Height(5)

	title := combinedTitle
	title += "\n"
	title += style.Render(fmt.Sprintf("Your current highscore is: %d", getHighScore()))

	options := ""

	for index, value := range m.choices {

		prefix := ""
		if index == m.cursor {
			prefix = "> "
		}

		options += style.Render(fmt.Sprintf("\n%s%s", prefix, value))
	}

	return title + options
}

func getHighScore() int {
	score, err := internal.GetScoreService().GetHighScore(context.Background())
	if err != nil {
		fmt.Printf("error: %s", err)
		return 0
	}

	return score.Value
}
