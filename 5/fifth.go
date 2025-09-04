package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	seconds := flag.Int("s", 5, "seconds to run") // можно использовать flag.Duration, но у нас только секунды
	flag.Parse()
	dur := time.Duration(*seconds) * time.Second
	fmt.Printf("%s to end\n", dur)
	start := time.Now().Add(dur)
	ch := make(chan int, 1)
	go func() {
		defer close(ch)
		timeout := time.After(dur)
		for {
			select {
			case <-timeout:
				return
			default:
				ch <- rand.Intn(10)
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()
	for v := range ch {
		fmt.Printf("val recieved %d, time2end %s\n", v, time.Until(start).Round(time.Millisecond))
	}
	fmt.Println("done")
}
