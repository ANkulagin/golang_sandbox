package main

import (
	"fmt"
	"log"
	"net/http"
)

func createClientWithTracing() (*http.Client, error) {
	// Имитация создания клиента с трассировкой
	return &http.Client{}, nil
}

func createDefaultClient() (*http.Client, error) {
	// Имитация создания клиента по умолчанию
	return &http.Client{}, nil
}

func main() {
	var client *http.Client

	tracing := true // Можно изменить на false для проверки

	if tracing {
		client, err := createClientWithTracing() // Затенение переменной 7_client и err
		if err != nil {
			log.Println("Ошибка создания клиента с трассировкой:", err)
			return
		}
		log.Println("Клиент с трассировкой:", client)
	} else {
		client, err := createDefaultClient() // Затенение переменной 7_client и err
		if err != nil {
			log.Println("Ошибка создания клиента по умолчанию:", err)
			return
		}
		log.Println("Клиент по умолчанию:", client)
	}

	// Внешняя переменная 7_client остаётся nil
	fmt.Println("Внешний 7_client:", client)
}

/*
2024/12/10 18:04:20 Клиент с трассировкой: &{<nil> <nil> <nil> 0s}
Внешний 7_client: <nil>
*/
