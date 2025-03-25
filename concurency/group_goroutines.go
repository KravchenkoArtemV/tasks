package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	ch := make(chan int)
	sum := 0
	var wg sync.WaitGroup

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			n := rand.Intn(100) + 1

			ch <- n
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for v := range ch {
		sum += v
	}

	fmt.Println("Сумма:", sum)
}
