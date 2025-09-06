package main

import "fmt"

func partition(arr []int, low, high int) int {
	pivot := arr[(low+high)/2]
	i := low
	j := high

	for i <= j {
		for arr[i] < pivot {
			i++
		}
		for arr[j] > pivot {
			j--
		}
		if i <= j {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		}
	}
	return i
}

func quicksort(arr []int, low, high int) {
	if low < high {
		index := partition(arr, low, high)
		quicksort(arr, low, index-1)
		quicksort(arr, index, high)
	}
}

func Quicksort(arr []int) {
	quicksort(arr, 0, len(arr)-1)
}
func main() {
	arr := []int{3, 6, 8, 10, 1, 2, 1}
	Quicksort(arr)
	fmt.Println(arr)
}
