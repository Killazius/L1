package main

import (
	"fmt"
	"strings"
)

func Unique(s string) bool {
	m := make(map[rune]struct{}, len(s))
	for _, r := range strings.ToLower(s) {
		if _, ok := m[r]; ok {
			return false
		}
		m[r] = struct{}{}
	}
	return true
}
func main() {
	fmt.Println(Unique("abcd"))
	fmt.Println(Unique("abCdefAaf"))
	fmt.Println(Unique("aabcd"))
	fmt.Println(Unique("AbCdefaf"))
}
