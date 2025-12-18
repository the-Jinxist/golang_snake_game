package views

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Mode int

const (
	ModeMenu = Mode(iota)
	ModeGame
	ModeLeaderboard
	ModeGameOver
)

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
