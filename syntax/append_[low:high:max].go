package main

import "fmt"

func main() {
	// создаём исходный слайс — backing array: [10 20 30 40 50]
	data := []int{10, 20, 30, 40, 50}

	// создаём слайс part от data:
	// от индекса 1 (включительно) до 3 (не включительно)
	// с максимальной capacity до индекса 3 (исключительно)
	// len(part) = 2 (элементы: 20, 30), cap(part) = 3 - 1 = 2
	part := data[1:3:3] // [low:high:max]

	// передаём слайс в функцию — он указывает на те же данные,
	// но с ограниченной capacity
	modify(part)

	// выводим оригинальный data — не должен измениться
	// потому что append в modify создаст новый backing array
	fmt.Println("data:", data)
}

func modify(nums []int) {
	// nums = [20 30], cap = 2, len = 2
	// append(nums, 99) превышает capacity,
	// поэтому создаётся новый массив (backing array)
	nums = append(nums, 99)

	// nums теперь указывает на НОВЫЙ массив: [20 30 99]
	// это НЕ влияет на оригинальный data
	fmt.Println("inside:", nums)
}
