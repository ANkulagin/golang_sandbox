package main

import (
	"fmt"
)

// Интерфейс для получения баланса
type BalanceGetter interface {
	GetBalance() float64
}

// Интерфейс для установки баланса
type BalanceSetter interface {
	SetBalance(amount float64)
}

// Структура Customer с не экспортированным полем balance
type Customer struct {
	balance float64
}

// Реализация метода GetBalance
func (c *Customer) GetBalance() float64 {
	return c.balance
}

// Реализация метода SetBalance
func (c *Customer) SetBalance(amount float64) {
	c.balance = amount
}

func main() {
	// Использование интерфейсов для доступа к полям
	var getter BalanceGetter = &Customer{}
	var setter BalanceSetter = &Customer{}

	// Установка баланса через сеттер
	setter.SetBalance(150.0)

	// Получение баланса через геттер
	currentBalance := getter.GetBalance()
	fmt.Printf("Текущий баланс: %.2f\n", currentBalance)
}
