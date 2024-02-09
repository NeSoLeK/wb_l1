package main

import "fmt"

// Функция быстрой сортировки
func quicksort(arr []int) []int {
	// Базовый случай: если длина массива меньше или равна 1, возвращаем его
	if len(arr) <= 1 {
		return arr
	}

	// Выбираем опорный элемент (pivot)
	pivot := arr[len(arr)/2]

	// Создаем три вспомогательных массива для элементов меньше, больше и равных опорному элементу
	var less, greater, equal []int

	// Разбиваем массив на три группы: элементы меньше, равные и больше опорного элемента
	for _, num := range arr {
		if num < pivot {
			less = append(less, num)
		} else if num > pivot {
			greater = append(greater, num)
		} else {
			equal = append(equal, num)
		}
	}

	// Рекурсивно сортируем подмассивы меньших и больших элементов
	less = quicksort(less)
	greater = quicksort(greater)

	// Объединяем отсортированные подмассивы вместе с равными элементами и возвращаем результат
	return append(append(less, equal...), greater...)
}

func t16() {
	// Пример неотсортированного массива
	arr := []int{9, 2, 5, 7, 1, 4, 8, 3, 6}

	// Вызываем функцию быстрой сортировки
	sortedArr := quicksort(arr)

	fmt.Println("Отсортированный массив:", sortedArr)
}
