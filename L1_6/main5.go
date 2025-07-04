package main

import (
	"fmt"
	"time"
)

// 1) В горутине worker5 используется defer с recover для перехвата паники
// 2) После первой итерации вызывается panic(), что приводит к остановке горутины
// 3) recover перехватывает панику и позволяет выполнить завершающие действия
// 4) Это не рекомендуется для обычной остановки горутин, только для обработки критических ошибок
func worker5() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()
	for {
		fmt.Println("Working...")
		time.Sleep(500 * time.Millisecond)
		panic("stop requested") // Принудительная остановка
	}
}

func main() {
	go worker5()
	time.Sleep(1 * time.Second)
}
