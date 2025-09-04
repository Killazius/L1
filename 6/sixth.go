package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	// exit with context cancel
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("1 gorutine stopped")
				return
			default:
				fmt.Println("1 gorutine working")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()
	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)

	// exit with close channel, сюда же можно сделать таймер в селект и по таймеру закрывать канал
	ch := make(chan struct{})
	go func() {
		for {
			select {
			case <-ch:
				fmt.Println("2 gorutine stopped")
				return
			default:
				fmt.Println("2 gorutine working")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()
	time.Sleep(2 * time.Second)
	close(ch)
	time.Sleep(1 * time.Second)

	// exit with Goexit
	go func() {
		fmt.Println("3 gorutine working")
		defer fmt.Println("3 gorutine stopped")
		runtime.Goexit()
	}()

	time.Sleep(2 * time.Second)

	// exit with condition
	stop := false
	go func() {
		for {
			if stop {
				fmt.Println("4 goroutine stopped")
				return
			}
			fmt.Println("4 goroutine working")
			time.Sleep(500 * time.Millisecond)
		}
	}()

	time.Sleep(2 * time.Second)
	stop = true
	time.Sleep(1 * time.Second)

}
