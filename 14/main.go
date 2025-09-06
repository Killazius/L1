package main

import "fmt"

func detect(v interface{}) string {
	switch v.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan int, chan string, chan bool:
		return "chan"
	default:
		return "unknown"
	}
}

func main() {
	fmt.Println(detect(42))             // int
	fmt.Println(detect("hello"))        // string
	fmt.Println(detect(true))           // bool
	fmt.Println(detect(make(chan int))) // chan
	fmt.Println(detect(3.14))           // unknown
}
