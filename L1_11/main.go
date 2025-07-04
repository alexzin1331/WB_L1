package main

import "fmt"

// Intersection находит пересечение двух слайсов любого comparable типа
// 1) Создаем мапу для элементов первого слайса
// 2) Проходим по второму слайсу и проверяем наличие в мапе
// 3) При нахождении добавляем в результат и удаляем из мапы (чтобы избежать дубликатов)
func Intersection[T comparable](a, b []T) []T {
	mp := make(map[T]struct{})
	for _, x := range a {
		mp[x] = struct{}{}
	}

	var result []T
	for _, val := range b {
		if _, ok := mp[val]; ok {
			result = append(result, val)
			delete(mp, val)
		}
	}
	return result
}

func main() {
	fmt.Println("result 1: ", Intersection([]int{1, 2, 3, 4, 7}, []int{2, 3, 3, 5}))
	fmt.Println("result 2: ", Intersection([]string{"a", "b", "cd"}, []string{"s", "c", "a", "cd"}))
}
