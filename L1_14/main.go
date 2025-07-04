package main

import "fmt"

func main() {
	//примеры типов данных
	items := []interface{}{
		[]int{1, 2, 3},
		[]string{"a", "b"},
		[]bool{true, false},
		true,
		1,
		"efwef",
	}
	//пытаемся определить тип в цикле
	for _, item := range items {
		fmt.Println(detectType(item))
	}
}

func detectType(v interface{}) string {
	switch v.(type) {
	case []int:
		return "[]int"
	case []string:
		return "[]string"
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	default:
		return "unknown type"
	}
}
