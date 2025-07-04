package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go worker3(ctx)
	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(100 * time.Millisecond)
}

// Реализация с контекстом:
// 1) Создаем отменяемый контекст (context.WithCancel)
// 2) В горутине worker3 проверяем состояние контекста через select
// 3) При вызове cancel() контекст отменяется, срабатывает case <-ctx.Done(), и горутина завершается
func worker3(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stopped")
			return
		default:
			fmt.Println("working")
			time.Sleep(200 * time.Millisecond)
		}
	}
}
