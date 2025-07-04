package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go worker4()
	time.Sleep(1 * time.Second)
}

// Реализация с runtime.Goexit:
// 1) Горутина worker4 работает в бесконечном цикле
// 2) После первой итерации вызывается runtime.Goexit(), который завершает текущую горутину
// 3) Это нестандартный способ остановки, обычно используется для аварийного завершения
func worker4() {
	for {
		fmt.Println("Working...")
		time.Sleep(500 * time.Millisecond)
		runtime.Goexit()
	}
}
