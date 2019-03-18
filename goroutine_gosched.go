package main

import (
	"fmt"
	"runtime"
)

// runtime包中的Gosched方法是让当前协程让出cpu的执行时间
// 让其它协程执行，等其它协程执行完毕后，再执行当前协程
func main() {

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("goroutine: = ", i)
		}
	}()

	for i := 0; i < 2; i++ {
		// 如果不加runtime.Gosched()，很有可能上面这个协程就执行不到，因为
		// 主协程退出后，其它协程就会退出，加了runtime.Gosched()之后会使主协程
		// 让出，让上面的协程先执行。
		runtime.Gosched()

		fmt.Println("main: = ", i)

		// 结果为：
		// goroutine: =  0
		// goroutine: =  1
		// goroutine: =  2
		// goroutine: =  3
		// goroutine: =  4
		// main: =  0
		// main: =  1
	}
}
