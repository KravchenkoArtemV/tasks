package main

import "fmt"

type person struct {
	age int
}

func main() {
	people := make([]person, 2) // len=2, cap=2, массив: [{0}, {0}]

	p1 := &people[1] // p1 указывает на people[1]
	fmt.Println(p1)  // Выведет адрес: &{0}

	p1.age++ // people[1].age = 1

	people = append(people, person{}, person{}, person{}) // cap=2 → создаётся новый массив cap=6

	// РЕШЕНИЕ После append() нужно обновить p1: p1 = &people[1] Теперь p1 указывает на новый массив

	fmt.Println(cap(people)) // Выведет: 6 (новая ёмкость)

	p1.age++ // p1 все ещё указывает на старый массив → p1.age = 2

	fmt.Println(people[1].age) // Выведет: 1 (в новом массиве)
	fmt.Println(p1.age)        // Выведет: 2 (в старом массиве)
}

/* ОТВЕТ
&{0}   Адрес people[1]
6      Новая ёмкость (cap) после append
1      people[1].age (в новом массиве)
2      p1.age (в старом массиве)
*/
