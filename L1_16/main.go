package main

import "fmt"

func quickSort(m []int) {
	if len(m) <= 1 {
		return // Базовый случай рекурсии
	}
	// Опорный элемент (последний)
	pivot := m[len(m)-1]
	i := 0
	// Разделение элементов относительно опорного
	for j := 0; j < len(m)-1; j++ {
		if m[j] < pivot {
			m[i], m[j] = m[j], m[i]
			i++
		}
	}
	// Помещаем опорный элемент на правильную позицию
	m[i], m[len(m)-1] = m[len(m)-1], m[i]
	quickSort(m[:i])
	quickSort(m[i+1:])
}

func main() {
	arr := []int{3, 11, 4, 6, 13, 6, 5, 3, 56, -1, 2, 9, -2, 4, -13, -5}
	fmt.Println("Исходный массив:", arr)
	quickSort(arr)
	fmt.Println("Отсортированный массив:", arr)
}
