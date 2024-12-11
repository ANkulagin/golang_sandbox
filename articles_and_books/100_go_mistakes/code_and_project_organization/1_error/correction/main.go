package main

import (
	"fmt"
	"log"
	"net/http"
)

func CreateClientWithTracing() (*http.Client, error) {
	return &http.Client{}, nil
}

func CreateDefaultClient() (*http.Client, error) {
	return &http.Client{}, nil
}

func main() {
	var client *http.Client
	var err error // Объявляем переменную err заранее

	tracing := true // Можно изменить на false для проверки

	if tracing {
		client, err = CreateClientWithTracing() // Присваиваем внешним переменным
	} else {
		client, err = CreateDefaultClient() // Присваиваем внешним переменным
	}

	if err != nil {
		log.Println("Ошибка при создании клиента:", err)
		return
	}

	log.Println("Используемый клиент:", client)
	// Внешняя переменная client имеет корректное значение
	fmt.Println("Внешний client:", client)
}

/*
2024/12/10 18:07:16 Используемый клиент: &{<nil> <nil> <nil> 0s}
Внешний client: &{<nil> <nil> <nil> 0s}

*/
