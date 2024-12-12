package main

import (
	"fmt"
	__client "github.com/ANkulagin/golang_sandbox/articles_and_books/100_go_mistakes/code_and_project_organization/7_error/error/7_client"
	__store "github.com/ANkulagin/golang_sandbox/articles_and_books/100_go_mistakes/code_and_project_organization/7_error/error/7_store"
)

func main() {
	// Создание конкретной реализации хранилища
	s := __store.NewInMemoryStore()

	// Создание клиента Foo с использованием конкретного типа
	foo := __client.NewFoo(s)

	// Вызов метода Bar
	foo.Bar()

	// Пример использования хранилища напрямую
	customer := __store.Customer{
		ID:       "123",
		Balance:  250.75,
		Contract: true,
	}

	err := s.StoreCustomer(customer)
	if err != nil {
		fmt.Println("Ошибка сохранения клиента:", err)
		return
	}

	fmt.Println("Клиент успешно сохранен")
}
