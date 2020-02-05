package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	scale      = 2
	screenSize = 192 * scale
)

const (
	boardSize      = 3
	shapeThickness = 20
)

// Tile bla
type Tile struct {
	position rl.Rectangle
	state    int32
}

// Game bla
type Game struct {
	gameOver bool
	mousePos rl.Vector2
	board    [boardSize ^ 2]Tile
}

func main() {
	rl.InitWindow(screenSize, screenSize, "TIC-TAC-TOE")
	rl.SetTargetFPS(60)

	game := NewGame()

	for !rl.WindowShouldClose() {
		game.Update()
		game.Draw()
		// rl.DrawFPS(0, 0)
	}

	rl.CloseWindow()
}

// NewGame bla
func NewGame() (g Game) {
	g.Init()
	return
}

// Init bla
func (g *Game) Init() {
	g.gameOver = false

	for i := 0; i < boardSize^2; i++ {
		g.board[i].position.Width = screenSize / 3
		g.board[i].position.Height = screenSize / 3
		// wymyslic jak generowac pozycje
		g.board[i].position.X = 0
		g.board[i].position.Y = 0
		g.board[i].state = 0
	}
}

// Update bla
func (g *Game) Update() {
	g.gameOver = true

	g.mousePos = rl.GetMousePosition()

	// Check if mouse is inside a tile
	for i := 0; i < boardSize^2; i++ {
		if rl.CheckCollisionPointRec(g.mousePos, g.board[i].position) {
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
				println("test")
			}
		}
	}
}

// Draw bla
func (g *Game) Draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.White)

	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {

		}
	}

	// Draw mouse position
	// rl.DrawText(fmt.Sprintf("x=%.0f y=%.0f", rl.GetMousePosition().X, rl.GetMousePosition().Y), 0, 20, 20, rl.Black)

	rl.EndDrawing()
}
