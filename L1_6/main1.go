package main

import (
	"fmt"
	"time"
)

func worker(stop *bool) {
	for !*stop {
		fmt.Println("Working")
		time.Sleep(500 * time.Millisecond)
	}
}

// реализация с каналом
// 1) создаем флаг
// 2) передаем его указатель в горутину
// 3) в бесконечном цикле проверям его, если он поменялся, то останавливаем горутину
func main() {
	stop := false
	go worker(&stop)
	time.Sleep(2 * time.Second)
	stop = true
}
