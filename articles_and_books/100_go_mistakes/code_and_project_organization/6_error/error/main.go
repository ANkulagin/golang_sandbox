package main

import (
	"github.com/ANkulagin/golang_sandbox/articles_and_books/100_go_mistakes/code_and_project_organization/6_error/correction/client"
	"github.com/ANkulagin/golang_sandbox/articles_and_books/100_go_mistakes/code_and_project_organization/6_error/correction/store"
)

func main() {
	s := &store.MySQLStore{}
	foo := client.NewFoo(s)
	foo.Bar()
}

/*
Клиент ID: 1, Баланс: 100.00, Контракт: true
Клиент ID: 2, Баланс: -50.00, Контракт: false

*/
