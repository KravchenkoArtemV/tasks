package main

import (
	"context"
	"fmt"
	"time"
)

// Задача - сделать горутину-воркер, которая пишет данные в канал.
// Через 1 секунду остановить её с помощью context.WithCancel().

func main() {
	ctx, cancel := context.WithCancel(context.Background()) // создаём контекст с возможностью отмены
	data := make(chan int)

	// запускаем воркера
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("горутина остановлена по сигналу cancel()")
				return
			case data <- 42:
				fmt.Println("отправлено значение 42")
				time.Sleep(200 * time.Millisecond)
			}
		}
	}(ctx)

	// читаем данные в другой горутине
	go func() {
		for val := range data {
			fmt.Println("main получил:", val)
		}
	}()

	time.Sleep(1 * time.Second)
	cancel() // останавливаем через контекст

	// подождём чуть, чтобы увидеть вывод
	time.Sleep(300 * time.Millisecond)
	fmt.Println("main завершён")
}
