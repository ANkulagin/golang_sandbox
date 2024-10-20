package main

import (
	"fmt"
)

func main() {
	// 1. Массивы: инициализация, доступ и изменение
	fmt.Println("Массивы:")
	arr := [5]int{1, 2, 3, 4}
	fmt.Println("Исходный массив:", arr)
	arr[4] = 5
	// Изменение значения массива
	arr[0] = 10
	fmt.Println("Изменённый массив:", arr)

	// Передача массива в функцию (копируется)
	modifyArray(arr)
	fmt.Println("Массив после передачи в функцию (без изменения):", arr)

	// 2. Срезы: создание, добавление, доступ, инициализация слайсов и поддиапазоны
	fmt.Println("\nСрезы:")
	slice := []int{1, 2, 3}
	fmt.Println("Исходный срез:", slice)

	// Добавление элементов с помощью append
	slice = append(slice, 4, 5)
	fmt.Println("Срез после добавления элементов:", slice)

	// Срез поддиапазона
	subSlice := slice[1:3]
	fmt.Println("Срез поддиапазона [1:3]:", subSlice)

	// Изменение среза через функцию (ссылка)
	modifySlice(slice)
	fmt.Println("Срез после изменения через функцию:", slice)

	// 3. Многомерные массивы и срезы
	fmt.Println("\nМногомерные массивы и срезы:")
	matrix := [2][2]int{
		{1, 2},
		{3, 4},
	}
	fmt.Println("Многомерный массив:", matrix)

	sliceOfSlices := [][]int{
		{1, 2, 3},
		{3, 4},
	}
	fmt.Println("Срез срезов:", sliceOfSlices)

	// 4. Карты: создание, добавление, удаление и проверка наличия ключей
	fmt.Println("\nКарты:")
	m := map[string]int{
		"one": 1,
		"two": 2,
	}
	fmt.Println("Исходная карта:", m)

	// Добавление ключей
	m["three"] = 3
	fmt.Println("Карта после добавления элемента:", m)

	// Проверка наличия ключа (comma ok)
	value, ok := m["four"]
	if !ok {
		fmt.Println("Ключ 'four' не найден в карте")
	} else {
		fmt.Println("Значение ключа 'four':", value)
	}

	// Удаление элемента
	delete(m, "two")
	fmt.Println("Карта после удаления ключа 'two':", m)

	// Итерация по карте
	fmt.Println("Итерация по карте:")
	for key, value := range m {
		fmt.Printf("%s: %d\n", key, value)
	}
}

// Функция для модификации массива (по значению, изменений не произойдёт)
func modifyArray(arr [5]int) {
	arr[0] = 100
	fmt.Println("Массив внутри функции (локальная копия):", arr)
}

// Функция для модификации среза (по ссылке)
func modifySlice(slice []int) {
	slice[0] = 100
	fmt.Println("Срез внутри функции (по ссылке):", slice)
}

/*
Массивы:
Исходный массив: [1 2 3 4 0]
Изменённый массив: [10 2 3 4 5]
Массив внутри функции (локальная копия): [100 2 3 4 5]
Массив после передачи в функцию (без изменения): [10 2 3 4 5]

Срезы:
Исходный срез: [1 2 3]
Срез после добавления элементов: [1 2 3 4 5]
Срез поддиапазона [1:3]: [2 3]
Срез внутри функции (по ссылке): [100 2 3 4 5]
Срез после изменения через функцию: [100 2 3 4 5]

Многомерные массивы и срезы:
Многомерный массив: [[1 2] [3 4]]
Срез срезов: [[1 2 3] [3 4]]

Карты:
Исходная карта: map[one:1 two:2]
Карта после добавления элемента: map[one:1 three:3 two:2]
Ключ 'four' не найден в карте
Карта после удаления ключа 'two': map[one:1 three:3]
Итерация по карте:
one: 1
three: 3
*/
