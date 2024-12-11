package main

import (
	"fmt"
)

// Экспортированное поле Balance
type Customer struct {
	Balance float64
}

func main() {
	customer := &Customer{}

	// Прямой доступ к полю Balance для установки значения
	customer.Balance = 100.0

	// Прямой доступ к полю Balance для получения значения
	fmt.Printf("Текущий баланс: %.2f\n", customer.Balance)

	// Прямой доступ к полю Balance для обнуления
	customer.Balance = 0
	fmt.Printf("Баланс после обнуления: %.2f\n", customer.Balance)
}
