package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

// Distance вычисляет расстояние между текущей точкой и другой точкой
func (p Point) Distance(other Point) float64 {
	dx := p.x - other.x
	dy := p.y - other.y
	//вычисляем по теореме пифагора
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	//создаем тестовые точки
	pointA := NewPoint(1.5, 2.7)
	pointB := NewPoint(4.2, 6.1)

	//вычисляем дистанцию
	distance := pointA.Distance(pointB)

	fmt.Printf("Расстояние между точкой A(%.1f, %.1f) и точкой B(%.1f, %.1f) = %.2f\n",
		pointA.x, pointA.y, pointB.x, pointB.y, distance)
}
