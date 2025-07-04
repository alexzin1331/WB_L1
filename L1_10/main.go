package main

import (
	"fmt"
	"math"
)

func main() {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	results := make(map[int][]float64)
	// Группировка температур по диапазонам:
	// 1) Для каждой температуры вычисляем ключ (кратный 10)
	// 2) Особый случай для отрицательных температур (-25.4 → -20)
	// 3) Добавляем температуру в соответствующую группу
	for _, temp := range temps {
		key := int(math.Floor(temp/10)) * 10
		if key < 0 {
			key += 10
		}
		results[key] = append(results[key], temp)
	}
	for key, val := range results {
		fmt.Printf("%d: %v\n", key, val)
	}
}
