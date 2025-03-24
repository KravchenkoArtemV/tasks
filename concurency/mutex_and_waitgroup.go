package main

import (
	"fmt"
	"sync"
)

// 5 горутин увеличивают общий счётчик. Нужно защитить его от гонки.

func main() {
	var mu sync.Mutex
	counter := 0
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				mu.Lock()   // захватываем доступ
				counter++   // безопасно увеличиваем
				mu.Unlock() // освобождаем доступ
			}
		}()
	}

	wg.Wait()
	fmt.Println("Итоговое значение:", counter) // ожидается 5000
}
