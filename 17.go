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

// Функция бинарного поиска
func binarySearch(arr []int, target int) int {
	// Определяем границы для поиска
	left, right := 0, len(arr)-1

	// Ищем целевой элемент, пока левая граница не превысит правую
	for left <= right {
		// Определяем средний индекс
		mid := left + (right-left)/2

		// Если целевой элемент находится в середине массива, возвращаем его индекс
		if arr[mid] == target {
			return mid
		}

		// Если целевой элемент меньше элемента в середине, сдвигаем правую границу
		if target < arr[mid] {
			right = mid - 1
		} else { // Если целевой элемент больше элемента в середине, сдвигаем левую границу
			left = mid + 1
		}
	}

	// Если элемент не найден, возвращаем -1
	return -1
}

func t17() {
	// Пример отсортированного массива
	arr := []int{9, 2, 5, 7, 1, 4, 8, 3, 6}
	sortedArr := quicksort(arr)
	// Целевое значение для поиска
	target := 13

	// Выполняем бинарный поиск
	index := binarySearch(sortedArr, target)

	// Выводим результат поиска
	if index != -1 {
		fmt.Printf("Элемент %d найден в позиции %d\n", target, index)
	} else {
		fmt.Printf("Элемент %d не найден в массиве\n", target)
	}
}
