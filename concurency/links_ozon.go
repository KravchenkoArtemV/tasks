package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	// список URL'ов для проверки
	urls := []string{
		"http://ozon.ru",
		"https://ozon.ru",
		"http://google.com",
		"http://somesite.com",
		"http://non-existent.domain.tld",
		"https://ya.ru",
		"http://ya.ru",
		"http://ёёёё",
	}

	// структурка для ответа
	type result struct {
		url string
		ok  bool
	}

	var wg sync.WaitGroup        // для ожидания завершения всех горутин
	results := make(chan result) // канал для получения результата от воркеров

	// для каждого URL запускаем воркера
	for _, url := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()

			// отправляем запрос
			resp, err := http.Get(u)
			if resp != nil {
				defer resp.Body.Close()
			}

			// отправляем результат в канал
			ok := err == nil && resp.StatusCode == 200
			results <- result{url: u, ok: ok}
		}(url)
	}

	// горутина-закрыватель канала: закрываем results, когда все воркеры завершат работу
	go func() {
		wg.Wait()
		close(results)
	}()

	// читаем и печатаем результаты в основном потоке
	for res := range results {
		if res.ok {
			fmt.Println(res.url, "- ok")
		} else {
			fmt.Println(res.url, "- not ok")
		}
	}
}
