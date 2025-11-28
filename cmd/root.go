/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/the-Jinxist/golang_snake_game/ui"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "super_snake",
	Short: "The best terminal snake game written in Go",
	Long:  `Run the super_snake command to start playing the classic snake game in your terminal!`,
	Run: func(cmd *cobra.Command, args []string) {
		p := tea.NewProgram(ui.InitalModel())
		if _, err := p.Run(); err != nil {
			log.Fatal(err)
		}

		os.Exit(1)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
