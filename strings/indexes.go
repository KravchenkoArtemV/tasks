//На вход подаются слова и индекс буквы на одной строке.
//Нужно собрать по буквам новое слово, состоящее из букв по указанным индексам.

// 1) решение через руны и конкатенацию, 2) решение со стрингбилдером

package main

import "fmt"

func main() {
	var count int
	fmt.Scan(&count)

	result := ""
	//  2. var sb strings.Builder

	for i := 0; i < count; i++ {
		var word string
		var index int
		fmt.Scan(&word, &index)

		// 2. sb.WriteByte(word[index])
		r := []rune(word)
		result += string(r[index])
	}
	fmt.Println(result)
	// 2. fmt.Println(sb.String())
}
