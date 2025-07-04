package main

import (
	"fmt"
	"strings"
)

func uniqueSymbolsCheck(s string) bool {
	//переводим буквы в нижний регистр
	lowerString := strings.ToLower(s)

	//мапа для проверки, был ли символ в строке
	charMap := make(map[rune]struct{})

	//проходимся по строке и проверяем, заполняем мапу
	for _, symbol := range lowerString {
		if _, ok := charMap[symbol]; ok {
			return false
		}
		charMap[symbol] = struct{}{}
	}
	return true
}

func main() {
	//тесты
	slices := []string{
		"abcd",
		"abCdefAaf",
		"aabcd",
		"",
		"12345",
		"11234",
		"Test",
		"test",
	}
	for _, m := range slices {
		fmt.Printf("'%s': %t\n", m, uniqueSymbolsCheck(m))
	}
}
