package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	value int
}

func main() {
	var wg sync.WaitGroup
	c := Counter{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.mu.Lock()
			defer c.mu.Unlock()
			c.value++
		}()
	}

	wg.Wait()
	c.mu.Lock()
	fmt.Println("Итоговое значение счётчика (Mutex):", c.value)
	c.mu.Unlock()
}
