package main

import "fmt"

//функция для удаления элемента
func removeElement(slice []int, i int) []int {
	//проверка валидности индекса
	if i < 0 || i >= len(slice) {
		return slice
	}

	//копирование слайса (удалиться i-тый элемент)
	copy(slice[i:], slice[i+1:])
	
	return slice[:len(slice)-1]
}

func main() {
	slice := []int{10, 20, 30, 40, 50}
	fmt.Println("Исходный слайс:", slice) // [10 20 30 40 50]

	slice = removeElement(slice, 2)
	fmt.Println("После удаления:", slice) // [10 20 40 50]

	slice = removeElement(slice, 0)
	fmt.Println("После удаления:", slice) // [20 40 50]
}
