package views

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Mode int

const (
	ModeMenu = Mode(iota)
	ModeGame
	ModeGame1
	ModeGame2
	ModeGame3
	ModeGame4
	ModeGame5
	ModeLeaderboard
	ModeGameOver
	ModeGameCompleted
)

func NextLevelModeFromCurrent(level int) Mode {

	if level == 0 {
		return ModeGame1
	}

	if level == 1 {
		return ModeGame2
	}

	if level == 2 {
		return ModeGame3
	}

	if level == 3 {
		return ModeGame4
	}

	if level == 4 {
		return ModeGame5
	}

	return ModeGame1
}

type SwitchModeMsg struct {
	Target Mode
}

type ExitGameMsg struct{}

func ClearScreen() tea.Cmd {
	return func() tea.Msg {
		return tea.ClearScreen()
	}
}

func SwitchModeCmd(target Mode) tea.Cmd {
	return func() tea.Msg {
		return SwitchModeMsg{
			Target: target,
		}
	}
}

func ExitGameCmd() tea.Cmd {
	return func() tea.Msg {
		return ExitGameMsg{}
	}
}
