package main

import (
	"fmt"
	"time"
)

// --- Пример 1: Горутины ---
func exampleGoroutine() {
	fmt.Println("Пример 1: Горутины")
	go func() {
		fmt.Println("Привет из горутины!")
	}()
	time.Sleep(500 * time.Millisecond) // Даем горутине время завершиться
}

// --- Пример 2: Замыкание и горутины ---
func exampleClosureGoroutine() {
	fmt.Println("\nПример 2: Замыкание и горутины")
	message := "Отложенное сообщение из горутины"
	delay := 1 * time.Second

	go func(msg string) {
		time.Sleep(delay)
		fmt.Println(msg)
	}(message)

	time.Sleep(2 * time.Second) // Даём горутине время завершиться
}

// --- Пример 3: Небуферизованный канал ---
func exampleUnbufferedChannel() {
	fmt.Println("\nПример 3: Небуферизованный канал")
	c := make(chan string)

	go func() {
		c <- "Сообщение из горутины в небуферизованном канале"
	}()

	message := <-c
	fmt.Println(message)
}

// --- Пример 4: Буферизованный канал ---
func exampleBufferedChannel() {
	fmt.Println("\nПример 4: Буферизованный канал")
	c := make(chan string, 2) // Создаем канал с буфером на 2 сообщения

	c <- "Сообщение 1"
	c <- "Сообщение 2"

	fmt.Println(<-c) // Чтение первого сообщения
	fmt.Println(<-c) // Чтение второго сообщения
}

// --- Пример 5: Семафор с использованием буферизованного канала ---
func exampleSemaphore() {
	fmt.Println("\nПример 5: Семафор с использованием буферизованного канала")

	sem := make(chan struct{}, 2) // Канал, ограничивающий одновременно выполняющиеся задачи

	for i := 1; i <= 5; i++ {
		sem <- struct{}{} // Блокирует, если одновременно выполняются 2 задачи
		go func(id int) {
			defer func() { <-sem }() // Освобождает слот после завершения
			fmt.Printf("Выполняется задача %d\n", id)
			time.Sleep(1 * time.Second)
		}(i)
	}

	time.Sleep(3 * time.Second) // Даём задачам время завершиться
}

// --- Пример 6: Параллельные вычисления ---
func exampleParallelComputation() {
	fmt.Println("\nПример 6: Параллельные вычисления")

	data := []int{1, 2, 3, 4, 5, 6, 7, 8}
	c := make(chan int)

	// Разделяем данные на 2 части и обрабатываем параллельно
	go func() {
		sum := 0
		for _, v := range data[:len(data)/2] {
			sum += v
		}
		c <- sum // Отправляем результат в канал
	}()

	go func() {
		sum := 0
		for _, v := range data[len(data)/2:] {
			sum += v
		}
		c <- sum // Отправляем результат в канал
	}()

	// Суммируем результаты
	total := <-c + <-c
	fmt.Println("Сумма всех элементов:", total)
}

// --- Пример 7: Каналы каналов для обратной связи ---
type Request struct {
	number     int
	resultChan chan int
}

func exampleChannelsOfChannels() {
	fmt.Println("\nПример 7: Каналы каналов для обратной связи")

	request := Request{number: 5, resultChan: make(chan int)}

	go func(req Request) {
		req.resultChan <- req.number * 2 // Отправляем результат обратно через канал
	}(request)

	result := <-request.resultChan
	fmt.Println("Результат обработки запроса:", result)
}

// --- Пример 8: Leaky Buffer для повторного использования объектов ---
type Buffer struct {
	data string
}

var freeList = make(chan *Buffer, 2) // Буфер для повторного использования

func exampleLeakyBuffer() {
	fmt.Println("\nПример 8: Leaky Buffer для повторного использования объектов")

	// Получение буфера из freeList или создание нового
	var buffer *Buffer
	select {
	case buffer = <-freeList: // Пытаемся взять буфер из freeList
		fmt.Println("Получили буфер из freeList")
	default:
		buffer = &Buffer{data: "Новый буфер"} // Создаем новый, если свободных нет
	}

	fmt.Println("Используем буфер с данными:", buffer.data)

	// Возвращаем буфер в freeList, если там есть место
	select {
	case freeList <- buffer:
		fmt.Println("Вернули буфер в freeList")
	default:
		fmt.Println("freeList полон, буфер не возвращён")
	}
}

// --- Основная функция для запуска всех примеров ---
func main() {
	exampleGoroutine()
	exampleClosureGoroutine()
	exampleUnbufferedChannel()
	exampleBufferedChannel()
	exampleSemaphore()
	exampleParallelComputation()
	exampleChannelsOfChannels()
	exampleLeakyBuffer()
}

/*
Пример 1: Горутины
Привет из горутины!

Пример 2: Замыкание и горутины
Отложенное сообщение из горутины

Пример 3: Небуферизованный канал
Сообщение из горутины в небуферизованном канале

Пример 4: Буферизованный канал
Сообщение 1
Сообщение 2

Пример 5: Семафор с использованием буферизованного канала
Выполняется задача 2
Выполняется задача 1
Выполняется задача 3
Выполняется задача 4
Выполняется задача 5

Пример 6: Параллельные вычисления
Сумма всех элементов: 36

Пример 7: Каналы каналов для обратной связи
Результат обработки запроса: 10

Пример 8: Leaky Buffer для повторного использования объектов
Используем буфер с данными: Новый буфер
Вернули буфер в freeList

*/
