package main

import (
	"fmt"
	"sync"
)

// задача на отправку из одной горутины данных и чтения в другой

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(2)
	go func() {
		defer wg.Done()
		ch <- 3
		ch <- 7
		ch <- 11
		close(ch) // закрывает канал отправитель

	}()

	go func() {
		defer wg.Done()
		for i := range ch { // проходим
			fmt.Println(i)
		}
	}()

	wg.Wait()

}
