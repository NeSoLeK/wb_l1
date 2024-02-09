package main

import (
	"fmt"
	"math"
	"sort"
)

//Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0,
//15.5, 24.5, -21.0, 32.5. Объединить данные значения в группы с шагом в 10
//градусов. Последовательность в подмножноствах не важна.

func group(temps []float64, step float64) map[string][]float64 {
	groups := make(map[string][]float64)
	// Сортируем температуры по возрастанию
	sort.Float64s(temps)
	for _, temp := range temps {
		// Находим ключ (диапазон температур) для текущей температуры
		key := fmt.Sprintf("%.0f-%.0f", math.Floor(temp/step)*step, math.Ceil(temp/step)*step)
		groups[key] = append(groups[key], temp)
	}
	return groups
}

func t10() {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	//Шаг. По ТЗ = 10
	step := 10.0
	//Группируем
	groups := group(temps, step)
	for key, val := range groups {
		fmt.Printf("%s: %v\n", key, val)
	}
}
