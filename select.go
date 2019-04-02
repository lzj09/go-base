package main

import (
	"fmt"
)

// select的用法有点类似switch，只是它的case条件是IO操作，
// 也即是判断channel是否可以读取值，是否可以设置值
func main() {

	// select实现斐波那契数列
	// 定义一个整型管道，用于传递数字
	ch := make(chan int)
	// 定义一个布尔型管道，用于传递是否结束
	flag := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			// 从管道中获取斐波那契数列值
			num := <-ch
			fmt.Println(num)
		}

		// 设置结束标志
		flag <- true
	}()

	// 调用
	fibonacci(ch, flag)

	// 结果为：
	// 1
	// 1
	// 2
	// 3
	// 5
	// 8
	// 13
	// 21
	// 34
	// 55
}

// 实现斐波那契数列
func fibonacci(ch chan<- int, flag <-chan bool) {
	// 斐波那契数列的规律为：
	// 1 1 2 3 5 8 13 21...
	// 即从第3开始的值，是从前2位之和

	// 定义两个初始化变量
	x, y := 1, 1
	for {
		select {
		case ch <- x:
			// 重新设置x和y的值
			x, y = y, x+y
		case <-flag:
			// 只当传递结束标志为true时，就结束
			return
		}
	}
}
