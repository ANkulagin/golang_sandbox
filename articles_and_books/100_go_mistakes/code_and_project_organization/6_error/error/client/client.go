package client

import (
	"github.com/ANkulagin/golang_sandbox/articles_and_books/100_go_mistakes/code_and_project_organization/6_error/error/store"
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
	_, err := f.getter.GetAllCustomers()
	if err != nil {
		// Обработка ошибки
		return
	}
	// Работа с полученными клиентами
}
