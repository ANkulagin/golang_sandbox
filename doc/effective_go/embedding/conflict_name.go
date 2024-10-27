package main

import "fmt"

type A struct {
	Name string
}

type B struct {
	Name string
}

type C struct {
	A // Встраивание A
	B // Встраивание B
}

type D struct{}

func (D) Print() {
	fmt.Println("Print from D")
}

type E struct{}

func (E) Print() {
	fmt.Println("Print from E")
}

type F struct {
	D
	E
}

type Inner struct {
	Name string
}

type Middle struct {
	Inner
}

type Outer struct {
	Name string // Поле Name на верхнем уровне
	Middle
}

func main() {
	c := C{}
	// Конфликт имён! Нельзя однозначно определить, к какому Name обращаться.
	// c.Name = "Test" // Ошибка компиляции

	// Чтобы решить конфликт, нужно обратиться к полю через встроенный тип:
	c.A.Name = "Name from A"
	c.B.Name = "Name from B"

	fmt.Println("A.Name:", c.A.Name) // Вывод: A.Name: Name from A
	fmt.Println("B.Name:", c.B.Name) // Вывод: B.Name: Name from B

	//-----------------------------------------------------------------

	f := F{}
	// f.Print() // Ошибка компиляции: конфликт методов

	// Решение конфликта: явно вызвать метод через встроенный тип
	f.D.Print() // Вывод: Print from D
	f.E.Print() // Вывод: Print from E

	//------------------------

	outer := Outer{
		Name: "Outer Name",
		Middle: Middle{
			Inner: Inner{
				Name: "Inner Name",
			},
		},
	}

	fmt.Println("Outer Name:", outer.Name)              // Доступ к Name верхнего уровня
	fmt.Println("Inner Name:", outer.Middle.Inner.Name) // Доступ к Name на нижнем уровне через вложенные структуры
}

/*
A.Name: Name from A
B.Name: Name from B
Print from D
Print from E
Outer Name: Outer Name
Inner Name: Inner Name
*/
