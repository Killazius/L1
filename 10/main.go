package main

import (
	"fmt"
	"math"
)

func main() {
	degrees := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	m := make(map[int][]float64)
	for _, deg := range degrees {
		var key int
		if deg >= 0 {
			key = int(math.Floor(deg/10.)) * 10
		} else {
			key = int(math.Ceil(deg/10.)) * 10
		}
		m[key] = append(m[key], deg)
	}
	for k, v := range m {
		fmt.Printf("%d: %v\n", k, v)
	}
}
