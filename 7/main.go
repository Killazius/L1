package main

import (
	"fmt"
	"sync"
)

type Data struct {
	mu sync.RWMutex
	m  map[string]string
}

func NewData() *Data {
	return &Data{
		m: make(map[string]string),
	}
}

func (d *Data) Set(key, value string) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.m[key] = value
}
func main() {
	var wg sync.WaitGroup
	data := NewData()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", i)
			value := fmt.Sprintf("val%d", i)
			data.Set(key, value)
		}(i)
	}
	wg.Wait()
}
