package main

import (
	"fmt"
	"sync"
	"time"
)

func printNumbers(id int, n int) {
	for i := 1; i <= n; i++ {
		fmt.Printf("Worker %d: number %d\n", id, i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			printNumbers(id, 5)
		}(i)
	}

	// было: time.Sleep(1 * time.Second) ненадёжно, не знаем, когда закончат
	// стало: wg.Wait() гарантированно ждёт завершения всех горутин
	wg.Wait()
	fmt.Println("Все воркеры завершили работу")
}