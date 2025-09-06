package main

import "fmt"

func unique(strs []string) []string {
	m := make(map[string]struct{})

	for _, s := range strs {
		m[s] = struct{}{}
	}
	res := make([]string, 0, len(m))
	for s := range m {
		res = append(res, s)
	}
	return res
}
func main() {
	strs := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(unique(strs))
}
