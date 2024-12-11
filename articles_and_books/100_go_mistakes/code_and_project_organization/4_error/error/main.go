package main

import (
	"fmt"
)

// Структура Customer с неэкспортированным полем balance
type Customer struct {
	balance float64
}

// Геттер для поля balance
func (c *Customer) Balance() float64 {
	return c.balance
}

// Сеттер для поля balance
func (c *Customer) SetBalance(amount float64) {
	c.balance = amount
}

func main() {
	customer := &Customer{}

	// Установка баланса через сеттер
	customer.SetBalance(100.0)

	// Получение баланса через геттер
	currentBalance := customer.Balance()
	fmt.Printf("Текущий баланс: %.2f\n", currentBalance)

	// Обнуление баланса через сеттер
	customer.SetBalance(0)
	fmt.Printf("Баланс после обнуления: %.2f\n", customer.Balance())
}
