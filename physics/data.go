package main

var (
	dx          float64
	dy          float64
	ForceFactor float64
	distance    float64
	r           float64
	Fy          float64
	Fx          float64
)

// Храктеристики тела 1
var (
	m1   float64 = 10.0 // масса
	x1   float64 = 2.0  // позиция по оси х
	y1   float64 = 2.0  // позиция по оси у
	v_x1 float64        // скорость по оси х
	v_y1 float64        // скорость по оси у
	a_x1 float64
	a_y1 float64
)

// Характеристики для тела 2
var (
	m2   float64 = 5.0
	x2   float64 = 5.0
	y2   float64 = 6.0
	v_x2 float64
	v_y2 float64
	a_x2 float64
	a_y2 float64
)

// константы
const (
	G   float64 = 2
	dt  float64 = 0.01
	eps float64 = 0.001
)
