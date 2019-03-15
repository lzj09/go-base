package main

import (
	"fmt"
	"time"
)

// 在Go语言进行开发时，一般会使用goroutine来处理并发任务
// goroutine是建立在线程之上的轻量级的抽象。它允许我们以非常低的代价在同一个地址空间中并行地执行多个函数或者方法。
// 相比于线程，它的创建和销毁的代价要小很多，并且它的调度是独立于线程的。
// 在golang中创建一个goroutine非常简单，使用“go”关键字即可

// 其中main()可以认为是一个主协程（主goroutine），其它通过go关键字开启的协程，称为子协程
// 注意: 当主协程退出时，所有的子协程也将退出
func main() {
	// 通过go关键字开启goroutine
	go func() {
		i := 0
		for {
			i++
			fmt.Println("子协程 === ", i)
			time.Sleep(time.Second)

		}
	}()

	for i := 0; i < 2; i++ {
		fmt.Println("主协程 === ", i)
		time.Sleep(time.Second)
	}

	// 结果为：(每一次运行的结果可能略有不同)
	// 主协程 ===  0
	// 子协程 ===  1
	// 子协程 ===  2
	// 主协程 ===  1

	// 注意：多运行几次，我们可以发现每当主协程出现两次后，退出出整个程序就退出了，也即证明了：
	// 当主协程退出时，所有的子协程也将退出
}
