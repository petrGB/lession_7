package main

import (
	"fmt"
	"time"
)

/*
* Домашнее заднание 7
* Перепишите программу-конвейер, ограничив количество передаваемых для обработки значений и обеспечив корректное завершение всех горутин.
* Петр, 22.05.2019
*/

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// генерация
	go func() {
		for x := 0; x < 5; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	// возведение в квадрат
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	// печать
	for y := range squares {
		fmt.Println(y)
		time.Sleep(1 * time.Second)
	}
}
