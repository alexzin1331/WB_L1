package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

// scanner и readInt для буферизованного ввода данных
var scanner = bufio.NewScanner(os.Stdin)

func readInt() int {
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	return num
}

func main() {
	scanner.Split(bufio.ScanWords)
	n := readInt()

	// Создаем канал для передачи данных
	dataChan := make(chan int)

	// Запускаем горутину-генератор данных
	go func() {
		for i := 1; ; i++ {
			dataChan <- i // Отправляем последовательные числа в канал
			time.Sleep(500 * time.Millisecond)
		}
	}()

	// Создаем таймер, который сработает через n секунд
	timeout := time.After(time.Duration(n) * time.Second)

	for {
		select {
		case data := <-dataChan: // Если пришли данные из канала, то выводим их
			fmt.Printf("Прочитано значение: %d\n", data)
		case <-timeout: // если закончился таймер
			fmt.Printf("Программа завершена, прошло %d секунд\n", n)
			return
		}
	}
}
