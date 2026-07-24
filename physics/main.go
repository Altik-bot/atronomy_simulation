package main

import (
	"fmt"
	"math"
)

func main() {
	Vector := func() float64 {
		dx = x2 - x1
		dy = y2 - y1
		distance = math.Sqrt(dx*dx+dy*dy) + eps
		fmt.Printf("Вектор смещения: (%.2f, %.2f)\n", dx, dy)
		fmt.Printf("Расстояние (модуль): %.4f\n\n", distance)
		return distance
	}

	forcefactor := func() float64 {
		ForceFactor = G * ((m1 * m2) / (math.Pow(distance, 3)))
		fmt.Printf("Сила: %f \n", ForceFactor)
		return ForceFactor
	}

	ForceXY := func() (float64, float64) {
		Fx = ForceFactor * dx
		Fy = ForceFactor * dy
		fmt.Printf("Fx, Fy: %f, %f \n", Fx, Fy)
		return Fx, Fy
	}

	acceleration := func() (float64, float64, float64, float64) {
		a_x1 = Fx / m1
		a_x2 = -Fx / m2
		a_y1 = Fy / m1
		a_y2 = -Fy / m2
		fmt.Printf("ACC: %f, %f, %f, %f \n", a_x1, a_x2, a_y1, a_y2)
		return a_x1, a_x2, a_y1, a_y2
	}

	speed := func() (float64, float64, float64, float64) {
		v_x1 = v_x1 + a_x1*dt
		v_y1 = v_y1 + a_y1*dt
		v_x2 = v_x2 + a_x2*dt
		v_y2 = v_y2 + a_y2*dt
		fmt.Printf("Speed: %f, %f. %f, %f \n", v_x1, v_x2, v_y1, v_y2)
		return v_x1, v_x2, v_y1, v_y2
	}

	position := func() (float64, float64, float64, float64) {
		x1 = x1 + v_x1*dt
		y1 = y1 + v_y1*dt

		x2 = x2 + v_x2*dt
		y2 = y2 + v_y2*dt
		fmt.Print("Position:\n")
		fmt.Printf("1st body: %f, %f \n", x1, y1)
		fmt.Printf("2nd body: %f, %f \n", x2, y2)
		return x1, x2, y1, y2
	}
	Vector()
	forcefactor()
	ForceXY()
	acceleration()
	speed()
	position()
}
