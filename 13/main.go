package main

import "fmt"

func main() {
	a, b := 1, 5
	a += b
	b = a - b
	a -= b
	fmt.Println(a, b)
}
