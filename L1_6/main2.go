package main

import (
	"fmt"
	"time"
)

func main() {
	stop := make(chan struct{})
	go worker2(stop)
	time.Sleep(2 * time.Second)
	close(stop) // Посылаем сигнал остановки
	time.Sleep(100 * time.Millisecond)
}

// Реализация с каналом для остановки:
// 1) Создаем канал stop типа struct{} (так как он не занимает память)
// 2) В горутине worker2 в бесконечном цикле проверяем канал через select
// 3) При закрытии канала (close(stop)) срабатывает case <-stop, и горутина завершается
func worker2(stop chan struct{}) {
	for {
		select {
		case <-stop:
			fmt.Println("stopped")
			return
		default:
			fmt.Println("working")
			time.Sleep(400 * time.Millisecond)
		}
	}
}
