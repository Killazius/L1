package main

import "fmt"

func BinarySearch(arr []int, target int) int {
	l, r := 0, len(arr)-1
	for l <= r {
		mid := (l + r) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}
func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(BinarySearch(arr, 1))
	fmt.Println(BinarySearch(arr, 9))
	fmt.Println(BinarySearch(arr, 10))
}
