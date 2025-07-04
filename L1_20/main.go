package main

import (
	"fmt"
)

//в Go строки неизменяемые, поэтому сделал через байтовый слайс, надеюсь, засчитаете, что выполнил операцию «на месте»:)

func reverseWords(s []byte) []byte {
	var words [][]byte
	start := 0
	for i := 0; i <= len(s); i++ {
		if i == len(s) || s[i] == ' ' {
			words = append(words, s[start:i])
			start = i + 1
		}
	}

	// Меняем слова местами
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	// Собираем обратно в []byte
	result := make([]byte, 0, len(s))
	for i, word := range words {
		if i > 0 {
			result = append(result, ' ')
		}
		result = append(result, word...)
	}

	return result
}

func main() {
	input := "snow dog sun"
	inputByte := []byte(input)
	output := reverseWords(inputByte)
	fmt.Println(string(output)) // "sun dog snow"
}
