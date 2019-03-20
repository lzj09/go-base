package main

import (
	"fmt"
	"time"
)

// 将字符串的每个字符打印
func test01(str string) {
	for _, item := range str {
		fmt.Printf("%c", item)
		time.Sleep(time.Second)
	}
	// 换行
	fmt.Println()
}

// 当利用多协程访问公共资源时，会出现资源竞争的现象，出现不正确的结果
// 在此通过channel来实现同步
func main() {

	//go test01("neil")
	//go test01("jane")
	// 结果为：(每次的结果可能有所不同，由于没有同步，所以出现我们不想要的结果)
	//njeainle

	// 通过make创建channel
	ch := make(chan int)

	go func() {
		test01("neil")
		// 往channel中放入值
		ch <- 1
	}()

	go func() {
		// 如果ch中没有值，此处会阻塞，通过此就可以实现同步
		<-ch
		test01("jane")
	}()

	// 结果为：
	// neil
	// jane

	// 防止主协程退出
	for {

	}
}
