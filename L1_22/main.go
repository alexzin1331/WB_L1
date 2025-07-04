package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Создаем очень большие числа
	a := &big.Int{}
	b := &big.Int{}
	a.SetString("1000000000000000000000", 10) // 10^21
	b.SetString("500000000000000000000", 10)  // 5*10^20
	fmt.Printf("a = %s, b = %s\n", a.String(), b.String())

	// Результаты операций
	result := &big.Int{}

	//сложение
	result.Add(a, b)
	fmt.Printf("Сумма: %s + %s = %s\n", a.String(), b.String(), result.String())

	//вычитание
	result.Sub(a, b)
	fmt.Printf("Разность: %s - %s = %s\n", a.String(), b.String(), result.String())

	//умножение
	result.Mul(a, b)
	fmt.Printf("Произведение: %s * %s = %s\n", a.String(), b.String(), result.String())

	//деление
	result.Div(a, b)
	fmt.Printf("Частное: %s / %s = %s\n", a.String(), b.String(), result.String())

	//остаток
	result.Mod(a, b)
	fmt.Printf("Остаток: %s %% %s = %s\n", a.String(), b.String(), result.String())
}
