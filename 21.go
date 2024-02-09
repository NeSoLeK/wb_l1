package main

import (
	"fmt"
)

// Структура Printer

type Printer struct{}

// Метод Print для Printer

func (p *Printer) Print(text string) {
	fmt.Println("Printing:", text)
}

// Адаптер для преобразования Printer.Print в метод Write

type PrinterAdapter struct {
	Printer *Printer
}

// Метод Write для адаптера

func (pa *PrinterAdapter) Write(data []byte) (int, error) {
	// Преобразуем срез байтов в строку и вызываем метод Print у Printer
	pa.Printer.Print(string(data))
	// Возвращаем количество байтов и nil ошибку, так как метод Print не возвращает ошибку
	return len(data), nil
}

func t21() {
	// Создаем экземпляр Printer
	printer := &Printer{}

	// Создаем адаптер
	adapter := &PrinterAdapter{Printer: printer}

	// Вызываем Write через адаптер
	bytesWritten, err := adapter.Write([]byte("Hello, world!"))

	if err == nil {
		fmt.Printf("Записано %d байт\n", bytesWritten)
	} else {
		fmt.Println("Ошибка:", err)
	}
}
