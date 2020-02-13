package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	scale          = 1
	tileSize       = 100 * scale
	screenSize     = tileSize*boardSize + padding*(boardSize+1)
	boardSize      = 3
	shapeThickness = 5
	menuHeight     = 20
	padding        = 10
)

// Tile parameters
type Tile struct {
	position rl.Rectangle
	state    string
	pressed  bool
	index    int
}

// Game parameters
type Game struct {
	gameOver bool
	mousePos rl.Vector2
	board    [boardSize][boardSize]Tile
	turn     string
	winner   string
}

func main() {
	rl.InitWindow(screenSize, screenSize+menuHeight, "TIC-TAC-TOE")
	rl.SetTargetFPS(100)

	game := NewGame()

	for !rl.WindowShouldClose() {
		game.Update()
		game.Draw()
	}

	rl.CloseWindow()
}

// NewGame initialize
func NewGame() (g Game) {
	g.Init()
	return
}

// Init method
func (g *Game) Init() {
	g.gameOver = false
	g.turn = "o"

	// Prepare the board specific parameters
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			g.board[i][j].position.Width = tileSize
			g.board[i][j].position.Height = tileSize
			g.board[i][j].position.X = tileSize*float32(i) + float32(padding*(i+1))
			g.board[i][j].position.Y = tileSize*float32(j) + float32(padding*(j+1))
			g.board[i][j].index = i + boardSize*j
			g.board[i][j].state = ""
		}
	}
}

// Update method
func (g *Game) Update() {
	g.gameOver = false
	g.mousePos = rl.GetMousePosition()

	// Check on where mouse was pressed
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) && g.winner == "" {
		for i := 0; i < boardSize; i++ {
			for j := 0; j < boardSize; j++ {
				if rl.CheckCollisionPointRec(g.mousePos, g.board[i][j].position) {
					if g.board[i][j].state == "" {
						g.board[i][j].state = g.turn
						CheckWin(i, j, g)
						// Change turn
						if g.turn == "x" {
							g.turn = "o"
						} else {
							g.turn = "x"
						}
					}
				}
			}
		}
	}
}

// Draw method
func (g *Game) Draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.White)

	// Draw the board
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			rl.DrawRectangleLinesEx(g.board[i][j].position, shapeThickness, rl.Black)
			rl.DrawText(fmt.Sprintf("%s", g.board[i][j].state), int32(g.board[i][j].position.X)+tileSize/4, int32(g.board[i][j].position.Y), tileSize, rl.Black)
		}
	}

	// Draw menu items
	rl.DrawText(fmt.Sprintf("Turn=%s", g.turn), 0+10, 0+screenSize, 20, rl.Black)
	rl.DrawText(fmt.Sprintf("Winner=%s", g.winner), screenSize/2+10, 0+screenSize, 20, rl.Black)

	// Draw game won screen
	if g.winner != "" {
		rl.DrawRectangle(0, 0, screenSize, screenSize+menuHeight, rl.Black)
		rl.DrawText(fmt.Sprintf("Game over \n  '%s' won!", g.winner), 40, 120, 50, rl.White)
		// rl.EndDrawing()
	}

	rl.EndDrawing()
}

// CheckWin checks if there is a winner
func CheckWin(i int, j int, g *Game) {
	var winner bool = true
	// Check row, column and diagonals
	for x := 0; x < boardSize; x++ {
		if g.board[i][x].state != g.turn {
			winner = false
			break
		}
	}
	if winner {
		g.winner = g.turn
		g.gameOver = true
	}
	winner = true
	for x := 0; x < boardSize; x++ {
		if g.board[x][j].state != g.turn {
			winner = false
			break
		}
	}
	if winner {
		g.winner = g.turn
		g.gameOver = true
	}
	winner = true
	for x := 0; x < boardSize; x++ {
		if g.board[x][x].state != g.turn {
			winner = false
			break
		}
	}
	if winner {
		g.winner = g.turn
		g.gameOver = true
	}
	winner = true
	for x := 0; x < boardSize; x++ {
		if g.board[x][boardSize-x-1].state != g.turn {
			winner = false
			break
		}
	}
	if winner {
		g.winner = g.turn
		g.gameOver = true
	}
}
