package main

import (
	"fmt"
	"time"
)

func main() {

	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}

// request 1 2024-10-06 21:48:03.261431 +0900 JST m=+0.201281542
// request 2 2024-10-06 21:48:03.461391 +0900 JST m=+0.401248709
// request 3 2024-10-06 21:48:03.661373 +0900 JST m=+0.601238084
// request 4 2024-10-06 21:48:03.861351 +0900 JST m=+0.801222876
// request 5 2024-10-06 21:48:04.061321 +0900 JST m=+1.001200251
// request 1 2024-10-06 21:48:04.061411 +0900 JST m=+1.001289792
// request 2 2024-10-06 21:48:04.061415 +0900 JST m=+1.001294667
// request 3 2024-10-06 21:48:04.061418 +0900 JST m=+1.001297084
// request 4 2024-10-06 21:48:04.262464 +0900 JST m=+1.202350709
// request 5 2024-10-06 21:48:04.462488 +0900 JST m=+1.402381792
