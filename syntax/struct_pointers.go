package main

import "fmt"

type Person struct {
	name string
	age  uint8
}

func changePerson(person *Person) {
	person = &Person{ // решение 1 - поправить на person.name = "Vladimir"
		name: "Vladimir",
		age:  25, // решение 1 - поправить на person.age = 25
	}
}

func main() {
	person := &Person{
		name: "Ivan",
		age:  30,
	}
	fmt.Println(person.name)
	changePerson(person)
	fmt.Println(person.name)
}

/*
ОТВЕТ
Ivan
Ivan

В changePerson передаётся указатель *Person,
но внутри функции переменной person присваивается новый указатель &Person{...}.
Это означает, что person теперь указывает на другой объект,
но это изменение не влияет на переданный в main указатель.
*/

//==================================================

/*
РЕШЕНИЕ 1

func changePerson(person *Person) {
	person.name = "Vladimir"
	person.age = 25
}
*/
//==================================================
/*
РЕШЕНИЕ 2
Если нужно заменить объект полностью, можно изменить сигнатуру функции и вернуть новый объект

func changePerson(person *Person) *Person {
    return &Person{
        name: "Vladimir",
        age: 25,
    }
}

func main() {
    person := &Person{
        name: "Ivan",
        age: 30,
    }
    fmt.Println(person.name) // Ivan
    person = changePerson(person)
    fmt.Println(person.name) // Vladimir
}
*/
