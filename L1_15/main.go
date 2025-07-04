package main

import "fmt"

/*
если запись 1 &lt;&lt; 10 означает 1 << 10, то проблема кода в том, что программа создает срез justString, который ссылается на исходную большую строку v. Соответственно, Garbage Collector не будет очищать память от крупной переменной v до тех пор, пока будет использоваться переменная justString, хотя justString содержит лишь малую часть от переменной v, то есть неэффективно расходуется память.
*/

func createHugeString(size int) string {
	return string(make([]byte, size))
}

func someFunc() string {
	v := createHugeString(1 << 10)
	justString := make([]byte, 100)
	// Копируем только первые 100 байт, чтобы избежать утечки памяти
	copy(justString, v[:100]) // после этого Garbage Collector удалит v из памяти
	return string(justString)
}

func main() {
	justString := someFunc()
	fmt.Println(justString)
}
