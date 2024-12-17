package main

import (
	"fmt"
	"sync"
	"time"
)

func gather(funcs []func() any) []any {
	// Выполните все функции параллельно, соберите результаты
	var wg sync.WaitGroup
	result := make([]any, len(funcs))

	for i, f := range funcs {
		wg.Add(1)
		go func(id int, f func() any) {
			defer wg.Done()
			result[id] = f()
		}(i, f)
	}
	wg.Wait()

	return result
}

func squared(n int) func() any {
	return func() any {
		time.Sleep(time.Duration(n) * 100 * time.Millisecond)
		return n * n
	}
}

func main() {
	funcs := []func() any{
		squared(2), squared(3), squared(4),
	}

	start := time.Now()
	nums := gather(funcs)
	elapsed := float64(time.Since(start)) / 1_000_000

	fmt.Println(nums)
	fmt.Printf("Took %.0f ms\n", elapsed)
}
