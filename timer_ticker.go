package main

import (
	"fmt"
	"time"
)

// timer定时器，时间到后就触发一次
func testTimer() {
	// 定义3秒后触发
	timer := time.NewTimer(3 * time.Second)
	// 停止
	defer timer.Stop()

	for i := 0; i < 5; i++ {
		<-timer.C
		fmt.Println("timer时间到...")

		// 可以通过Reset，重置定时触发时间
		timer.Reset(3 * time.Second)
	}
}

// ticker定时器，每隔固定时间都会触发
func testTicker() {
	// 定义每3秒触发
	ticker := time.NewTicker(3 * time.Second)
	// 停止
	defer ticker.Stop()

	for i := 0; i < 5; i++ {
		<-ticker.C
		fmt.Println("ticker时间到...")
	}
}

func main() {
	go testTimer()
	go testTicker()

	for {
	}
}
