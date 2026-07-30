package main

import (
	"image/color"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth  = 900
	screenHeight = 900

	worldMinX = 0.0
	worldMaxX = 30.0
	worldMinY = 0.0
	worldMaxY = 30.0

	stepsPerFrame = 5
)

type Game struct{}

func worldToScreen(x, y float64) (float32, float32) {
	sx := float32((x - worldMinX) / (worldMaxX - worldMinX) * screenWidth)

	// Переворачиваем Y, чтобы вверх был положительным
	sy := float32(screenHeight - (y-worldMinY)/(worldMaxY-worldMinY)*screenHeight)

	return sx, sy
}

// Выполняет физику без вывода в консоль
func silentPhysicsStep() {
	old := os.Stdout
	devNull, _ := os.OpenFile(os.DevNull, os.O_WRONLY, 0)

	os.Stdout = devNull
	runConsole()
	os.Stdout = old

	devNull.Close()
}

// Границы мира
func enforceBoundaries() {

	if x1 <= worldMinX {
		x1 = worldMinX
		v_x1 = -v_x1
	}
	if x1 >= worldMaxX {
		x1 = worldMaxX
		v_x1 = -v_x1
	}
	if y1 <= worldMinY {
		y1 = worldMinY
		v_y1 = -v_y1
	}
	if y1 >= worldMaxY {
		y1 = worldMaxY
		v_y1 = -v_y1
	}

	if x2 <= worldMinX {
		x2 = worldMinX
		v_x2 = -v_x2
	}
	if x2 >= worldMaxX {
		x2 = worldMaxX
		v_x2 = -v_x2
	}
	if y2 <= worldMinY {
		y2 = worldMinY
		v_y2 = -v_y2
	}
	if y2 >= worldMaxY {
		y2 = worldMaxY
		v_y2 = -v_y2
	}
}

func (g *Game) Update() error {

	for i := 0; i < stepsPerFrame; i++ {
		silentPhysicsStep()
		enforceBoundaries()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	screen.Fill(color.RGBA{20, 20, 20, 255})

	xBody1, yBody1 := worldToScreen(x1, y1)
	xBody2, yBody2 := worldToScreen(x2, y2)

	// Тело 1
	vector.DrawFilledCircle(
		screen,
		xBody1,
		yBody1,
		14,
		color.RGBA{255, 70, 70, 255},
		false,
	)

	// Тело 2
	vector.DrawFilledCircle(
		screen,
		xBody2,
		yBody2,
		7,
		color.RGBA{70, 170, 255, 255},
		false,
	)
}

func (g *Game) Layout(_, _ int) (int, int) {
	return screenWidth, screenHeight
}

func main() {

	ebiten.SetWindowTitle("Симуляция двух тел")
	ebiten.SetWindowSize(screenWidth, screenHeight)

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
