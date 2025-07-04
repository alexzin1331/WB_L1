package main

import "fmt"

func main() {
	a, b := 2, 7
	//1) способ (скорее всего он не считается, но дополнительную память не он использует:) )
	a, b = b, a
	fmt.Println(a, b)
	a, b = b, a
	//2) способ
	a = a + b
	b = a - b
	a = a - b
	fmt.Println(a, b)
	a, b = b, a
	//3) способ
	a = a ^ b
	b = a ^ b
	a = a ^ b
	fmt.Println(a, b)
}
