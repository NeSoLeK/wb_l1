## 1. Какой самый эффективный способ конкатенации строк?

- Использование пакета `strings` с помощью типа `Builder`.

## 2. Что такое интерфейсы, как они применяются в Go?

- Интерфейсы в Go служат для абстрагирования данных и определения методов.

## 3. Чем отличаются RWMutex от Mutex?

- RWMutex позволяет параллельное чтение нескольким горутинам, в то время как Mutex разрешает доступ только одной горутине за раз.

## 4. Чем отличаются буферизированные и не буферизированные каналы?

- В не буферизированном канале отправка блокируется, пока данные не будут получены, в то время как в буферизированном канале отправка блокируется только при заполненном буфере.

## 5. Какой размер у структуры struct{}{}?

- 0 байт

## 6. Есть ли в Go перегрузка методов или операторов?

- Нет

## 7. В какой последовательности будут выведены элементы map[int]int?

- Порядок вывода элементов `map[int]int` зависит от способа итерации. Используя `fmt.Println`, элементы будут отсортированы, но при использовании цикла `range`, порядок может быть случайным.

## 8. В чем разница make и new?

- `make` используется для инициализации сложных структур данных, таких как каналы и мапы, в то время как `new` выделяет память для структур.

## 9. Сколько существует способов задать переменную типа slice или map?

- Два способа:
  1. с использованием `make`
  2. через объявление переменной

## 10. Что выведет данная программа и почему?

- 1 1
- Поскольку изменения копии указателя которая внутри функции `update` и `main` не затрагивает


```go
func update(p *int) {
	b := 2
	p = &b
}
func main() {
	var (
		a = 1
		p = &a
	)
	fmt.Println(*p)
	update(p)
	fmt.Println(*p)
}
```


## 11. Что выведет данная программа и почему?

- Deadlock. Нет указателя для `wg sync.WaitGroup`

```go
func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(wg sync.WaitGroup, i int) {
			fmt.Println(i)
			wg.Done()
		}(wg, i)
	}
	wg.Wait()
	fmt.Println("exit")
}
```

## 12 Что выведет данная программа и почему?

- 0. Переменная n в цикле записана через `:=` -> другая переменная

```go
func main() {
	n := 0
	if true {
		n := 1
		n++
	}
	fmt.Println(n)
}
```

## 13 Что выведет данная программа и почему?

- 100, 2, 3, 4, 5
- В функции `someAction` при добавлении нового элемента в массив `v` создается новый массив, так как исходный массив `a` не имеет достаточно места для добавления элемента. Это означает, что изменения, внесенные в `v`, не будут отражены на исходном массиве `a`.

```go
func someAction(v []int8, b int8) {
	v[0] = 100
	v = append(v, b)
}

func main() {
	var a = []int8{1, 2, 3, 4, 5}
	someAction(a, 6)
	fmt.Println(a)
}
```

## 14 Что выведет данная программа и почему?

- bba, aa
- Во время выполнения операции `append` в анонимной функции создается новый срез, который не связан с срезом в основной функции `main`. Поэтому изменения, внесенные в этот новый срез, не отразятся на исходном срезе в `main`.

```go
func main() {
	slice := []string{"a", "a"}
	func(slice []string) {
		slice = append(slice, "a")
		slice[0] = "b"
		slice[1] = "b"
		fmt.Print(slice)
	}(slice)
	fmt.Print(slice)
}
```