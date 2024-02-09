package main

//Реализовать конкурентную запись данных в map.

import (
	"fmt"
	"sync"
)

// Структура MyMap, хранящая в себе map и mutex

type MyMap struct {
	data map[int]int
	mt   sync.RWMutex // Встраивание Mutex
}

//Мутод для добавления данных в map

func (mm *MyMap) AddData(i int, j int) {
	mm.mt.Lock()
	mm.data[i] = j
	mm.mt.Unlock()
}

func t7() {
	userdata := MyMap{data: make(map[int]int)}
	wg := new(sync.WaitGroup)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			userdata.AddData(i, i*2)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(userdata.data)
}
