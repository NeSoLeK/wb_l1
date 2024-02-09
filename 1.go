package main

//Дана структура Human (с произвольным набором полей и методов).
//Реализовать встраивание методов в структуре Action от родительской структуры
//Human (аналог наследования)

//Структура Human

type Human struct {
	name  string
	age   int
	email string
}

// Метод NewHuman (Human)

func NewHuman(name string, age int, email string) *Human {
	return &Human{
		name:  name,
		age:   age,
		email: email,
	}
}

// Метод UpdateName (Human)

func (h *Human) UpdateName(name string) {
	h.name = name
}

//Метод UpdateEmail (Human)

func (h *Human) UpdateEmail(email string) {
	h.email = email
}

//Структура Action (Human)

type Action struct {
	*Human
}

// Метод NewAction (Action)

func NewAction(human *Human) *Action {
	return &Action{
		Human: human,
	}
}

// Метод NewAge (Action)

func (a *Action) NewAge(age int) {
	a.age = age
}

func t1() {
	//Создаем экземпляр Human
	user := NewHuman("Ilya", 18, "test@gmail.com")
	//Создаем экземпляр Action, хранящего user(Human)
	action := NewAction(user)
	//Вызываем метод NewAge(), унаследованный от Human
	action.NewAge(19)
}
