package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

/*
* Домашнее заднание 7
* Добавить для time-сервера возможность его корректного завершения при вводе команды exit.
* Петр, 22.05.2019
*/

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Для остановки введите: exit")
	cancel := make(chan int, 2)
	go checkCommand(cancel)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn, cancel)
	}
}

func handleConn(c net.Conn, cancel chan int) {
	defer c.Close()

	tick := make(<-chan time.Time)
	tick = time.Tick(1 * time.Second)

	for {
		select {
		case <-tick:
			_, err := io.WriteString(c, time.Now().Format("15:04:05\n\r"))
			if err != nil {
				return
			}
		case <-cancel:
			_, err := io.WriteString(c, "the server is down\n\r")
			if err != nil {
				log.Print(err)
				os.Exit(1)
			}

			fmt.Println("Программа завершена")
			os.Exit(0)
		}
	}
}

func checkCommand(cancel chan int) {
	var command string
	fmt.Scanln(&command)

	if command == "exit" {
		fmt.Println("Программа завершается")
		cancel <- 1
		//завершить программу если соединений нет
		time.Sleep(2 * time.Second)
		fmt.Println("Программа завершена")
		os.Exit(0)
	} else {
		go checkCommand(cancel)
	}
}