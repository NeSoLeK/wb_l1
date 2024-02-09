package main

import (
	"fmt"
	"reflect"
)

// Функция для определения типа переменной
func checkType(v interface{}) {
	// Используем пакет reflect для определения типа переменной
	switch reflect.TypeOf(v).Kind() {
	case reflect.Int:
		fmt.Println("Type: int")
	case reflect.String:
		fmt.Println("Type: string")
	case reflect.Bool:
		fmt.Println("Type: bool")
	case reflect.Chan:
		fmt.Println("Type: channel")
	default:
		fmt.Println("Unknown type")
	}
}

func t14() {

	// Проверяем типы переменных
	checkType("123123")
	checkType(123)
	checkType(true)
	checkType(make(chan int))
}
