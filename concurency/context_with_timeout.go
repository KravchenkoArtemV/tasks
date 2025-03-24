package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// создаём контекст с таймаутом 1 секунда
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel() // обязательно освобождаем ресурсы

	result := make(chan string)

	go func() {
		// симулируем долгую работу — 3 секунды
		time.Sleep(3 * time.Second)
		result <- "успешно завершено"
	}()

	select {
	case res := <-result:
		fmt.Println("Результат:", res)
	case <-ctx.Done():
		// сработает через 1 секунду
		fmt.Println("Операция отменена:")
	}
}
