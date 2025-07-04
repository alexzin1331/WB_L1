package main

import "fmt"

// функция для переворота строки
func reverseString(s string) string {
	runes := []rune(s)
	//проходимся по строке и меняем буквы местами
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func main() {
	//тесты
	examples := []string{
		"главрыба",
		"Hello, 世界",
		"👋🌍",
		"123абв",
	}

	for _, s := range examples {
		fmt.Printf("Оригинал: %q\n", s)
		fmt.Printf("Перевёрнуто: %q\n\n", reverseString(s))
	}
}
