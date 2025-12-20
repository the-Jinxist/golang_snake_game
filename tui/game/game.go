package game

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/Broderick-Westrope/charmutils"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/the-Jinxist/golang_snake_game/tui/views"
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

	Food          Food
	Direction     Direction
	Score         int
	IsGameOver    bool
	IsOutOfBounds bool
	spinner       spinner.Model
	isPaused      bool
}

func InitalGameModel(gameConfig GameStartConfig) *GameModel {

	s := spinner.New()
	s.Spinner = spinner.Dot

	currentScore, err := gameConfig.ScoreService.GetCurrentScore(context.Background())
	if err != nil {
		log.Fatalf("Failed to get current score: %s", err)
	}

	gameMod := &GameModel{
		Snake: []Position{
			{
				X: gameConfig.Rows / 2,
				Y: gameConfig.Columns / 2,
			},
		},
		Config:    gameConfig,
		Direction: Right,
		Score:     currentScore,
		spinner:   s,
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
	return x > g.Config.Rows-1 || y > g.Config.Columns-1 || x < 0 || y < 0
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

		if g.IsGameOver {
			g.Config.SessionManager.DestroyCurrentSession()
			if utils.KeyMatchesInput(input, utils.Esc, utils.Space) {
				return g, tea.Batch(views.ClearScreen(), views.SwitchModeCmd(views.ModeMenu))
			}

			return g, nil
		}

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

		if !g.isPaused && !g.IsGameOver {
			g.moveSnake()
		}

		return g, tea.Batch(g.Tick(), g.spinner.Tick)
	default:
		return g, tea.Batch(g.Tick(), g.spinner.Tick)

	}

}

func (g *GameModel) hasReachedLevelThreshold() bool {
	return g.Score == g.Config.ScoreThreshold
}

func (g *GameModel) Tick() tea.Cmd {

	if g.hasReachedLevelThreshold() {

		if g.Config.Level == 5 {
			return tea.Batch(views.SwitchModeCmd(views.ModeGameCompleted))
		}

		time.Sleep(2 * time.Second)
		return tea.Batch(views.SwitchModeCmd(views.ModeGame))
	}

	return tea.Tick(time.Second/4, func(t time.Time) tea.Msg {
		return Tick{}
	})
}

func (g *GameModel) moveSnake() {
	pos := g.directionToPosition(g.Direction)

	currentSnakeHead := g.Snake[0]

	movingToX := currentSnakeHead.X + pos.X
	movingToY := currentSnakeHead.Y + pos.Y

	//If snake is out of bounds
	if g.hasHitWall(movingToX, movingToY) {
		g.IsGameOver = true
	}

	// If snakes eats itself
	if g.isSnake(movingToX, movingToY) {
		g.IsGameOver = true
	}

	if g.IsGameOver || g.hasReachedLevelThreshold() {
		return
	}

	if g.isFood(movingToX, movingToY) {
		g.instantiateFood()
		newSnakeHead := Position{
			X: movingToX,
			Y: movingToY,
		}

		g.increaseScore()

		g.Snake = append([]Position{newSnakeHead}, g.Snake...)

	}

	//Snap snake back to the opp. side of stage if he goes out of bounds
	if g.isOutOfBounds(movingToX, movingToY) ||
		g.isOutOfBounds(currentSnakeHead.X, currentSnakeHead.Y) {

		// g.IsOutOfBounds = true
		if movingToX > g.Config.Rows-1 || currentSnakeHead.X > g.Config.Rows-1 {
			movingToX = 0
		}

		if movingToX < 0 || currentSnakeHead.X < 0 {
			movingToX = g.Config.Rows - 1
		}

		if movingToY > g.Config.Columns-1 || currentSnakeHead.Y > g.Config.Columns-1 {
			movingToY = 0
		}

		if movingToY < 0 || currentSnakeHead.Y < 0 {
			movingToY = g.Config.Columns - 1
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

	if g.hasReachedLevelThreshold() {
		levelingUpMsg := "We're going up!"
		levelingUpMsg += "/n"
		levelingUpMsg += lipgloss.NewStyle().
			AlignHorizontal(lipgloss.Center).
			Render(g.spinner.View())
		output, _ = charmutils.OverlayCenter(output, levelingUpMsg, false)

	}

	if g.IsGameOver {
		gameOverMessage := gameOverMsg
		gameOverMessage += "\n"
		gameOverMessage += lipgloss.NewStyle().
			AlignHorizontal(lipgloss.Center).
			Render(fmt.Sprintf("Your final score is %d\nPress SPACE to go back to menu", g.Score))
		output, _ = charmutils.OverlayCenter(output, gameOverMessage, false)
	}

	return output
}
