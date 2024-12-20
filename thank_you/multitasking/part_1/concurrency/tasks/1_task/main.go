package main

import "fmt"

// count генерирует числа начиная со значения start и кладёт их в канал out,
// пока не будет получен сигнал отмены из done.
func count(done <-chan struct{}, start int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out) // Закрываем канал по окончании горутины
		for i := start; ; i++ {
			select {
			case <-done:
				return
			case out <- i:
			}
		}
	}()
	return out
}

// take читает первые n чисел из канала in и перенаправляет их в out,
// пока не будет достигнуто n или не придёт сигнал отмены. Если канал in
// закроется раньше времени, просто завершимся.
func take(done <-chan struct{}, in <-chan int, n int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out) // Закрываем канал по окончании горутины
		for i := 0; i < n; i++ {
			select {
			case <-done:
				return
			case val, ok := <-in:
				if !ok {
					// Канал in закрылся раньше, чем мы успели взять n элементов
					return
				}
				out <- val
			}
		}
	}()
	return out
}

func main() {
	done := make(chan struct{})
	defer close(done)

	stream := take(done, count(done, 10), 5)

	first := <-stream
	second := <-stream
	third := <-stream

	fmt.Println(first, second, third)
}
