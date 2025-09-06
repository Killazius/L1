package main

import "fmt"

// корректно работает с двумя одинаковыми числами в каждом слайсе, если надо по одному то можно мапу с пустой структурой
func main() {
	A := []int{1, 2, 3, 2}
	B := []int{2, 3, 4, 2}
	m := make(map[int]int)
	for _, v := range A {
		m[v]++
	}
	var res []int
	for _, v := range B {
		if m[v] > 0 {
			res = append(res, v)
			m[v]--
		}
	}
	fmt.Println(res)
}
