package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	count int
	mu    sync.Mutex
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func main() {
	var wg sync.WaitGroup
	counter := Counter{}

	for i := 0; i < 5; i++ {
		wg.Go(func() {
			for k := 0; k < 1000; k++ {
				counter.Increment()
			}
		})
	}
	wg.Wait()

	fmt.Println(counter.count)
}
