package main

//Разработать программу нахождения расстояния между двумя точками, которые
//представлены в виде структуры Point с инкапсулированными параметрами x,y и
//конструктором.

import (
	"fmt"
	"math"
)

// Структура Point представляет точку в двумерном пространстве

type Point struct {
	x, y float64
}

// Конструктор для создания новой точки

func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

// Метод для вычисления расстояния между двумя точками

func (p1 *Point) DistanceTo(p2 *Point) float64 {
	dx := p1.x - p2.x
	dy := p1.y - p2.y
	return math.Sqrt(dx*dx + dy*dy)
}

func t24() {
	// Создаем две точки
	point1 := NewPoint(0, 0)
	point2 := NewPoint(3, 4)

	// Вычисляем расстояние между ними
	distance := point1.DistanceTo(point2)

	fmt.Printf("Расстояние между точками (%.1f, %.1f) и (%.1f, %.1f) равно %.2f\n", point1.x, point1.y, point2.x, point2.y, distance)
}
