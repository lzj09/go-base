package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建带缓存的channel
	ch := make(chan int, 3)
	// 显示channel的长度、容量
	fmt.Printf("len(ch) = %d, cap(ch) = %d\n", len(ch), cap(ch))

	go func() {
		for i := 0; i < 6; i++ {
			ch <- i
			fmt.Println("子协程 i = ", i)
		}

		// 关闭channel
		// 关闭channel后，是不能往channel里面写数据，但可以从channel里面读数据
		close(ch)
	}()

	// 让主协程sleep
	time.Sleep(3 * time.Second)

	for {
		// 从channel里读取数据，
		// 如果channel没有关闭，那么ok为true
		// 如果channel关闭，那么ok为false
		if data, ok := <-ch; ok {
			fmt.Println("data = ", data)
		} else {
			break
		}
	}

	// 结果为：（可能每次运行的结果略有不同）
	// len(ch) = 0, cap(ch) = 3
	// 子协程 i =  0
	// 子协程 i =  1
	// 子协程 i =  2
	// data =  0
	// data =  1
	// data =  2
	// data =  3
	// 子协程 i =  3
	// 子协程 i =  4
	// 子协程 i =  5
	// data =  4
	// data =  5

	// 通过运行结果可以发现，如果缓存的channel里面放置的数据满了时，如果此时没有读取数据，那么放置数据端会阻塞

}
