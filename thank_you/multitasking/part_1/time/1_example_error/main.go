package main

import "fmt"

func rangeGen(start, stop int) <-chan int {
	out := make(chan int)
	go func() {
		for i := start; i < stop; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func main() {
	for n := range rangeGen(1, 10) {
		if n == 2 {
			break
		}
		fmt.Println(n)
	}
}

//Поскольку main() вышла из цикла на числе 42, то цикл ✓ внутри rangeGen() тоже не завершился. Он навсегда заблокировался на строчке ✔ при попытке отправить число 43 в канал out. Горутина зависла. Канал out тоже не закрылся, так что если бы от него зависели другие горутины — зависли бы и они.
