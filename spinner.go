package main

import (
	"fmt"
	"time"
)

/*
* Домашнее заднание 7
* Заставить спиннер вращаться в течение некоего времени
* Петр, 22.05.2019
*/

func main() {
	go spinner(50 * time.Millisecond)

	fmt.Printf("Начинаем крутить спиннер: %s", time.Now().Format("15:04:05\n\r"))
	sleep(5)
	fmt.Printf("Закончили крутить спиннер: %s", time.Now().Format("15:04:05\n\r"))

	//вариант с каналами
	fmt.Printf("Начинаем крутить спиннер: %s", time.Now().Format("15:04:05\n\r"))
	sleepOverChannel(10)
	fmt.Printf("Закончили крутить спиннер: %s", time.Now().Format("15:04:05\n\r"))
}

func spinner(delay time.Duration) {
	for {
		for _, r := range "-\\|/" {
			fmt.Printf("%c\r", r)
			time.Sleep(delay)
		}
	}
}

func sleep(seconds int) {
	time.Sleep(time.Duration(seconds) * time.Second)
}

func sleepOverChannel(seconds int) {
	tick := make(<-chan time.Time)    // создаем однонаправленный канал
	tick = time.Tick(time.Duration(seconds) * time.Second) // создаём поток 10 секундных "тиков"
	<-tick // получим зи канала через 10 секунд
}