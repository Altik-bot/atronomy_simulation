package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// Параметры визуализации
const (
	gridWidth     = 60 // ширина поля в символах
	gridHeight    = 24 // высота поля в символах
	stepsTotal    = 1000
	stepsPerFrame = 5 // ускорил анимацию, просчитывая больше шагов физики за 1 кадр
	frameDelay    = 40 * time.Millisecond
)

// Границы мира (только первая четверть от 0.0)
const (
	worldMinX = 0.0
	worldMaxX = 14.0
	worldMinY = 0.0
	worldMaxY = 14.0
)

func worldToGrid(x, y float64) (int, int) {
	col := int((x - worldMinX) / (worldMaxX - worldMinX) * float64(gridWidth-1))
	row := int((y - worldMinY) / (worldMaxY - worldMinY) * float64(gridHeight-1))
	return col, row
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

// silentPhysicsStep вызывает готовую логику из physics.go,
// "затыкая" её вывод.
func silentPhysicsStep() {
	original := os.Stdout
	devNull, err := os.OpenFile(os.DevNull, os.O_WRONLY, 0)
	if err == nil {
		os.Stdout = devNull
		defer func() {
			os.Stdout = original
			devNull.Close()
		}()
	}
	// Вызов твоей функции из physics.go (в предыдущих примерах она называлась stepPhysics или runConsole)
	runConsole()
}

// enforceBoundaries удерживает тела строго в первой четверти (x > 0, y > 0)
// и в пределах видимости (worldMax).
func enforceBoundaries() {
	// Границы для тела 1
	if x1 <= worldMinX {
		x1 = worldMinX
		v_x1 = -v_x1 // упругий отскок, меняем направление скорости
	}
	if y1 <= worldMinY {
		y1 = worldMinY
		v_y1 = -v_y1
	}
	if x1 >= worldMaxX {
		x1 = worldMaxX
		v_x1 = -v_x1
	}
	if y1 >= worldMaxY {
		y1 = worldMaxY
		v_y1 = -v_y1
	}

	// Границы для тела 2
	if x2 <= worldMinX {
		x2 = worldMinX
		v_x2 = -v_x2
	}
	if y2 <= worldMinY {
		y2 = worldMinY
		v_y2 = -v_y2
	}
	if x2 >= worldMaxX {
		x2 = worldMaxX
		v_x2 = -v_x2
	}
	if y2 >= worldMaxY {
		y2 = worldMaxY
		v_y2 = -v_y2
	}
}

func drawFrame(step int, t float64) {
	grid := make([][]rune, gridHeight)
	for i := range grid {
		grid[i] = make([]rune, gridWidth)
		for j := range grid[i] {
			grid[i][j] = ' ' // Заменил точки на пробелы для более чистого отображения
		}
	}

	// Отрисовка рамок (наши физические границы)
	for j := 0; j < gridWidth; j++ {
		grid[0][j] = '-'
		grid[gridHeight-1][j] = '-'
	}
	for i := 0; i < gridHeight; i++ {
		grid[i][0] = '|'
		grid[i][gridWidth-1] = '|'
	}

	// Перевод координат в сетку
	col1, row1 := worldToGrid(x1, y1)
	col2, row2 := worldToGrid(x2, y2)

	// Отрисовка Тела 1
	if row1 >= 0 && row1 < gridHeight && col1 >= 0 && col1 < gridWidth {
		grid[row1][col1] = 'O' // Кружок для 1 тела
	}

	// Отрисовка Тела 2 с проверкой на наложение
	if row2 >= 0 && row2 < gridHeight && col2 >= 0 && col2 < gridWidth {
		if grid[row2][col2] == 'O' {
			grid[row2][col2] = 'X' // Тела столкнулись
		} else {
			grid[row2][col2] = 'o' // Маленький кружок для 2 тела
		}
	}

	// Вывод на экран
	clearScreen()
	fmt.Printf("Шаг: %d   Время: %.2f c \n", step, t)
	fmt.Println(strings.Repeat("=", gridWidth))
	for _, row := range grid {
		fmt.Println(string(row))
	}
	fmt.Println(strings.Repeat("=", gridWidth))
	fmt.Printf("Тело 1: x=%.2f  y=%.2f | v_x=%.2f v_y=%.2f\n", x1, y1, v_x1, v_y1)
	fmt.Printf("Тело 2: x=%.2f  y=%.2f | v_x=%.2f v_y=%.2f\n", x2, y2, v_x2, v_y2)
	fmt.Printf("Расстояние: %.3f\n", distance)
}

func main() {
	for step := 1; step <= stepsTotal; step++ {
		for i := 0; i < stepsPerFrame; i++ {
			silentPhysicsStep() // 1. Считаем физику
			enforceBoundaries() // 2. Проверяем и корректируем границы (отскок)
		}

		drawFrame(step*stepsPerFrame, float64(step*stepsPerFrame)*dt)
		time.Sleep(frameDelay)
	}
}
