package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
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

	// Создаем контекст, который отменится по Ctrl+C
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Канал для системных сигналов
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	ch := make(chan int64)
	var wg sync.WaitGroup

	wg.Add(n)
	for i := 0; i < n; i++ {
		go workerProcess(i, ch, &wg)
	}

	//горутина, в которой высылаем данные в канал
	go func() {
		for {
			select {
			case <-ctx.Done():
				close(ch)
				return
			default:
				ch <- time.Now().Unix()
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	// Ожидаем сигнал завершения
	go func() {
		<-sigChan
		fmt.Println("Shutting down")
		cancel() // Отменяем контекст
	}()

	wg.Wait() // Ждем завершения всех воркеров
	fmt.Println("All workers finished")
}

// обработка сообщений из канала
func workerProcess(id int, ch <-chan int64, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range ch {
		fmt.Printf("Worker %d: %d\n", id, task)
	}
	fmt.Printf("Worker %d shutting down...\n", id)
}
