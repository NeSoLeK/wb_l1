package main

import (
	"fmt"
	"time"
)

// Создаем функцию для отправки данных в канал
func send(ch chan<- int) {
	for i := 1; ; i++ {
		//Отправляем значения в канал
		ch <- i
		//Задрежка
		time.Sleep(time.Second)
	}
}

func t5() {
	//Кол-во времени работы (секунды)
	seconds := 10
	//Создаем канал для передачи данных
	channel := make(chan int)
	//Запускаем горутину для отправки
	go send(channel)
	//Запускаем таймер, который по истечению N секунд завершит программу
	timer := time.NewTimer(time.Duration(seconds) * time.Second)
	defer timer.Stop()

	//Ожидаем данные из канала или истечение таймера
	for {
		select {
		case data := <-channel:
			fmt.Printf("Полученные данные: %d\n", data)
		case <-timer.C:
			fmt.Println("Программа завершена!")
			return
		}
	}
}
