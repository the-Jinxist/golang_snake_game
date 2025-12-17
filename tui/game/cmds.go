package game

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/the-Jinxist/golang_snake_game/internal"
)

type Tick struct{}
type GameStartConfig struct {
	Rows           int
	Columns        int
	Pillars        []Position
	IsWalled       bool
	Scoring        int
	ScoreService   internal.ScoreService
	SessionManager internal.SessionManager
}

type LeaderboardConfig struct {
	ScoreService   internal.ScoreService
	SessionManager internal.SessionManager
}

func TickGame() tea.Cmd {
	return func() tea.Msg {
		return Tick{}
	}
}

func DefaultGameConfig() GameStartConfig {
	return GameStartConfig{
		Rows:           30,
		Columns:        25,
		Scoring:        10,
		IsWalled:       false,
		ScoreService:   internal.GetScoreService(),
		SessionManager: internal.GetSessionManager(),
	}
}

func Level1GameConfig() GameStartConfig {
	return GameStartConfig{
		Rows:    40,
		Columns: 25,

		ScoreService:   internal.GetScoreService(),
		SessionManager: internal.GetSessionManager(),
	}
}

func Level2GameConfig() GameStartConfig {
	return GameStartConfig{
		Rows:    45,
		Columns: 30,

		ScoreService:   internal.GetScoreService(),
		SessionManager: internal.GetSessionManager(),
	}
}

func Level3GameConfig() GameStartConfig {
	return GameStartConfig{
		Rows:    50,
		Columns: 35,

		ScoreService:   internal.GetScoreService(),
		SessionManager: internal.GetSessionManager(),
	}
}

func Level4GameConfig() GameStartConfig {
	return GameStartConfig{
		Rows:    55,
		Columns: 40,

		ScoreService:   internal.GetScoreService(),
		SessionManager: internal.GetSessionManager(),
	}
}

func Level5GameConfig(
	scoreService internal.ScoreService,
	sessionManager internal.SessionManager,
) GameStartConfig {
	return GameStartConfig{
		Rows:    60,
		Columns: 45,

		ScoreService:   internal.GetScoreService(),
		SessionManager: internal.GetSessionManager(),
	}
}
