package client

import (
	"fmt"

	"github.com/ANkulagin/golang_sandbox/articles_and_books/100_go_mistakes/code_and_project_organization/6_error/correction/store"
)

// Интерфейс, определенный на стороне потребителя
type customersGetter interface {
	GetAllCustomers() ([]store.Customer, error)
}

// Структура Foo, использующая интерфейс customersGetter
type Foo struct {
	getter customersGetter
}

// Фабричная функция для создания Foo
func NewFoo(g customersGetter) Foo {
	return Foo{getter: g}
}

func (f Foo) Bar() {
	customers, err := f.getter.GetAllCustomers()
	if err != nil {
		fmt.Println("Ошибка получения клиентов:", err)
		return
	}
	for _, customer := range customers {
		fmt.Printf("Клиент ID: %s, Баланс: %.2f, Контракт: %t\n", customer.ID, customer.Balance, customer.Contract)
	}
}
