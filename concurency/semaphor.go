package main

import (
	"fmt"
	"sync"
)

// Организуйте выполнение программы таким образом, чтобы на экран выводились два числа строго в порядке: 1 и 2.
// Воспользуйтесь для решения буферизованным каналом размером 1.

func main() {
	sem := make(chan struct{}, 1)
	wg := &sync.WaitGroup{}

	wg.Add(2)
	go func() {
		defer wg.Done()
		fmt.Println("1")
		sem <- struct{}{}
	}()

	go func() {
		defer wg.Done()
		<-sem
		fmt.Println("2")
	}()

	wg.Wait()
}
