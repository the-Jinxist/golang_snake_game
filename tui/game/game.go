package game

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/Broderick-Westrope/charmutils"
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

type Food struct {
	Position
	BigFish bool
}

type GameModel struct {
	Config GameStartConfig
	Snake  []Position

	Food       Food
	Direction  Direction
	Score      int
	IsGameOver bool
	isPaused   bool
}

func InitalGameModel(gameConfig GameStartConfig) *GameModel {
	gameMod := &GameModel{
		Snake: []Position{
			{
				X: gameConfig.Rows / 2,
				Y: gameConfig.Columns / 2,
			},
		},
		Config:    gameConfig,
		Direction: Right,
		Score:     0,
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

func (g *GameModel) hasHitWall(x int, y int) bool {
	return g.isOutOfBounds(x, y) && g.Config.IsWalled
}

func (g *GameModel) isOutOfBounds(x int, y int) bool {
	return x > g.Config.Rows || y > g.Config.Columns || x < 0 || y < 0
}

// Init implements tea.Model.
func (g *GameModel) Init() tea.Cmd {
	rand.NewSource(time.Now().UnixNano())
	g.instantiateFood()
	return tea.Batch(g.Tick())
}

func (g *GameModel) instantiateFood() {

	rows := g.Config.Rows
	columns := g.Config.Columns

	randomX := rand.Intn(rows)
	randomY := rand.Intn(columns)

	for g.isSnake(randomX, randomY) {
		randomX = rand.Intn(rows)
		randomY = rand.Intn(columns)
	}

	g.Food = Food{
		Position: Position{
			X: randomX,
			Y: randomY,
		},
	}
}

// Update implements tea.Model.
func (g *GameModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		input := msg.String()
		if utils.KeyMatchesInput(input, utils.Space) {
			g.isPaused = !g.isPaused
		}

		if g.isPaused {
			return g, nil
		}

		if utils.KeyMatchesInput(input, utils.KeyUp) {
			if g.Direction == Left || g.Direction == Right {
				g.Direction = Up
			}
		}

		if utils.KeyMatchesInput(input, utils.KeyRight) {
			if g.Direction == Up || g.Direction == Down {
				g.Direction = Right
			}
		}

		if utils.KeyMatchesInput(input, utils.KeyDown) {
			if g.Direction == Left || g.Direction == Right {
				g.Direction = Down
			}
		}

		if utils.KeyMatchesInput(input, utils.KeyLeft) {
			if g.Direction == Up || g.Direction == Down {
				g.Direction = Left
			}
		}

		return g, nil

	case Tick:

		if !g.isPaused {
			g.moveSnake()
		}

		return g, tea.Batch(g.Tick())
	default:
		return g, tea.Batch(g.Tick())

	}

}

func (g *GameModel) Tick() tea.Cmd {
	return tea.Tick(time.Second/4, func(t time.Time) tea.Msg {
		return Tick{}
	})
}

func (g *GameModel) moveSnake() {

	pos := g.directionToPosition(g.Direction)

	currentSnakeHead := g.Snake[0]

	//If snake is out of bounds
	if g.hasHitWall(currentSnakeHead.X, currentSnakeHead.Y) {
		g.IsGameOver = true
	}

	movingToX := currentSnakeHead.X + pos.X
	movingToY := currentSnakeHead.Y + pos.Y

	// If snakes eats itself
	if g.isSnake(movingToX, movingToY) {
		g.IsGameOver = true
	}

	if g.IsGameOver {
		return
	}

	if g.isFood(currentSnakeHead.X, currentSnakeHead.Y) {
		g.instantiateFood()
		newSnakeHead := Position{
			X: movingToX,
			Y: movingToY,
		}

		g.increaseScore()

		g.Snake = append([]Position{newSnakeHead}, g.Snake...)

	}

	//Snap snake back to the opp. side of stage if he goes out of bounds
	if g.isOutOfBounds(movingToX, movingToY) {
		if movingToX > g.Config.Rows {
			movingToX = 0
		}

		if movingToX < 0 {
			movingToX = g.Config.Rows
		}

		if movingToY > g.Config.Columns {
			movingToY = 0
		}

		if movingToY < 0 {
			movingToY = g.Config.Columns
		}
	}

	newSnakeHead := Position{
		X: movingToX,
		Y: movingToY,
	}

	g.Snake = g.Snake[:len(g.Snake)-1]
	g.Snake = append([]Position{newSnakeHead}, g.Snake...)
}

func (g *GameModel) increaseScore() {
	g.Score += g.Config.Scoring

	go func() {
		g.Config.ScoreService.SetCurrentScore(context.Background(), g.Score)
	}()

}

// View implements tea.Model.
func (g *GameModel) View() string {

	cellStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#CCCCCC"))
	var output string
	for i := range g.Config.Columns {
		for j := range g.Config.Rows {

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

	output = lipgloss.NewStyle().Border(lipgloss.ASCIIBorder(), g.Config.IsWalled).Render(output)

	output += "\n"

	if g.isPaused {
		output += lipgloss.NewStyle().
			AlignHorizontal(lipgloss.Center).
			Render(fmt.Sprintf("[ PAUSED ]. Your score: %d. Press SPACE to resume!", g.Score))
	} else {
		output += lipgloss.NewStyle().
			AlignHorizontal(lipgloss.Center).
			Render(fmt.Sprintf("Your score: %d. Press SPACE to pause!", g.Score))
	}

	if g.IsGameOver {
		gameOverMessage := GameOver
		gameOverMessage += "\n"
		gameOverMessage += lipgloss.NewStyle().
			AlignHorizontal(lipgloss.Center).
			Render("Press EXIT or HOLD to continue")
		output, _ = charmutils.OverlayCenter(output, gameOverMessage, false)
	}

	return output
}
