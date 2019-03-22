package main

import (
	"fmt"
	"time"
)

// 无缓存通道，即该通道中是不能存储值的，即：
// 如果读取无缓存通道时，此时没有值，则会阻塞
// 如果往无缓存通道写值，此时通道中有值，则同样会阻塞
func main() {

	// 定义一个无缓存通道
	ch := make(chan int) // 或者：make(chan int, 0)这样也是创建无缓存的通道

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Printf("go：开始写第%d次\n", (i + 1))
			// 往通道中写值
			ch <- i
			fmt.Println("go：i = ", i)
		}
	}()

	// 主协程sleep一会儿
	time.Sleep(3 * time.Second)

	for i := 0; i < 3; i++ {
		// 从通道中取值
		num := <-ch
		fmt.Println("main: num = ", num)
	}

	// 结果为：（可能运行每次的结果会不一样）
	// go：开始写第1次
	// main: num =  0
	// go：i =  0
	// go：开始写第2次
	// go：i =  1
	// go：开始写第3次
	// main: num =  1
	// main: num =  2

	// 多次运行后可以发现，如果不从通道中将数据读取，往通道写的操作会阻塞
}
