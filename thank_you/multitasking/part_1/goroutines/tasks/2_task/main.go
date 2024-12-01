/*
1 ВОПРОС:

	Правильно ли я понимаю, что так как мы закрываем канал в горутине, то цикл for message := range messages { ... } не начнёт свою работу, пока канал открыт?

Ответ:

	Нет, цикл for message := range messages { ... } начинает свою работу сразу же, как только он достигается в основном потоке. Этот цикл будет блокироваться при попытке чтения из канала messages, если в нем нет доступных данных. То есть, если горутина еще не отправила ничего в канал, основной поток будет ждать, пока данные не появятся.

	Цикл for ... range продолжает считывать значения из канала до тех пор, пока канал не будет закрыт и все значения из него не будут прочитаны. Как только канал закрыт и все данные получены, цикл автоматически завершается.
	Почему нужно закрывать канал

2 ВОПРОС:

	Почему вообще канал надо закрывать?

Ответ:
Закрытие канала важно по нескольким причинам:

  - Сигнализация о завершении отправки данных:
    Закрывая канал, вы сообщаете получателю, что больше не будет отправлено новых данных. Это позволяет получателю (в данном случае, циклу for ... range) корректно завершить свою работу.

  - Предотвращение блокировки:
    Если канал не закрыть, цикл for message := range messages { ... } будет ожидать новые данные бесконечно, даже если их больше не будет. Это приведет к блокировке программы.

  - Безопасность:
    Закрытие канала гарантирует, что никакие дополнительные данные не будут случайно отправлены в канал после завершения его использования, что может привести к панике (runtime panic).
*/
package main

import (
	"fmt"
	"strings"
	"sync"
	"unicode"
)

type counter map[string]int

func countDigits(str string) int {
	count := 0
	for _, char := range str {
		if unicode.IsDigit(char) {
			count++
		}
	}
	return count
}

func countDigitsInWords(phrase string) counter {
	words := strings.Fields(phrase)
	stats := make(counter)

	// начало решения
	messages := make(chan struct {
		word  string
		count int
	})

	var wg sync.WaitGroup
	wg.Add(1)
	// Запустите одну горутину.
	// В горутине пройдите по словам, посчитайте количество цифр в каждом,
	// и запишите результаты в канал counted.
	// В основной функции считайте значения из канала и заполните stats.
	go func() {
		for _, word := range words {
			count := countDigits(word)
			messages <- struct {
				word  string
				count int
			}{word, count}
		}
		defer wg.Done()
		close(messages)
	}()

	for message := range messages {
		stats[message.word] = message.count
	}
	wg.Wait()

	// конец решения
	return stats

}

func main() {
	input := "Hello 123 world 4567"
	stats := countDigitsInWords(input)
	for word, count := range stats {
		fmt.Printf("Word: %s, Digits: %d\n", word, count)
	}
}

/*
Word: Hello, Digits: 0
Word: 123, Digits: 3
Word: world, Digits: 0
Word: 4567, Digits: 4
*/
