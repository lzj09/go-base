package main

import (
	"fmt"
)

func main() {
	// 定义匿名函数
	f1 := func() {
		fmt.Println("do anonymous func...")
	}
	// 匿名函数调用
	f1()

	// 匿名函数定义及调用
	func() {
		fmt.Println("hello go...")
	}()

	// 匿名函数定义及调用
	max, min := func(a, b int) (max, min int) {
		if a > b {
			max = a
			min = b
		} else {
			max = b
			min = a
		}

		return
	}(33, 66)
	fmt.Printf("max = %d, min = %d\n", max, min)

	// 结果为：
	// do anonymous func...
	// hello go...
	// max = 66, min = 33
}
