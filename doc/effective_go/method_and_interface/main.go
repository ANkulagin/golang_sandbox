package main

import (
	"fmt"
	"net/http"
	"os"
	"sort"
)

// ---------------- Методы для значений и указателей ----------------

// MySlice - тип, представляющий срез байтов
type MySlice []byte

// Append - метод, работающий с копией MySlice (значение)
func (slice MySlice) Append(data []byte) MySlice {
	return append(slice, data...)
}

// AppendPointer - метод, работающий с указателем на MySlice
func (p *MySlice) AppendPointer(data []byte) {
	*p = append(*p, data...)
}

// ---------------- Пример интерфейса ----------------

type Stringer interface {
	String() string
}

type Sequence []int

// Реализация интерфейса Stringer для Sequence
func (s Sequence) String() string {
	copyS := make(Sequence, len(s))
	copy(copyS, s)
	sort.Ints(copyS)
	return fmt.Sprint([]int(copyS))
}

// Len, Less и Swap для реализации sort.Interface
func (s Sequence) Len() int           { return len(s) }
func (s Sequence) Less(i, j int) bool { return s[i] < s[j] }
func (s Sequence) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// ---------------- Утверждение типа и конверсия ----------------

func typeAssertions(value interface{}) {
	// Type assertion (утверждение типа)
	if str, ok := value.(string); ok {
		fmt.Println("Это строка:", str)
	} else {
		fmt.Println("Значение не является строкой")
	}

	// Type switch (switch по типу)
	switch v := value.(type) {
	case string:
		fmt.Println("Это строка (switch):", v)
	case int:
		fmt.Println("Это целое число:", v)
	default:
		fmt.Println("Неизвестный тип")
	}
}

// ---------------- Пример HTTP-сервера ----------------

// SimpleCounter - простой HTTP-счётчик
type SimpleCounter int

func (ctr *SimpleCounter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	*ctr++
	fmt.Fprintf(w, "Counter = %d\n", *ctr)
}

// RequestNotifier - уведомление через канал
type RequestNotifier chan *http.Request

// метода ServeHTTP, который является реализацией интерфейса http.Handler
func (notifier RequestNotifier) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	notifier <- req
	fmt.Fprint(w, "Notification sent\n")
}

// FunctionHandler - использование функции в качестве обработчика
func FunctionHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Аргументы программы: %v\n", os.Args)
}

func main() {
	// ----------- Методы для значений и указателей -----------

	// Создание среза и добавление данных с помощью методов
	slice := MySlice{}
	fmt.Println("Оригинальный срез:", slice)

	// Использование метода для значения (возвращает новый срез)
	newSlice := slice.Append([]byte{1, 2, 3})
	fmt.Println("После Append (значение):", newSlice)

	// Использование метода для указателя (изменяет оригинальный срез)
	slice.AppendPointer([]byte{4, 5, 6})
	fmt.Println("После AppendPointer (указатель):", slice)

	// ------------- Интерфейсы и сортировка --------------

	seq := Sequence{5, 3, 1, 4, 2}
	fmt.Println("Оригинальная последовательность:", seq)
	fmt.Println("Отсортированная строковая версия:", seq.String())

	// -------------- Утверждение типа и конверсия --------------

	fmt.Println("\nType Assertion Examples:")
	typeAssertions("Hello, Go!")
	typeAssertions(123)
	typeAssertions(3.14)

	// -------------- HTTP-сервер ----------------

	fmt.Println("\nЗапуск HTTP-сервера на порту 8080...")

	// Счётчик посещений
	counter := new(SimpleCounter)
	http.Handle("/counter", counter)

	// Уведомление через канал
	notifier := make(RequestNotifier)
	http.Handle("/notify", notifier) //можем передать notifer только потому что он реализует интерфейс http.Handler
	go func() {
		for req := range notifier {
			fmt.Println("Получен запрос на:", req.URL.Path)
		}
	}()

	// Обработчик на основе функции
	http.Handle("/args", http.HandlerFunc(FunctionHandler))

	// Запуск сервера
	logErr := http.ListenAndServe(":8080", nil)
	if logErr != nil {
		fmt.Println("Ошибка при запуске сервера:", logErr)
	}
}

/*
Оригинальный срез: []
После Append (значение): [1 2 3]
После AppendPointer (указатель): [4 5 6]
Оригинальная последовательность: [1 2 3 4 5]
Отсортированная строковая версия: [1 2 3 4 5]

Type Assertion Examples:
Это строка: Hello, Go!
Это строка (switch): Hello, Go!
Значение не является строкой
Это целое число: 123
Значение не является строкой
Неизвестный тип

Запуск HTTP-сервера на порту 8080...
*/
