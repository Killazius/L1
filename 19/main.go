package main

import "fmt"

func reverse(s string) string {
	sRune := []rune(s)
	for i, j := 0, len(sRune)-1; i < j; i, j = i+1, j-1 {
		sRune[i], sRune[j] = sRune[j], sRune[i]
	}
	return string(sRune)
}

func main() {
	fmt.Println(reverse("hello world"))
	fmt.Println(reverse("嗨。"))
	fmt.Println(reverse("главрыба"))
}
