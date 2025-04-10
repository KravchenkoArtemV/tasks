package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// необходимо получить слайс интов с только уникальными числами

func uniqRand(n int) []int {
	var res []int               // результирующий слайс
	m := make(map[int]struct{}) // мапа для сета

	for len(res) < n {
		temp := rand.Intn(100)
		if _, ok := m[temp]; !ok {
			res = append(res, temp)
		}
		m[temp] = struct{}{}
	}
	sort.Ints(res)
	return res
}

func main() {
	fmt.Println(uniqRand(10))
}
