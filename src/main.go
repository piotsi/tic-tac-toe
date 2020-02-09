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

	// Check if mouse is inside a tile
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		fmt.Println("click")
		for i := 0; i < boardSize; i++ {
			for j := 0; j < boardSize; j++ {
				if rl.CheckCollisionPointRec(g.mousePos, g.board[i][j].position) {
					if g.board[i][j].state == "" {
						fmt.Println(g.board[i][j].index)
						g.board[i][j].state = g.turn
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

	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			rl.DrawRectangleLinesEx(g.board[i][j].position, shapeThickness, rl.Black)
			rl.DrawText(fmt.Sprintf("%s", g.board[i][j].state), int32(g.board[i][j].position.X)+tileSize/4, int32(g.board[i][j].position.Y), tileSize, rl.Black)
		}
	}

	//Draw mouse position
	rl.DrawText(fmt.Sprintf("x=%4.0f y=%4.0f turn=%s", rl.GetMousePosition().X, rl.GetMousePosition().Y, g.turn), 0, 0+screenSize, 20, rl.Black)

	rl.EndDrawing()
}
