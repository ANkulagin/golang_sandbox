package main

import (
	"errors"
	"fmt"
	"os"
	"syscall"
)

// ErrorExample демонстрирует использование типа error.
func ErrorExample(filename string) {
	fmt.Println("\n--- Error Example ---")
	file, err := os.Open(filename)
	if err != nil {
		if pathErr, ok := err.(*os.PathError); ok {
			fmt.Printf("Ошибка: операция %s не удалась для файла %s: %v\n", pathErr.Op, pathErr.Path, pathErr.Err)
		} else {
			fmt.Printf("Ошибка при открытии файла: %v\n", err)
		}
		return
	}
	defer file.Close()
	fmt.Println("Файл успешно открыт:", file.Name())
}

// CustomErrorExample демонстрирует создание пользовательских ошибок.
func CustomErrorExample(value int) error {
	fmt.Println("\n--- Custom Error Example ---")
	if value < 0 {
		return errors.New("значение не может быть отрицательным")
	}
	fmt.Println("Значение корректно:", value)
	return nil
}

// PanicExample демонстрирует использование panic для критических ошибок.
func PanicExample(value int) {
	fmt.Println("\n--- Panic Example ---")
	if value < 0 {
		panic(fmt.Sprintf("Недопустимое значение: %d", value))
	}
	fmt.Println("Значение допустимо:", value)
}

// RecoverExample демонстрирует использование recover для восстановления после panic.
func RecoverExample(value int) {
	fmt.Println("\n--- Recover Example ---")
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Восстановление после panic: %v\n", r)
		}
	}()
	PanicExample(value)
	fmt.Println("Этот код не будет выполнен, если возникла panic")
}

// PathErrorExample демонстрирует проверку типа ошибки с использованием type assertion.
func PathErrorExample(filename string) {
	fmt.Println("\n--- Path Error Example ---")
	for try := 0; try < 2; try++ {
		file, err := os.Create(filename)
		if err == nil {
			fmt.Println("Файл успешно создан:", file.Name())
			file.Close()
			return
		}
		if e, ok := err.(*os.PathError); ok && e.Err == syscall.ENOSPC {
			fmt.Println("Нет места на устройстве. Попытка освободить место...")
			deleteTempFiles() // Функция для освобождения места.
			continue
		}
		fmt.Printf("Ошибка при создании файла: %v\n", err)
		return
	}
}

// deleteTempFiles демонстрирует освобождение места на устройстве.
func deleteTempFiles() {
	fmt.Println("Удаление временных файлов для освобождения места...")
}

func main() {
	// Пример 1: Использование error для обработки ошибок
	ErrorExample("nonexistent_file.txt")

	// Пример 2: Пользовательские ошибки
	err := CustomErrorExample(-5)
	if err != nil {
		fmt.Println("Ошибка:", err)
	}
	CustomErrorExample(10)

	// Пример 3: Использование panic для критических ошибок
	// PanicExample(-1) // Раскомментируйте, чтобы увидеть работу panic (программа завершится)

	// Пример 4: Использование recover для восстановления после panic
	RecoverExample(-1)

	// Пример 5: Использование type assertion для проверки типа ошибки
	PathErrorExample("/dev/full") // /dev/full - специальное устройство, симулирующее "нет места"
}

/*
--- Error Example ---
Ошибка: операция open не удалась для файла nonexistent_file.txt: no such file or directory

--- Custom Error Example ---
Ошибка: значение не может быть отрицательным

--- Custom Error Example ---
Значение корректно: 10

--- Recover Example ---

--- Panic Example ---
Восстановление после panic: Недопустимое значение: -1

--- Path Error Example ---
Файл успешно создан: /dev/full

*/
