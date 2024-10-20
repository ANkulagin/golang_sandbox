package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	_ "net/http/pprof" // Импорт для побочных эффектов
	"os"
)

// Структура для демонстрации реализации интерфейса json.Marshaler
type RawMessage struct {
	Message string `json:"message"`
}

// Реализация метода MarshalJSON для интерфейса json.Marshaler
func (rm *RawMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]string{
		"message": rm.Message,
	})
}

func main() {

	// 🔍 Проверка реализации интерфейса json.Marshaler
	rm := &RawMessage{Message: "Hello, Go!"}
	if _, ok := interface{}(rm).(json.Marshaler); ok {
		fmt.Printf("Тип *RawMessage реализует интерфейс json.Marshaler\n")
	} else {
		fmt.Printf("Тип *RawMessage не реализует интерфейс json.Marshaler\n")
	}
	// 📌 Пример с использованием пустого идентификатора в множественном присваивании
	path := "doc/effective_go/blank_identifier__/test.go"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Printf("%s не существует\n", path)
	}

	// ❌ Пример плохой практики (Игнорирование ошибок)
	// Попробуйте удалить этот блок и убедитесь, что программа выдаёт предупреждение о неиспользуемой ошибке.
	fi, _ := os.Stat(path)
	if fi != nil && fi.IsDir() {
		fmt.Printf("%s — это директория\n", path)
	}

	// ⚠️ Работа с неиспользуемыми импортами и переменными
	var _ = fmt.Printf // Используем для подавления ошибки неиспользуемого пакета fmt
	var _ io.Reader    // Используем для подавления ошибки неиспользуемого пакета io

	fd, err := os.Open("test.go")
	if err != nil {
		log.Fatal(err)
	}
	_ = fd // Подавление ошибки неиспользуемой переменной

	// 🔄 Импорт с побочными эффектами
	// net/http/pprof регистрирует обработчики HTTP для профилирования.
	fmt.Println("\nЗапуск HTTP-сервера для профилирования (на порту 6060)")
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	// Пример использования интерфейса json.Marshaler
	jsonData, err := json.Marshal(rm)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("JSON:", string(jsonData))

	// 📌 Пример использования пустого идентификатора в for range
	data := []int{1, 2, 3, 4, 5}
	fmt.Println("\nПеребор значений с использованием пустого идентификатора:")
	for _, value := range data {
		fmt.Println("Значение:", value)
	}

	// Ожидание завершения работы сервера
	fmt.Println("\nНажмите любую клавишу для завершения...")
	var input string
	fmt.Scanln(&input)
}
