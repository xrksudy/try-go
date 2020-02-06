package main

import (
	"fmt"
	"time"
)

func main() {
	// 每隔200毫秒处理请求
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(200 * time.Microsecond)

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// 先缓存了3个，同时处理请求，后面每隔200毫秒进行处理
	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Microsecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequest := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequest <- i
	}
	close(burstyRequest)

	for req := range burstyRequest {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}

}
