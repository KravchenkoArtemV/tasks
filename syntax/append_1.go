package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	addNum(nums[0:2])
	fmt.Println(nums) // 1 2 4
	addNums(nums[0:2])
	fmt.Println(nums) // 1 2 4
}
func addNum(nums []int) {
	nums = append(nums, 4) // здесь замена посл элемента len 2 cap 3 - все ок
}
func addNums(nums []int) {
	nums = append(nums, 5, 6) // здесь не помещаемся по cap, пересоздаем базовый массив, поэтому в main нет изменений
}
