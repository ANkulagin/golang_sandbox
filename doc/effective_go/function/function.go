package main

import (
	"fmt"
	"io"
	"os"
)

// Функция с множественными возвращаемыми значениями
func nextInt(b []byte, i int) (int, int) {
	for ; i < len(b) && !isDigit(b[i]); i++ {
	}
	x := 0
	for ; i < len(b) && isDigit(b[i]); i++ {
		x = x*10 + int(b[i]) - '0'
	}
	return x, i
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

// Использование defer для закрытия файла
func Contents(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	var result []byte
	buf := make([]byte, 100)
	for {
		n, err := f.Read(buf[0:])
		result = append(result, buf[0:n]...)
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}
	}
	return string(result), nil
}

func main() {
	// Пример множественного возвращаемого значения
	b := []byte("12345")
	i := 0
	for i < len(b) {
		x, nextPos := nextInt(b, i)
		fmt.Println("Value:", x, "Next Position:", nextPos)
		i = nextPos
	}

	// Тестирование defer
	content, err := Contents("/home/ankul/project/golang_sandbox/doc/effective_go/function/test.txt")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("File content:", content)
	}
}
