package main

import "fmt"

func foo(x int, k uint, v int) int {
	if v == 1 {
		return x | (1 << k)
	}
	return x &^ (1 << k)
}
func main() {
	fmt.Println(foo(5, 1, 0)) // 5 = 101 &^ 010 = 101
	fmt.Println(foo(5, 1, 1)) // 5 = 101 | 010 = 111
	fmt.Println(foo(4, 1, 0)) // 4 = 100 &^ 010 = 100
	fmt.Println(foo(4, 1, 1)) // 4 = 100 | 010 = 110
}
