package main

import (
	"fmt"
)

func reverse(data []byte, i, j int) {
	for i < j {
		data[i], data[j] = data[j], data[i]
		i++
		j--
	}
}

// если без массива data, то принимает массив байт изначально
func reverseWords(s string) string {
	data := []byte(s)
	reverse(data, 0, len(data)-1) // переворачиваем всю строку

	// переворачиваем каждое слово
	l := 0
	for r := 0; r <= len(data); r++ {
		if r == len(data) || data[r] == ' ' {
			reverse(data, l, r-1)
			l = r + 1
		}
	}

	return string(data)
}

func main() {
	fmt.Println(reverseWords("snow dog sun"))
}
