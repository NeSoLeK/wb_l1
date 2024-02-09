package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

//Реализовать постоянную запись данных в канал (главный поток). Реализовать
//набор из N воркеров, которые читают произвольные данные из канала и
//выводят в stdout. Необходима возможность выбора количества воркеров при
//старте.

//go run 4.go -workers=15
//default workers = 5

func main() {
	workers := flag.Int("workers", 5, "an int")
	flag.Parse()
	t4(workers)
}

func t4(workers *int) {
	mainChannel := make(chan time.Time)

	wg := new(sync.WaitGroup)

	c := make(chan os.Signal)
	// Создается хендлер сигнала прерывания
	// который при срабатывание отправляет сигнал в канал
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	//https://gobyexample.com/signals

	// Запускаем воркеры
	for i := 1; i <= *workers; i++ {
		//Добавление одной горутины в очередь
		wg.Add(1)
		go worker(mainChannel, wg)
	}

	// В этой горутине добавляются данные в главный канал
	// Срабатывает каждые 3 секунды, передает time.Time
	go func() {
		count := 1
		for tick := range time.Tick(3 * time.Second) {
			// отправляю данные в общий канал
			mainChannel <- tick
			count++
		}
	}()

	// Главный поток в формате ожидания ctrl+c
	<-c
	close(mainChannel)
	//Ожидаем wg.Done для каждой функции
	wg.Wait()

}

// Воркер, который читает данные из канала и выводит их
func worker(ch <-chan time.Time, wg *sync.WaitGroup) {
	defer wg.Done()
	for {

		data, ok := <-ch
		// Если if отрабатывает значит данные в канале были все прочитаны и он закрыт
		if !ok {
			return
		}

		fmt.Println(data)
	}
}
