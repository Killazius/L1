package main

import "fmt"

// еще есть вариант через append
func deleteElement(s []int, index int) []int {
	if index < 0 || index >= len(s) {
		return s
	}
	copy(s[index:], s[index+1:])
	return s[:len(s)-1]
}
func main() {
	s := []int{1, 2, 3, 4, 5}
	s = deleteElement(s, 3)
	fmt.Println(s) // [1 2 3 5]

}
