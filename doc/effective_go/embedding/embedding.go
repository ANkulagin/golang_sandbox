package main

import (
	"fmt"
	"log"
	"os"
)

// Пример интерфейсов io.Reader и io.Writer
type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

// Объединение интерфейсов Reader и Writer в новый интерфейс ReadWriter
type ReadWriter interface {
	Reader
	Writer
}

// Пример структуры с встраиванием
type CustomReader struct{}

func (r *CustomReader) Read(p []byte) (n int, err error) {
	copy(p, "Hello from CustomReader")
	return len(p), nil
}

type CustomWriter struct{}

func (w *CustomWriter) Write(p []byte) (n int, err error) {
	fmt.Printf("CustomWriter: %s\n", string(p))
	return len(p), nil
}

type CustomReadWriter struct {
	*CustomReader
	*CustomWriter
}

func (rw *CustomReadWriter) ReadAndWrite() {
	buffer := make([]byte, 20)
	n, _ := rw.Read(buffer)
	rw.Write(buffer[:n])
}

// Пример встраивания структуры log.Logger в другую структуру
type Job struct {
	Command string
	*log.Logger
}

// Конструктор для структуры Job
func NewJob(command string) *Job {
	logger := log.New(os.Stderr, "Job: ", log.Ldate|log.Ltime)
	return &Job{Command: command, Logger: logger}
}

func (job *Job) PrintJobStatus() {
	job.Printf("Executing command: %s\n", job.Command)
}

func main() {
	fmt.Println("=== Пример с встраиванием интерфейсов ===")
	customRW := &CustomReadWriter{
		CustomReader: &CustomReader{},
		CustomWriter: &CustomWriter{},
	}

	customRW.ReadAndWrite()

	fmt.Println("\n=== Пример с встраиванием структур ===")
	job := NewJob("backup")
	job.Println("Запуск задачи...")
	job.PrintJobStatus()
}

/*
=== Пример с встраиванием интерфейсов ===
CustomWriter: Hello from CustomRea

=== Пример с встраиванием структур ===
Job: 2024/10/27 15:16:07 Запуск задачи...
Job: 2024/10/27 15:16:07 Executing command: backup

*/
