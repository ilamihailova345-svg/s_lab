package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	sites := []string{"google.com", "yandex.ru", "github.com", "stackoverflow.com"}

	var wg sync.WaitGroup

	// для каждого сайта запускаем отдельную горутину
	for _, site := range sites {
		wg.Add(1)

		// site передаём как аргумент, чтобы каждая горутина видела свой сайт
		go func(s string) {
			defer wg.Done()

			fmt.Printf("Начинаю проверку %s\n", s)

			// имитация разной скорости ответа: от 1 до 3 секунд
			// rand.Intn(3) даёт 0, 1, 2. После +1 будет 1, 2, 3.
			duration := time.Duration(rand.Intn(3)+1) * time.Second
			time.Sleep(duration)

			fmt.Printf("Проверка %s завершена\n", s)
		}(site)
	}

	// ждём завершения всех проверок
	wg.Wait()
	fmt.Println("Проверка всех сайтов завершена")
}