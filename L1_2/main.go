package main

import (
	"fmt"
	"sync"
)

func main() {
	m := []int{2, 4, 6, 8, 10}
	wg := &sync.WaitGroup{}
	//проходимся по слайсу и в каждой горутине вычисляем значение
	for _, v := range m {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			fmt.Println(v * v)
		}(v)
	}
	wg.Wait()
}
