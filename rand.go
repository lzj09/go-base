package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 设置随机种子
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 10; i++ {
		fmt.Println("rand.Int() = ", rand.Int())

		// 指定范围的随机数
		fmt.Println("rand.Intn(1000) = ", rand.Intn(1000))
	}
}
