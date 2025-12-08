package game

import (
	"fmt"
	"math/rand"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/the-Jinxist/golang_snake_game/utils"
)

var _ tea.Model = &GameModel{}

type Direction int

type Position struct {
	X int
	Y int
}

const (
	Up Direction = iota
	Down
	Left
	Right
)

type GameModel struct {
	Rows    int
	Columns int
	Snake   []Position

	Food      Position
	Direction Direction
}

func InitalGameModel() *GameModel {

	rows := 50
	column := 30

	gameMod := &GameModel{
		Rows:    rows,
		Columns: column,
		Snake: []Position{
			{
				X: rows / 2,
				Y: column / 2,
			},
		},
		Direction: Right,
	}

	return gameMod
}

func (g *GameModel) isSnake(x int, y int) bool {
	isSnake := false
	for _, snakeBody := range g.Snake {
		if snakeBody.X == x && snakeBody.Y == y {
			isSnake = true
		}
	}

	return isSnake
}

func (g *GameModel) isFood(x int, y int) bool {
	food := g.Food
	return x == food.X && y == food.Y
}

func (g *GameModel) directionToPosition(direction Direction) Position {
	position := Position{
		Y: 0, X: 1,
	}

	switch direction {
	case Up:
		position = Position{Y: -1, X: 0}
	case Down:
		position = Position{Y: 1, X: 0}
	case Left:
		position = Position{Y: 0, X: -1}
	}

	return position
}

func (g *GameModel) isWall(x int, y int) bool {
	return x > g.Rows || y > g.Columns || x < 0 || y < 0
}

// Init implements tea.Model.
func (g *GameModel) Init() tea.Cmd {
	rand.NewSource(time.Now().UnixNano())
	g.instantiateFood()
	return tea.Batch(g.Tick())
}

func (g *GameModel) instantiateFood() {
	randomX := rand.Intn(g.Rows)
	randomY := rand.Intn(g.Columns)

	for g.isSnake(randomX, randomY) {
		randomX = rand.Intn(g.Rows)
		randomY = rand.Intn(g.Columns)
	}

	g.Food = Position{
		X: randomX,
		Y: randomY,
	}
}

// Update implements tea.Model.
func (g *GameModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		input := msg.String()
		if utils.KeyMatchesInput(input, utils.KeyUp) {
			g.Direction = Up
		}

		if utils.KeyMatchesInput(input, utils.KeyRight) {
			g.Direction = Right
		}

		if utils.KeyMatchesInput(input, utils.KeyDown) {
			g.Direction = Down
		}

		if utils.KeyMatchesInput(input, utils.KeyLeft) {
			g.Direction = Left
		}

		if utils.KeyMatchesInput(input, utils.Space) {
			fmt.Println("Paused")
		}

		return g, nil

	case Tick:
		g.moveSnake()
		return g, tea.Batch(g.Tick())
	default:
		return g, tea.Batch(g.Tick())

	}

}

func (g *GameModel) Tick() tea.Cmd {
	return tea.Tick(time.Second/2, func(t time.Time) tea.Msg {
		return Tick{}
	})

}

func (g *GameModel) moveSnake() {

	pos := g.directionToPosition(g.Direction)

	currentSnakeHead := g.Snake[0]

	if g.isWall(currentSnakeHead.X, currentSnakeHead.Y) {
		fmt.Println("Game over")
	}

	movingToX := currentSnakeHead.X + pos.X
	movingToY := currentSnakeHead.Y + pos.Y

	if g.isFood(currentSnakeHead.X, currentSnakeHead.Y) {
		g.instantiateFood()
		newSnakeHead := Position{
			X: movingToX,
			Y: movingToY,
		}

		g.Snake = append([]Position{newSnakeHead}, g.Snake...)

	}

	newSnakeHead := Position{
		X: movingToX,
		Y: movingToY,
	}

	g.Snake = g.Snake[:len(g.Snake)-1]
	g.Snake = append([]Position{newSnakeHead}, g.Snake...)
}

// View implements tea.Model.
func (g *GameModel) View() string {

	cellStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#CCCCCC"))
	var output string
	for i := range g.Columns {
		for j := range g.Rows {

			if g.isSnake(j, i) {
				output += cellStyle.Render(FilledCell)
			} else if g.isFood(j, i) {
				output += cellStyle.Foreground(lipgloss.Color("205")).Render(FilledCell)
			} else {
				output += cellStyle.Render(EmptyCell)
			}

		}
		output += "\n"
	}

	output = lipgloss.NewStyle().Margin(5).Border(lipgloss.DoubleBorder(), true).Render(output)
	return output
}
