package correction_10

import (
	"sync"
)

// Структура InMem с не встроенным и не экспортируемым Mutex
type InMem struct {
	mu sync.Mutex
	m  map[string]int
}

// Конструктор для InMem
func NewInMem() *InMem {
	return &InMem{m: make(map[string]int)}
}

// Метод Get с использованием внутреннего Mutex
func (i *InMem) Get(key string) (int, bool) {
	i.mu.Lock()
	defer i.mu.Unlock()
	v, contains := i.m[key]
	return v, contains
}
