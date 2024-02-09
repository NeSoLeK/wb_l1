package main

import "fmt"

//Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

func Bit(val int64, pos uint, bitVal bool) int64 {
	if bitVal {
		return val | (1 << pos)
	} else {
		return val &^ (1 << pos)
	}
}

func t8() {
	var val int64 = 50
	fmt.Printf("Начальное значение: %d\n", val)

	updVal := Bit(val, 3, true)
	fmt.Printf("Значение, после добавления 3 бита в 1: %d\n", updVal)

	updVal = Bit(val, 5, false)
	fmt.Printf("Значение, после добавления 5 бита в 0: %d\n", updVal)
}
