package main

import (
	"fmt"
	"sync"
)

// считает факториал числа n
func factorial(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func main() {
	numbers := []int{11, 4, 6, 10}

	// выделяем память под срез результатов той же длины, что и numbers
	// это позволяет каждой горутине писать в свой индекс без гонки данных
	results := make([]int, len(numbers))

	var wg sync.WaitGroup

	// запускаем горутину для каждого числа
	for i, num := range numbers {
		wg.Add(1) // увеличиваем счётчик

		// передаём i и num как аргументы
		go func(idx, n int) {
			defer wg.Done() // уменьшаем счётчик

			// каждая горутина пишет в свой индекс среза
			results[idx] = factorial(n)
		}(i, num)
	}

	// ждём, пока все горутины закончат вычисления
	wg.Wait()

	for i, num := range numbers {
		fmt.Printf("Факториал %d = %d\n", num, results[i])
	}
}
