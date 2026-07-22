package main

import (
	"fmt"
	"math"
)

func form(t1 int, t2 int, r int) float64 {
	const G int = 2
	var F float64 = float64(G * ((t1 * t2) / (r * r)))
	fmt.Println(F)
	return F
}

func dist(x1, y1, x2, y2 float64) float64 {
	var distance float64 = math.Sqrt((math.Pow(x2-x1, 2)) + (math.Pow(y2-y1, 2)))
	fmt.Println(distance)
	return distance
}

func main() {
	var m1 int = 10
	var m2 int = 20
	var radius int = 2
	var m1_x float64 = 1
	var m1_y float64 = 1
	var m2_x float64 = 7
	var m2_y float64 = 6
	form(m1, m2, radius)
	dist(m1_x, m1_y, m2_x, m2_y)
}
