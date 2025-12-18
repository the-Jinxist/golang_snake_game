package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/the-Jinxist/golang_snake_game/tui/game"
	"github.com/the-Jinxist/golang_snake_game/tui/leaderboard"
	"github.com/the-Jinxist/golang_snake_game/tui/menu"
	"github.com/the-Jinxist/golang_snake_game/tui/views"
)

type SuperSnake struct {
	child tea.Model

	width  int
	height int
}

func NewModel() *SuperSnake {
	return &SuperSnake{
		child: menu.InitalModel(),
	}
}

func (s *SuperSnake) setChild(mode views.Mode) {
	switch mode {
	case views.ModeGame:
		s.child = game.InitalGameModel(game.DefaultGameConfig())
		return
	case views.ModeLeaderboard:
		s.child = leaderboard.NewLeaderboardModel(
			leaderboard.DefaultLeaderboardConfig(),
		)
		return
	}

	s.child = menu.InitalModel()
}

func (s *SuperSnake) Init() tea.Cmd {
	return s.initChild()
}

// Update is called when a message is received. Use it to inspect messages
// and, in response, update the model and/or send a command.
func (s *SuperSnake) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":

			return s, tea.Quit
		}
	case views.SwitchModeMsg:
		s.setChild(msg.Target)
	}

	var cmd tea.Cmd
	s.child, cmd = s.child.Update(msg)
	return s, cmd
}

// View renders the program's UI, which is just a string. The view is
// rendered after every Update.
func (s *SuperSnake) View() string {
	return s.child.View()
}

func (m *SuperSnake) initChild() tea.Cmd {
	var cmds []tea.Cmd
	cmd := m.child.Init()
	cmds = append(cmds, cmd)
	m.child, cmd = m.child.Update(tea.WindowSizeMsg{Width: 100, Height: 40})
	cmds = append(cmds, cmd)
	return tea.Batch(cmds...)
}
