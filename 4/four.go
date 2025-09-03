package main

import (
	"context"
	"fmt"
	"math/rand"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(ctx context.Context, id int, wg *sync.WaitGroup) {
	sleepTime := time.Duration(rand.Intn(1000)) * time.Millisecond
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("worker %d done\n", id)
			return
		default:
			fmt.Printf("worker %d work,sleep for %s\n", id, sleepTime)
			time.Sleep(sleepTime)
		}
	}

}

// Решение с NotifyContext
func main() {
	workers := 5
	var wg sync.WaitGroup
	wg.Add(workers)
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT)
	defer stop()
	for i := 0; i < workers; i++ {
		go worker(ctx, i, &wg)
	}

	<-ctx.Done()

	wg.Wait()
	fmt.Println("done")
}

//func main() {
//	workers := 5
//	var wg sync.WaitGroup
//	wg.Add(workers)
//	ctx, cancel := context.WithCancel(context.Background())
//	for i := 0; i < workers; i++ {
//		go worker(ctx, i, &wg)
//	}
//	stop := make(chan os.Signal, 1)
//	signal.Notify(stop, syscall.SIGINT)
//	<-stop
//	cancel()
//	wg.Wait()
//	fmt.Println("done")
//}
