package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}

	tasks := make(chan int)
	results := make(chan int)

	wg := &sync.WaitGroup{}

	// Горутина-генератор:
	// 1) Посылает числа из слайса в канал tasks
	// 2) После завершения закрывает канал tasks
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(tasks)

		for _, num := range numbers {
			tasks <- num
		}
	}()

	// Горутина-обработчик:
	// 1) Читает числа из канала tasks
	// 2) Умножает их на 2 и посылает в results
	// 3) После завершения закрывает канал results
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(results)
		for num := range tasks {
			results <- num * 2
		}
	}()

	// Вывод результатов
	wg.Add(1)
	go func() {
		defer wg.Done()
		for result := range results {
			fmt.Println(result)
		}
	}()
	wg.Wait()
}
