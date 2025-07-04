package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

var scanner = bufio.NewScanner(os.Stdin)

func readInt() int {
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	return num
}

func main() {
	scanner.Split(bufio.ScanWords)
	//вводим количество воркеров
	n := readInt()
	_, cancel := context.WithCancel(context.Background())
	defer cancel()
	wg := &sync.WaitGroup{}
	//канал для вывода результата
	ch := make(chan int64)
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			workerProcess(i, ch, wg)
		}()
	}

	go func() {
		defer close(ch)
		for {
			//посылаем сообщение в канал
			ch <- time.Now().Unix()
			time.Sleep(500 * time.Millisecond)
		}
	}()
	wg.Wait()
}

// обрабатываем сообщения, посланные в канал
func workerProcess(id int, ch <-chan int64, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range ch {
		fmt.Printf("Worker %d: %d\n", id, task)
	}
}
