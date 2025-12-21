███████╗██╗   ██╗██████╗ ███████╗██████╗  ███████╗███╗   ██╗ █████╗ ██╗  ██╗███████╗
██╔════╝██║   ██║██╔══██╗██╔════╝██╔══██╗ ██╔════╝████╗  ██║██╔══██╗██║ ██╔╝██╔════╝
███████╗██║   ██║██████╔╝█████╗  ██████╔╝ ███████╗██╔██╗ ██║███████║█████═╝ █████╗  
╚════██║██║   ██║██╔═══╝ ██╔══╝  ██╔══██╗ ╚════██║██║╚██╗██║██╔══██║██╔═██╗ ╚════╝  
███████║╚██████╔╝██║     ███████╗██║  ██║ ███████║██║ ╚████║██║  ██║██║  ██╗███████╗
╚══════╝ ╚═════╝ ╚═╝     ╚══════╝╚═╝  ╚═╝ ╚══════╝╚═╝  ╚═══╝╚═╝  ╚═╝╚═╝  ╚═╝╚══════╝

A classic Snake game implementation written in Go with a beautiful Terminal User Interface (TUI). Play multiple difficulty levels, compete on the leaderboard, and enjoy smooth gameplay in your terminal.

## 📋 Table of Contents

- [Features](#features)
- [Project Structure](#project-structure)
- [Installation](#installation)
- [Usage](#usage)
- [Architecture](#architecture)
- [Component Documentation](#component-documentation)
- [Game Mechanics](#game-mechanics)
- [Database Schema](#database-schema)
- [Key Controls](#key-controls)

---

## ✨ Features

- **Classic Snake Gameplay**: Navigate your snake to eat food and grow longer
- **Progressive Difficulty Levels**: 5 distinct levels with increasing complexity
- **Obstacle Pillars**: Navigate through static obstacles that increase in difficulty
- **Score Tracking**: Persistent high score storage using SQLite
- **Leaderboard**: View all-time high scores
- **Session Management**: Track individual game sessions
- **Terminal UI**: Beautiful ASCII art and styled terminal interface using Charm.sh
- **Pause Support**: Pause and resume gameplay
- **Game Over Detection**: Intelligent collision detection

---

## 📁 Project Structure

```
golang_snake_game/
├── main.go                 # Application entry point
├── go.mod                  # Go module definition
├── go.sum                  # Go dependencies checksum
├── LICENSE                 # Project license
├── cmd/
│   └── root.go            # Cobra CLI root command
├── internal/
│   ├── internal.go        # Global configuration and initialization
│   ├── db.go              # Database setup and management
│   ├── session.go         # Session management (in-memory)
│   └── score.go           # Score service (persistence layer)
├── tui/
│   ├── super_snake.go     # Main TUI model (Bubble Tea)
│   ├── menu/
│   │   └── start_game.go  # Main menu screen
│   ├── game/
│   │   ├── game.go        # Core game logic
│   │   ├── game_over.go   # Game over screen
│   │   ├── cmds.go        # Game commands/messages
│   │   ├── styles.go      # Game styling (Lipgloss)
│   │   ├── pillars.go     # Obstacle definitions
│   │   ├── pillar_two.go  # Additional pillar data
│   │   └── pillars.go     # Level pillar configurations
│   ├── leaderboard/
│   │   ├── leaderboard.go # Leaderboard display
│   │   └── cmd.go         # Leaderboard commands
│   └── views/
│       └── mode.go        # View mode definitions
└── utils/
    └── keys.go            # Keyboard input mappings
```

---

## 🚀 Installation

### Prerequisites

- Go 1.24.2 or higher
- Terminal with support for ANSI colors

### Build from Source

```bash
# Clone the repository
git clone https://github.com/the-Jinxist/golang_snake_game.git
cd golang_snake_game

# Install dependencies
go mod download

# Build the executable
go build -o super_snake

# Run the game
./super_snake
```

### Direct Run

```bash
go run main.go
```

---

## 💻 Usage

### Starting the Game

```bash
./super_snake
```

### Main Menu

When you launch the game, you'll see the main menu with three options:
- **Start Game**: Begin playing at level 1
- **Leaderboard**: View top high scores
- **Exit**: Quit the game

### In-Game Controls

| Action | Keys |
|--------|------|
| Move Up | `W`, `A` (arrow up) |
| Move Down | `S`, `B` (arrow down) |
| Move Left | `A`, `D` (arrow left) |
| Move Right | `D`, `C` (arrow right) |
| Pause Game | `Space` |
| Quit Game | `Ctrl+C`, `Q` |

---

## 🏗️ Architecture

### High-Level Overview

The application follows a **layered architecture** pattern:

```
┌─────────────────────────────────┐
│         Main Entry Point        │  (main.go)
│      (Cobra CLI Framework)      │
└────────────────┬────────────────┘
                 │
┌─────────────────▼────────────────┐
│      TUI Layer (Bubble Tea)      │
│    - SuperSnake (root model)     │
│    - Game Logic                  │
│    - Menu & Leaderboard          │
└────────────────┬────────────────┘
                 │
┌─────────────────▼────────────────┐
│     Internal Services Layer      │
│   - ScoreService                 │
│   - SessionManager               │
│   - Database Management          │
└────────────────┬────────────────┘
                 │
┌─────────────────▼────────────────┐
│      Data Layer (SQLite)         │
│   - Score Persistence            │
└─────────────────────────────────┘
```

### Key Design Patterns

- **Model-View-Update (MVU)**: Implemented via Bubble Tea framework
- **Dependency Injection**: Services injected through configurations
- **Interface Segregation**: `ScoreService` and `SessionManager` as interfaces
- **State Management**: Message-based event handling

---

## 📚 Component Documentation

### 1. **Main Entry Point** (`main.go`)

Entry point that initializes the Cobra CLI framework and launches the Bubble Tea application.

```go
func main() {
    cmd.Execute()
}
```

### 2. **Command Layer** (`cmd/root.go`)

**Purpose**: Cobra command definition for the `super_snake` CLI tool.

**Key Components**:
- `rootCmd`: Base command that starts the game
- Initializes database and config via `internal.IntializeConfigs()`
- Launches TUI using `tea.NewProgram()`

**Responsibilities**:
- Initialize application state
- Handle CLI arguments
- Launch the Bubble Tea program

### 3. **TUI Root Model** (`tui/super_snake.go`)

**Type**: `SuperSnake`

**Purpose**: Central routing model that manages view transitions between menu, game, and leaderboard.

**Key Methods**:
- `NewModel()`: Creates the root model with initial menu
- `setChild(mode views.Mode)`: Routes to appropriate child view based on mode
- `Update()`: Handles mode switching messages
- `Init()`: Initializes the application

**View Modes**:
- `ModeMenu`: Main menu
- `ModeGame`: Default game (level 1)
- `ModeGame1-5`: Progressive difficulty levels
- `ModeLeaderboard`: High scores display
- `ModeGameCompleted`: Game over screen

### 4. **Game Model** (`tui/game/game.go`)

**Type**: `GameModel`

**Purpose**: Core game logic implementing the snake game mechanics.

**Key Structures**:

```go
type GameModel struct {
    Config GameStartConfig      // Game configuration
    Snake  []Position          // Snake body segments
    Food   Food                // Current food location
    Direction Direction        // Current snake direction
    Score  int                 // Current score
    IsGameOver bool            // Game state
    IsOutOfBounds bool         // Boundary collision
    spinner spinner.Model      // UI spinner
    isPaused bool              // Pause state
}

type Position struct {
    X int  // Horizontal position
    Y int  // Vertical position
}

type Food struct {
    Position Position  // Food location
    BigFish bool       // Special food flag
}
```

**Game Mechanics**:
1. Snake moves in current direction each tick
2. Check collisions with food, walls, pillars, self
3. Food regenerated at random position
4. Score increases when food is eaten
5. Snake grows by one segment per food eaten
6. Game over conditions: out of bounds, self-collision

**Key Methods**:
- `InitalGameModel()`: Initialize game with configuration
- `Update()`: Process input and game state
- `View()`: Render game board
- `isSnake()`: Check if position is part of snake body
- `moveSnake()`: Move snake in current direction
- `spawnFood()`: Generate new food position

### 5. **Menu Model** (`tui/menu/start_game.go`)

**Type**: `StartGameModel`

**Purpose**: Main menu interface for game selection and navigation.

**Features**:
- Displays current high score
- Menu options: Start Game, Leaderboard, Exit
- Keyboard navigation with cursor selection

**Key Methods**:
- `InitalModel()`: Create menu with default options
- `Update()`: Handle menu navigation
- `View()`: Render menu with ASCII art title

### 6. **Leaderboard Model** (`tui/leaderboard/leaderboard.go`)

**Type**: `Leaderboard`

**Purpose**: Display top scores and game statistics.

**Features**:
- Shows all recorded high scores
- Displays user, session, score, and timestamp
- Styled display with Lipgloss
- Navigation back to menu

**Key Methods**:
- `NewLeaderboardModel()`: Fetch and display scores
- `Update()`: Handle navigation input
- `View()`: Render leaderboard table

### 7. **Score Service** (`internal/score.go`)

**Interface**: `ScoreService`

**Purpose**: Handles all score persistence and retrieval operations.

**Methods**:

```go
type ScoreService interface {
    GetHighScore(ctx context.Context) (Score, error)
    GetScores(ctx context.Context) ([]Score, error)
    SetCurrentScore(ctx context.Context, value int) error
    GetCurrentScore(ctx context.Context) (int, error)
}
```

**Implementation**: `ScoreServiceImpol`

**Responsibilities**:
- Save scores to SQLite database
- Retrieve high scores
- Track per-session scores
- Query by user and timestamp

### 8. **Session Manager** (`internal/session.go`)

**Interface**: `SessionManager`

**Purpose**: Manage user sessions using in-memory storage.

**Methods**:

```go
type SessionManager interface {
    CreateNewSession(value any) (string, error)
    DestroyCurrentSession() error
    GetCurrentSession() (string, error)
}
```

**Implementation**: `InMemeorySessiomManagerImpl`

**Features**:
- Generate random session IDs (20-character base64)
- Track current active session
- Session cleanup on exit

### 9. **Database Management** (`internal/db.go`)

**Purpose**: SQLite database initialization and schema management.

**Key Functions**:
- `CreateDB()`: Initialize database connection
- `execScoreTableCreation()`: Create scores table

**Schema**:

```sql
CREATE TABLE IF NOT EXISTS scores (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    user TEXT,
    session TEXT UNIQUE,
    value INTEGER,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
)
```

### 10. **Configuration & Initialization** (`internal/internal.go`)

**Purpose**: Global service initialization and configuration management.

**Functions**:
- `IntializeConfigs()`: Set up database, session manager, and score service
- `GetSessionManager()`: Retrieve session manager instance
- `GetScoreService()`: Retrieve score service instance

**Initialization Flow**:
1. Create SQLite database
2. Get system hostname
3. Initialize session manager
4. Initialize score service with user and database

### 11. **View Modes** (`tui/views/mode.go`)

**Purpose**: Define application view states and mode transitions.

**Mode Constants**:
- `ModeMenu`, `ModeGame`, `ModeLeaderboard`, `ModeGameOver`, `ModeGameCompleted`
- `ModeGame1-5`: Difficulty levels

**Helper Functions**:
- `NextLevelModeFromCurrent()`: Calculate next level based on current
- `SwitchModeCmd()`: Create mode switch command
- `ClearScreen()`: Clear terminal display

### 12. **Game Styles** (`tui/game/styles.go`)

**Purpose**: Define visual styling for game elements using Lipgloss.

**Style Definitions**:
- Board styling and colors
- Food and snake rendering
- Obstacle/pillar styling
- Game over screen styling

### 13. **Obstacles/Pillars** (`tui/game/pillars.go`, `tui/game/pillar_two.go`)

**Purpose**: Define static obstacle positions for each difficulty level.

**Data Structure**:
- Arrays of `Position` structs for each level
- Level 1-5 progressively more complex obstacle patterns
- Used in collision detection during gameplay

### 14. **Keyboard Input** (`utils/keys.go`)

**Purpose**: Centralized keyboard input mapping and validation.

**Key Definitions**:
- `KeyUp`, `KeyDown`, `KeyLeft`, `KeyRight`: Movement keys with multiple options
- `Enter`, `Esc`, `Space`: Action keys

**Helper Function**:
- `KeyMatchesInput()`: Check if input matches any defined key

---

## 🎮 Game Mechanics

### Core Loop

1. **Initialize**: Snake starts at center, food spawns randomly
2. **Input Processing**: Player input updates direction
3. **Movement**: Snake moves one position per tick
4. **Collision Detection**: Check for:
   - Food collision (grow snake, increase score)
   - Self collision (game over)
   - Wall collision (game over)
   - Obstacle collision (game over)
5. **Rendering**: Draw game board with snake, food, obstacles
6. **Score Update**: Persist score to database

### Difficulty Levels

| Level | Feature | Pillars |
|-------|---------|---------|
| Level 1 | Basic gameplay | 1 horizontal bar |
| Level 2 | Increased speed | 2 obstacles |
| Level 3 | More complex layout | Complex patterns |
| Level 4 | Advanced patterns | Dense obstacles |
| Level 5 | Expert mode | Maximum complexity |

### Scoring

- Each food eaten: +1 to score
- Score persisted to database upon completion
- High scores tracked across sessions
- User and session identifiers recorded

---

## 🗄️ Database Schema

### Scores Table

```sql
CREATE TABLE scores (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user TEXT,                          -- System hostname
    session TEXT UNIQUE,                -- Random session ID
    value INTEGER,                      -- Final score
    created_at DATETIME DEFAULT NOW()   -- Timestamp
)
```

**Data Flow**:
1. Session created on game start
2. Score updated during gameplay
3. Final score saved to database when game ends
4. Leaderboard queries all scores ordered by value

---

## 🛠️ Technologies & Dependencies

### Core Framework
- **Bubble Tea**: TUI framework (github.com/charmbracelet/bubbletea)
- **Lipgloss**: Terminal styling (github.com/charmbracelet/lipgloss)
- **Bubbles**: TUI components (github.com/charmbracelet/bubbles)
- **Charm Utils**: Utility helpers (github.com/Broderick-Westrope/charmutils)

### CLI & Utilities
- **Cobra**: CLI framework (github.com/spf13/cobra)

### Database
- **SQLite3**: Embedded database (github.com/mattn/go-sqlite3)

### Development
- **Go 1.24.2**: Programming language
- **Standard library**: time, context, database/sql, log, etc.

---

## 📝 Development Notes

### Code Organization

- **Separation of Concerns**: Game logic, UI, and persistence are separate
- **Interface-Based Design**: Services defined as interfaces for flexibility
- **Configuration Pattern**: Game configurations passed through config structs
- **Context Usage**: Database operations use context for cancellation/timeouts

### Styling & Theming

- Colors defined using Lipgloss styles
- ASCII art titles for visual appeal
- Responsive layout adapting to terminal size
- Custom spinner for loading states

### Error Handling

- Database errors logged and handled gracefully
- Context errors propagated properly
- Missing score/session handled with defaults
- Fatal errors exit with status code 1

---

## 🚀 Future Enhancement Ideas

1. Network multiplayer support
2. Custom difficulty settings
3. Achievement/badge system
4. Saved game state
5. Themes/color schemes
6. Sound effects (terminal bells)
7. Replay functionality
8. Advanced AI opponent
9. Mobile app version
10. Web version using WASM

---

## 📄 License

See [LICENSE](LICENSE) file for details.

---

## 👨‍💻 Development

### Building

```bash
go build -o super_snake
```

### Testing

```bash
go test ./...
```

### Running with Debug

```bash
go run main.go
```

### Checking Dependencies

```bash
go mod tidy
go mod verify
```

---

## 🤝 Contributing

Contributions are welcome! Please feel free to submit issues and pull requests.

---

**Enjoy the game! 🎮**
