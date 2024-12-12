package __client

import (
	"fmt"
	"github.com/ANkulagin/golang_sandbox/articles_and_books/100_go_mistakes/code_and_project_organization/7_error/error/7_store"
)

type Foo struct {
	storage __store.CustomerStorage
}

func NewFoo(storage __store.CustomerStorage) *Foo {
	return &Foo{storage: storage}
}

func (f *Foo) Bar() {
	// Пример реализации метода Bar
	fmt.Println("Bar called")
}
