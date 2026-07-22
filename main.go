package main

import (
	"fmt"
	"math"
)

func main() {
	const g int = 2
	var m1 int = 10
	var m2 int = 20
	var radius int = 2
	var m1_x float64 = 1
	var m1_y float64 = 1
	var m2_x float64 = 7
	var m2_y float64 = 6
	form := func() float64 {
		var F float64 = float64(g * ((m1 * m2) / (radius * radius)))
		fmt.Println(F)
		return F
	}
	dist := func() float64 {
		var distance float64 = math.Sqrt((math.Pow(m2_x-m1_x, 2)) + (math.Pow(m2_y-m1_y, 2)))
		fmt.Println(distance)
		return distance
	}
	form()
	dist()
}
