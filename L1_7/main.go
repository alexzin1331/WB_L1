package main

import (
	"fmt"
	"sync"
)

type Map struct {
	mu sync.Mutex
	m  map[int]int
}

func main() {
	mp := &Map{
		mu: sync.Mutex{},
		m:  make(map[int]int),
	}
	wg := &sync.WaitGroup{}
	//намеренно создаем конкурентную запись, перезаписывая одни и те же ячейки в мапе (чтобы показать, что mu.Lock() работает)
	// 1) Используем sync.Mutex для защиты доступа к мапе
	// 2) 100 горутин пытаются записать в 20 ячеек мапы
	// 3) Без мьютекса была бы data race, но с ним - нет
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mp.mu.Lock()
			mp.m[i%20] = i * i
			mp.mu.Unlock()
		}(i)
	}

	wg.Wait()

	for i := range mp.m {
		fmt.Println(mp.m[i])
	}
}

//при запуске go run -race гонки нет
