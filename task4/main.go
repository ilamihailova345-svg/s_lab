package main

import (
	"fmt"
	"sync"
	"time"
)

type Order struct {
	ID       int
	Customer string
	Amount   float64
}

// обработка одного заказа
func processOrder(order Order) {
	fmt.Printf("Обработка заказа %d для %s...\n", order.ID, order.Customer)

	// имитация работы: валидация + отправка на склад = 1 секунда
	time.Sleep(1 * time.Second)

	fmt.Printf("Заказ %d на сумму %.2f успешно оформлен\n", order.ID, order.Amount)
}

func main() {
	orders := []Order{
		{ID: 1, Customer: "Иван Петров", Amount: 1500.00},
		{ID: 2, Customer: "Мария Сидорова", Amount: 2300.50},
		{ID: 3, Customer: "Пётр Иванов", Amount: 899.99},
		{ID: 4, Customer: "Анна Смирнова", Amount: 4500.00},
		{ID: 5, Customer: "Сергей Волков", Amount: 1200.75},
	}

	var wg sync.WaitGroup

	// каждый заказ в отдельной горутине
	for _, order := range orders {
		wg.Add(1)

		// передаём order как аргумент
		// без этого все горутины могли бы увидеть последний заказ
		go func(o Order) {
			defer wg.Done()
			processOrder(o)
		}(order)
	}

	// ждём завершения всех 5 горутин
	wg.Wait()
	fmt.Println("Все заказы обработаны!")
}