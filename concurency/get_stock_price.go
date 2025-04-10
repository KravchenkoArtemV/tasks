package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

// getStockPrice имитирует получение цены с биржи (медленно)
func getStockPrice() int64 {
	time.Sleep(1 * time.Second) // задержка 1 секунда
	return rand.Int63n(1000)    // случайное значение от 0 до 999
}

// глобальные переменные для хранения последней цены и защиты доступа к ней
var result int64    // актуальная цена
var mu sync.RWMutex // мьютекс с поддержкой RLock для чтения

func main() {
	// быстрый HTTP-обработчик — возвращает кэшированное значение
	http.HandleFunc("/stocks/instant", func(resp http.ResponseWriter, req *http.Request) {
		mu.RLock()                      // безопасно читаем цену
		fmt.Fprintf(resp, "%d", result) // отправляем как строку
		mu.RUnlock()
	})

	// запускаем фоновую горутину, которая постоянно обновляет цену
	go func() {
		for {
			temp := getStockPrice() // получаем новую цену (медленно)
			mu.Lock()               // защищаем запись
			result = temp
			mu.Unlock()
		}
	}()

	// запускаем HTTP-сервер
	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}
