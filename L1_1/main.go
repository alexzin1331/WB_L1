package main

import "fmt"

type Human struct {
	Age  int
	Name string
	Sex  string
}

// создаем структуру с вложенной Human
type Action struct {
	Field string
	Human
}

func main() {
	a := Action{
		Field: "Some_field",
		Human: Human{
			Age:  27,
			Name: "McLovin",
			Sex:  "M",
		},
	}
	//тестируем поля
	fmt.Println("Age:", a.Age, "\nSex: ", a.Sex)
	fmt.Println("name: ", a.Name, "\nField: ", a.Field)
}
