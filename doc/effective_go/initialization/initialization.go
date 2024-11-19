package main

import (
	"fmt"
	"log"
	"os"
)

type ByteSize float64

const (
	_           = iota // Пропускаем первое значение
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	INT int = 1
)

var (
	home   = os.Getenv("HOME")
	user   = os.Getenv("USER")
	gopath = os.Getenv("GOPATH")
)

func init() {
	if user == "" {
		log.Fatal("$USER не установлен")
	}
	if home == "" {
		home = "/home/" + user
	}
	if gopath == "" {
		gopath = home + "/go"
	}
}

func (b ByteSize) String() string {
	switch {
	case b >= PB:
		return fmt.Sprintf("%.2fPB", b/PB)
	case b >= TB:
		return fmt.Sprintf("%.2fTB", b/TB)
	case b >= GB:
		return fmt.Sprintf("%.2fGB", b/GB)
	case b >= MB:
		return fmt.Sprintf("%.2fMB", b/MB)
	}
	return fmt.Sprintf("%.2fKB", b/KB)
}

func main() {
	// Работа с ByteSize
	fmt.Println(ByteSize(1e13)) // Выведет 9.09TB
	fmt.Printf("KB: %v, Type: %T\n", KB, KB)
	fmt.Printf("MB: %v, Type: %T\n", MB, MB)
	fmt.Printf("GB: %v, Type: %T\n", GB, GB)
	fmt.Printf("TB: %v, Type: %T\n", TB, TB)
	fmt.Printf("TB: %v, Type: %T\n", INT, INT)

	// Переменные окружения
	fmt.Println("\nПеременные окружения:")
	fmt.Println("HOME:", home)
	fmt.Println("USER:", user)
	fmt.Println("GOPATH:", gopath)
}

/*
9.09TB
KB: 1.00KB, Type: 2_example.ByteSize
MB: 1.00MB, Type: 2_example.ByteSize
GB: 1.00GB, Type: 2_example.ByteSize
TB: 1.00TB, Type: 2_example.ByteSize
TB: 1, Type: int

Переменные окружения:
HOME: /home/ankul
USER: ankul
GOPATH: /home/ankul/go

*/
