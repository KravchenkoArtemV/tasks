// Необходимо создать  горутину, которая бесконечно пишет в канал data.
// В main нужно остановить эту горутину через 1 секунду, используя done-канал и select.

package main

import (
	"fmt"
	"time"
)

func main() {
	data := make(chan int)
	done := make(chan struct{})

	// запускаем воркер в горутине
	go func() {
		for {
			select {
			case <-done:
				// если пришёл сигнал в канал done — выходим из горутины
				fmt.Println("остановка горутины")
				return
			case data <- 1:
				// иначе пишем в канал данные
				fmt.Println("отправлено значение")
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()

	// читаем данные из канала в main
	go func() {
		for val := range data {
			fmt.Println("main получил:", val)
		}
	}()

	// ждём 1 секунду
	time.Sleep(1 * time.Second)
	close(done) // отправляем сигнал завершения

	// ждём немного, чтобы увидеть вывод
	time.Sleep(300 * time.Millisecond)
	fmt.Println("main завершён")
}
