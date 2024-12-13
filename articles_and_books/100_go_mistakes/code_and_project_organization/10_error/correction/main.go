package main

import (
	"fmt"
	correction10 "github.com/ANkulagin/golang_sandbox/articles_and_books/100_go_mistakes/code_and_project_organization/10_error/correction/correction_10"
)

func main() {
	store := correction10.NewInMem()
	// Получение значения
	value, exists := store.Get("foo")
	if exists {
		fmt.Printf("Ключ 'foo' имеет значение %d\n", value)
	}
	//Теперь нельзя
	//	store.mu.Lock() // Ошибка: mu приват
	//	store.mu.Unlock() // Ошибка: mu приват

	// Правильное использование методов Store
}
