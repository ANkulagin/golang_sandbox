package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

// Функция для создания и инициализации подключения к базе данных
func createDB(dataSourceName string) (*sql.DB, error) {
	// Открытие соединения с базой данных
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("не удалось открыть соединение с базой данных: %w", err)
	}

	// Проверка доступности базы данных
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("не удалось подключиться к базе данных: %w", err)
	}

	return db, nil
}

func main() {
	// Получение строки подключения из переменной окружения
	dataSourceName := os.Getenv("MYSQL_DATA_SOURCE_NAME")
	if dataSourceName == "" {
		log.Fatal("переменная окружения MYSQL_DATA_SOURCE_NAME не установлена")
	}

	// Инициализация подключения к базе данных
	db, err := createDB(dataSourceName)
	if err != nil {
		log.Fatalf("ошибка инициализации базы данных: %v", err)
	}
	defer db.Close() // Закрытие соединения при завершении программы

	// Использование инициализированного подключения
	err = db.Ping()
	if err != nil {
		log.Println("Ошибка подключения к базе данных:", err)
		return
	}
	log.Println("Подключение к базе данных успешно")
}
