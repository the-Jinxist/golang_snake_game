package game

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/the-Jinxist/golang_snake_game/internal"
)

type Tick struct{}

type TriggerNextLevel struct{}
type GameStartConfig struct {
	Rows           int
	Columns        int
	Pillars        []Position
	IsWalled       bool
	Level          int
	IsFinalLevel   bool
	ScoreThreshold int
	Scoring        int
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
		IsWalled:       true,
		ScoreThreshold: 100,
		ScoreService:   internal.GetScoreService(),
		SessionManager: internal.GetSessionManager(),
	}
}

func Level1GameConfig() GameStartConfig {
	return GameStartConfig{
		Rows:           40,
		Columns:        25,
		ScoreThreshold: 500,
		Level:          1,
		ScoreService:   internal.GetScoreService(),
		SessionManager: internal.GetSessionManager(),
	}
}

func Level2GameConfig() GameStartConfig {
	return GameStartConfig{
		Rows:           45,
		Columns:        30,
		ScoreThreshold: 1000,
		Level:          2,
		ScoreService:   internal.GetScoreService(),
		SessionManager: internal.GetSessionManager(),
	}
}

func Level3GameConfig() GameStartConfig {
	return GameStartConfig{
		Rows:           50,
		Columns:        35,
		ScoreThreshold: 1500,
		Level:          3,
		ScoreService:   internal.GetScoreService(),
		SessionManager: internal.GetSessionManager(),
	}
}

func Level4GameConfig() GameStartConfig {
	return GameStartConfig{
		Rows:           55,
		Columns:        40,
		ScoreThreshold: 2500,
		Level:          4,
		ScoreService:   internal.GetScoreService(),
		SessionManager: internal.GetSessionManager(),
	}
}

func Level5GameConfig() GameStartConfig {
	return GameStartConfig{
		Rows:           60,
		Columns:        45,
		ScoreThreshold: 5000,
		Level:          5,
		IsFinalLevel:   true,
		ScoreService:   internal.GetScoreService(),
		SessionManager: internal.GetSessionManager(),
	}
}
