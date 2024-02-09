package main

import (
	"fmt"
	"sync"
)

type Count struct {
	value int
	mt    sync.RWMutex // Встраивание Mutex
}

// Метод Add()

func (c *Count) Add() {
	//Блокируем доступ для других горутин
	c.mt.Lock()
	c.value++
	c.mt.Unlock()
}

// Метод PrintValue()
func (c *Count) PrintValue() int {
	c.mt.Lock()
	//Отложенная разблокировка
	defer c.mt.Unlock()
	return c.value
}

func t18() {
	var wg sync.WaitGroup
	count := Count{}
	//Создаем 20 горутин
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			count.Add()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Value:", count.PrintValue())
}
