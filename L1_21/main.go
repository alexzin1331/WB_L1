package main

import (
	"fmt"
	"math"
)

// Новый интерфейс
type Shape interface {
	Area() float64
}

// Адаптируемые структуры (старая система)

// Старый квадрат
type Square struct {
	side float64
}

// функция для получения площади для квадрата
func (s *Square) GetSquareArea() float64 {
	return s.side * s.side
}

// Старый круг
type Circle struct {
	radius float64
}

// функция для получения площади для круга
func (c *Circle) GetCircleArea() float64 {
	return math.Pi * c.radius * c.radius
}

// Адаптеры для нового интерфейса

// Адаптер для квадрата
type SquareAdapter struct {
	square *Square
}

func (a *SquareAdapter) Area() float64 {
	return a.square.GetSquareArea()
}

// Адаптер для круга
type CircleAdapter struct {
	circle *Circle
}

func (a *CircleAdapter) Area() float64 {
	return a.circle.GetCircleArea()
}

func printArea(shape Shape) {
	fmt.Printf("Площадь фигуры: %.2f\n", shape.Area())
}

func main() {
	// Создаем фигуры старого формата
	square := &Square{side: 5}
	circle := &Circle{radius: 3}

	// Создаем адаптеры
	squareAdapter := &SquareAdapter{square: square}
	circleAdapter := &CircleAdapter{circle: circle}

	// Используем адаптеры в новом коде
	printArea(squareAdapter)
	printArea(circleAdapter)
}
