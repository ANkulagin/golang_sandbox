package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// Функция инициализации для открытия базы данных
func init() {
	// Получение строки подключения из переменной окружения
	dataSourceName := os.Getenv("MYSQL_DATA_SOURCE_NAME")
	d, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Panic(err) // Паника при ошибке открытия базы данных
	}
	err = d.Ping()
	if err != nil {
		log.Panic(err) // Паника при ошибке пинга базы данных
	}
	db = d // Присвоение глобальной переменной db
}

func main() {
	// Использование глобальной переменной db
	// Пример запроса к базе данных
	err := db.Ping()
	if err != nil {
		log.Println("Ошибка подключения к базе данных:", err)
		return
	}
	log.Println("Подключение к базе данных успешно")
}
