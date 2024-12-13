package main

import (
	"fmt"
	"sync"
)

// Структура InMem с встроенным Mutex
type InMem struct {
	sync.Mutex
	m map[string]int
}

// Конструктор для InMem
func NewInMem() *InMem {
	return &InMem{m: make(map[string]int)}
}

// Метод Get с использованием встроенного Mutex
func (i *InMem) Get(key string) (int, bool) {
	i.Lock() // Прямой доступ к методу Lock
	defer i.Unlock()
	v, contains := i.m[key]
	return v, contains
}

func main() {
	store := NewInMem()

	// Установка значения
	store.m["foo"] = 42

	// Получение значения
	value, exists := store.Get("foo")
	if exists {
		fmt.Printf("Ключ 'foo' имеет значение %d\n", value)
	}

	// Нежелательный доступ к методам Mutex
	store.Lock()   // Неожиданный вызов Lock
	store.Unlock() // Неожиданный вызов Unlock

	// Эти вызовы могут привести к неправильной синхронизации
}
