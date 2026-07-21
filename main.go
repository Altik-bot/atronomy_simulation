package main

import "fmt"

func form(t1 int, t2 int, r int) float64 {
	const G int = 2
	var F float64 = float64(G * ((t1 + t2) / (r * r)))
	fmt.Println(F)
	return F
}
func main() {
	var mass1 int = 10
	var mass2 int = 20
	var radius int = 2
	form(mass1, mass2, radius)
}
