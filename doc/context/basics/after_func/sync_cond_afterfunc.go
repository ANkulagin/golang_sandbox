package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	// Функция, которая ждёт выполнения условия или завершения контекста
	waitWithContext := func(ctx context.Context, cond *sync.Cond, conditionMet func() bool) error {
		// После завершения контекста пробуждаем все горутины
		stop := context.AfterFunc(ctx, func() {
			cond.L.Lock()
			defer cond.L.Unlock()
			cond.Broadcast() // Пробуждаем всех ожидающих
		})
		defer stop() // Останавливаем AfterFunc после завершения

		cond.L.Lock()
		defer cond.L.Unlock()

		// Пока условие не выполнено, продолжаем ждать
		for !conditionMet() {
			cond.Wait()           // Ждём сигнала
			if ctx.Err() != nil { // Если контекст завершён, возвращаем ошибку
				return ctx.Err()
			}
		}
		return nil
	}

	// Создаём sync.Cond для координации горутин
	cond := sync.NewCond(&sync.Mutex{})

	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			// Контекст с таймаутом
			ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
			defer cancel()

			// Ждём выполнения условия
			err := waitWithContext(ctx, cond, func() bool { return false })
			if err != nil {
				fmt.Printf("Горутина %d: %s\n", id, err)
			}
		}(i)
	}

	wg.Wait() // Ждём завершения всех горутин
	fmt.Println("Все горутины завершены.")
}

/*
Горутина 1: context deadline exceeded
Горутина 2: context deadline exceeded
Горутина 0: context deadline exceeded
Горутина 3: context deadline exceeded
Все горутины завершены.

*/
