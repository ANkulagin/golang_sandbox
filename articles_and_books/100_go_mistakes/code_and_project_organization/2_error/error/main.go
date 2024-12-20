package main

import (
	"errors"
	"fmt"
	"log"
)

// Функция для объединения двух строк с проверкой
func join(s1, s2 string, max int) (string, error) {
	if s1 == "" {
		return "", errors.New("s1 is empty")
	} else {
		if s2 == "" {
			return "", errors.New("s2 is empty")
		} else {
			concat, err := concatenate(s1, s2) // Вызов функции concatenate
			if err != nil {
				return "", err
			} else {
				if len(concat) > max {
					return concat[:max], nil
				} else {
					return concat, nil
				}
			}
		}
	}
}

// Вспомогательная функция для конкатенации строк
func concatenate(s1 string, s2 string) (string, error) {
	// Простая реализация конкатенации
	return s1 + s2, nil
}

func main() {
	result, err := join("Hello, ", "World!", 20)
	if err != nil {
		log.Println("Ошибка:", err)
		return
	}
	fmt.Println("Результат:", result)
}
