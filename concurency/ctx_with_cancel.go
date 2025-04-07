package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type task struct {
	value int
	delay time.Duration
}

func doWork(ctx context.Context, cancel context.CancelFunc, wg *sync.WaitGroup, ch chan task) {
	defer wg.Done()
	for {
		select {
		case v := <-ch:
			if v.value == 0 {
				cancel()
				fmt.Println("end")
				return
			}
			time.Sleep(v.delay)
		case <-ctx.Done():
			fmt.Println("end")
			return
		}
	}
}

func main() {
	var num int
	fmt.Scan(&num)

	ctx, cancel := context.WithCancel(context.Background())

	inputCh := make(chan task, num)

	for i := 0; i < num; i++ {
		var value, delay int
		fmt.Scan(&value, &delay)
		inputCh <- task{value, time.Duration(delay) * time.Millisecond}
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go doWork(ctx, cancel, &wg, inputCh)
	go doWork(ctx, cancel, &wg, inputCh)

	wg.Wait()
}
