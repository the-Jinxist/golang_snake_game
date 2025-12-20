package game

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/the-Jinxist/golang_snake_game/tui/views"
	"github.com/the-Jinxist/golang_snake_game/utils"
)

var _ tea.Model = &GameCompleted{}

type GameCompleted struct {
	Score int
}

func NewGameCompletedModel(score int) *GameCompleted {
	return &GameCompleted{
		Score: score,
	}
}

// Init implements tea.Model.
func (g *GameCompleted) Init() tea.Cmd {
	return nil
}

// Update implements tea.Model.
func (g *GameCompleted) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		input := msg.String()

		if utils.KeyMatchesInput(input, utils.Esc, utils.Space) {
			return g, tea.Batch(views.ClearScreen(), views.SwitchModeCmd(views.ModeMenu))
		}
	default:
		return g, nil
	}

	return g, nil
}

// View implements tea.Model.
func (g *GameCompleted) View() string {
	gameCompletedMsg := "You are officially a"
	gameCompletedMsg += "/n"
	gameCompletedMsg += superSnakeMsg
	gameCompletedMsg += "/n"
	gameCompletedMsg += fmt.Sprintf("Your final score is %d\nPress SPACE to go back to menu", g.Score)

	return gameCompletedMsg
}
