package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{2, 4, 6, 8, 10}
	results := make([]int, len(arr))

	var wg sync.WaitGroup
	wg.Add(len(arr))

	for i := range arr {
		go func(idx int) {
			defer wg.Done()
			results[idx] = arr[idx] * arr[idx]
		}(i)
	}
	wg.Wait()

	for _, val := range results {
		fmt.Println(val)
	}
}
