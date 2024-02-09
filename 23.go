package main

//Удалить i-ый элемент из слайса.

import "fmt"

func t23() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	i := 5
	//Перезапись без x[i]
	arr = append(arr[:i], arr[i+1:]...)

	fmt.Println(arr)
}
