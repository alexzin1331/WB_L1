package main

import (
	"context"
	"fmt"
	"time"
)

// остановка с помощью контекста с таймаутом
func sleep1(duration time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()
	<-ctx.Done()
}

// остановка с помощью таймера
func sleep2(duration time.Duration) {
	timer := time.NewTimer(duration)
	<-timer.C
}

func main() {
	fmt.Println("Stop for 5 sec")
	sleep1(5 * time.Second)
	fmt.Println("duration is expired with ctx")
	sleep2(5 * time.Second)
	fmt.Println("duration is expired with timer")
}
