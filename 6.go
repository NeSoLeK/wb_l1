package main

import (
	"fmt"
	"os"
	"runtime"
	"sync"
)

// С помощью сигнала
func t6_1() {
	//Создаем канала для передачи сигнала
	channel := make(chan bool)

	go func() {
		for {
			select {
			//При получении сигнала завершения выходим из горутины
			case <-channel:
				return
			default:
				//Горутина работает
			}
		}
	}()
	//Отправляем сигнал о завершении
	channel <- true
}

// runtime.Goexit()

func t6_2() {
	var wg sync.WaitGroup
	// Запускаем несколько горутин
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go t6_2_worker(i, &wg)
	}

	// Подождем, пока все горутины завершатся
	wg.Wait()

	fmt.Println("All Goroutines have completed.")
}

func t6_2_worker(id int, wg *sync.WaitGroup) {
	// Уведомляем, что эта горутина завершилась
	defer wg.Done()

	fmt.Printf("Goroutine %d is working...\n", id)

	//Some work
	for i := 0; i < 5; i++ {
		fmt.Printf("Goroutine %d is processing task %d\n", id, i)
	}

	// Выходим из горутины
	runtime.Goexit()
}

// Выйти из программы полностью
func t6_3() {
	for i := 1; i < 5; i++ {
		go func() {
			for {
				fmt.Println("Working...")
			}
		}()
	}
	os.Exit(1)

}
