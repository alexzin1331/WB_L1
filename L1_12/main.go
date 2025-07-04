package main

import "fmt"

// Создание множества (уникальных элементов):
// 1) Используем мапу, где ключи - элементы последовательности
// 2) Повторные элементы автоматически отбрасываются
// 3) Значение bool не важно, используется только факт наличия ключа
func main() {
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]bool)

	for _, item := range sequence {
		set[item] = true
	}

	fmt.Println(set)

}
