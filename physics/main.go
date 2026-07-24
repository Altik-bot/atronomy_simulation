package main

import (
	"fmt"
	"math"
)

func main() {
	vector := func() float64{
		r = (x2-x2 y2-y1)
	}
	dist := func() float64 {
		distance = math.Sqrt((math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))) + eps
		fmt.Printf("Dist: %f \n", distance)
		return distance
	}


	force := func() float64 {
		F = G * ((m1 * m2) / (distance * distance))
		fmt.Printf("Сила: %f \n", F)
		return F
	}


	acceleration := func() (float64, float64) {
		a1 = F / m1
		a2 = -F / m2
		fmt.Printf("Acc: %f %f", a1, a2)
		return a1, a2
	}
	dist()
	force()
	acceleration()
}
