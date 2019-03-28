package main

import (
	"fmt"
)

// 单向的channel，即要么只能写数据；要么只能读数据
func main() {
	// 定义一个channel，默认的channel是双向的，即可以写，也可以读
	ch := make(chan int)

	// chan<- int即是定义只能往channel写int数据
	go func(writeChan chan<- int) {
		for i := 1; i < 6; i++ {
			writeChan <- i
		}

		// 关闭channel
		close(writeChan)
	}(ch) // 注意：channel的传递是引用传递，同时双向channel是可以转换成单向channel

	// <-chan int即是定义只能读取channel中的int数据
	func(readChan <-chan int) {
		// 通过range迭代channel
		for num := range readChan {
			fmt.Println("num = ", num)
		}
	}(ch)

	// 结果为：
	//num =  1
	// num =  2
	// num =  3
	// num =  4
	// num =  5
}
