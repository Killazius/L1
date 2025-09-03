package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d received job %d\n", id, job)
		time.Sleep(300 * time.Millisecond)
	}

}

func main() {
	workers := flag.Int("workers", 5, "Number of concurrent workers")
	flag.Parse()
	if *workers < 1 {
		panic("incorrent num of workers")
	}
	fmt.Printf("Starting workers %d\n", *workers)
	jobs := make(chan int, *workers)
	var wg sync.WaitGroup

	for i := 0; i < *workers; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	go func() {
		for i := 0; i < *workers*5; i++ {
			jobs <- i
		}
		close(jobs)
	}()
	wg.Wait()
	fmt.Println("All workers finished")
}
