package main

import (
	"fmt"
	"sync"
)

//Написать программу, которая конкурентно рассчитает значение квадратов чисел
//взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

func t2() {
	array := [5]int{2, 4, 6, 8, 10}
	// Канал для передачи результатов (
	channel := make(chan int, len(array))

	// Итерируемся по массиву и передаем в горутину числа для вычисления их квадратов.
	for _, num := range array {
		go func(num int) {
			// Передаем квадраты числе в канал
			channel <- num * num
		}(num)
	}
	// В цикле выводим в консоль значения квадратов чисел из канала.
	for i := 0; i < len(array); i++ {
		fmt.Println(<-channel)
	}
}

func t2_2() {

	array := [5]int{2, 4, 6, 8, 10}

	wg := sync.WaitGroup{}
	//Итерация по всем элементам массива array
	for _, num := range array {
		//Добавление одной горутины в очередь
		wg.Add(1)
		go func(num int) {
			//Откладываем сообщение о закрытии функции к ее завершению
			defer wg.Done()
			fmt.Println(num * num)
		}(num)
	}
	//Ожидаем wg.Done для каждой функции
	wg.Wait()
}

//https://dev.to/karanpratapsingh/go-course-sync-package-5c3m
