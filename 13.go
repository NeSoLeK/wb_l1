package main

import "fmt"

// Синтаксический способ
func t13_1(a int, b int) {

	a, b = b, a

	fmt.Println(a)
	fmt.Println(b)
}

// Математический способ
func t13_2(a int, b int) {

	a += b
	b = a - b
	a = a - b

	fmt.Println(a)
	fmt.Println(b)
}
