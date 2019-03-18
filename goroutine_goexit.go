package main

import (
	"fmt"
	"runtime"
)

func test() {
	defer fmt.Println("cccccccccc")
	// 执行该语句时，会使协程退出
	runtime.Goexit()
	fmt.Println("dddddddddddd")
}

// runtime的Goexit方法退出当前协程，但是defer语句会照常执行
func main() {

	go func() {
		fmt.Println("aaaaaaaaaaa")
		test()
		// 由于test()方法有runtime.Goexit()，使该协程退出，所以以下语句都不会执行
		fmt.Println("bbbbbbbbbbbb")
	}()

	// 为了使主协程不退出
	for {
	}

	// 运行结果为：
	// aaaaaaaaaaa
	// cccccccccc
}
