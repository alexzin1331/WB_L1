package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := int64(5)
	i1 := int64(2)
	i2 := int64(1)
	fmt.Printf("Исходное число: %d (%s)\n", n, strconv.FormatInt(n, 2))
	// установка i-го бита в 1
	//0100 -> 0101: 0100|(i<<5)
	n = n | (1 << (i1 - 1))
	fmt.Printf("После установки %d-го бита в 1: %d (%s)\n", i1, n, strconv.FormatInt(n, 2))
	//установка i-го бита в 0
	//0101 -> 0100: 0101& ^(i<<5)
	n = n &^ (1 << (i2 - 1))
	fmt.Printf("После установки %d-го бита в 0: %d (%s)\n", i2, n, strconv.FormatInt(n, 2))
}
