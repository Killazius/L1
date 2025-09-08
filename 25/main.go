package main

import (
	"fmt"
	"time"
)

func Sleep(d time.Duration) {
	done := make(chan struct{})
	time.AfterFunc(d, func() {
		close(done)
	})
	<-done
}

// через канал и горутину + таймер
//
//	func Sleep(d time.Duration) {
//		done := make(chan struct{})
//		go func() {
//			timer := time.NewTimer(d)
//			<-timer.C
//			close(done)
//		}()
//		<-done
//	}
func main() {
	start := time.Now()
	Sleep(time.Second)
	fmt.Println(time.Since(start))
}
