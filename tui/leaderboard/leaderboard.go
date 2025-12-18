package leaderboard

import (
	"context"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/the-Jinxist/golang_snake_game/internal"
	"github.com/the-Jinxist/golang_snake_game/tui/views"
)

const leaderboardTitle = `
_      _____   ___  ______  _____  ______  ______  _____   ___  ______ ______ 
| |    |  ___| / _ \ |  _  \|  ___| | ___ \ | ___ \|  _  | / _ \ | ___ \|  _  \
| |    | |__  / /_\ \| | | || |__   | |_/ / | |_/ /| | | |/ /_\ \| |_/ /| | | |
| |    |  __| |  _  || | | ||  __|  |    /  | ___ \| | | ||  _  ||    / | | | |
| |____| |___ | | | || |/ / | |___  | |\ \  | |_/ /\ \_/ /| | | || |\ \ | |/ / 
\_____/\____/ \_| |_/|___/  \____/  \_| \_| \____/  \___/ \_| |_/\_| \_||___/

`

var (
	_ tea.Model = &Leaderboard{}

	highscoreStyle = lipgloss.NewStyle().MarginLeft(2).PaddingLeft(1).PaddingRight(1).Background(lipgloss.Color("#87CEFA"))
	scoreStyle     = lipgloss.NewStyle().Bold(true)
	descStyle      = lipgloss.NewStyle().Italic(true).Foreground(lipgloss.Color("#DDDDDD"))
)

type Leaderboard struct {
	Scores []internal.Score
	Config LeaderboardConfig
}

func NewLeaderboardModel(config LeaderboardConfig) *Leaderboard {

	const defaultWidth = 20
	scores, _ := config.ScoreService.GetScores(context.Background())

	return &Leaderboard{
		Config: config,
		Scores: scores,
	}
}

// Init implements tea.Model.
func (l *Leaderboard) Init() tea.Cmd {
	return nil
}

// Update implements tea.Model.
func (l *Leaderboard) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c":
			return l, tea.Quit
		case "esc":
			return l, tea.Batch(views.ClearScreen(), views.SwitchModeCmd(views.ModeMenu))
		}
	}

	var cmd tea.Cmd
	return l, cmd
}

// View implements tea.Model.
func (l *Leaderboard) View() string {

	title := leaderboardTitle
	description := "\n\n"

	for i, value := range l.Scores {
		description += scoreStyle.Render(fmt.Sprintf("Score: %d", value.Value))
		if i == 0 {
			description += highscoreStyle.Render("Highscore")
		}

		description += "\n"
		description += descStyle.Render(fmt.Sprintf("Recorded at %s", value.CreatedAt.GoString()))

		description += "\n\n"
	}

	help := "\n\n\nPress [esc] to return back to menu screen"

	return title + description + help
}
