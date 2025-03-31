package main

import (
	"context"
	"fmt"
	"time"
)

//	Один общий небуферизованный канал.
//	Каждая горутина отправляет в него ch <- struct{}{}.
//	main() просто дважды получает из канала, чтобы дождаться обеих.

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan struct{})

	go func() {
		time.Sleep(time.Millisecond * 100)
		fmt.Println("done1")
		cancel()
		ch <- struct{}{}
	}()

	go func() {
		select {
		case <-time.After(time.Second * 100):
		case <-ctx.Done():
		}
		ch <- struct{}{}
		fmt.Println("done2")
	}()

	<-ch
	<-ch
}
