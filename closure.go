package main

import (
	"fmt"
)

func main() {
	num := 3
	str := "hello go"

	// 在匿名函数中修改num及str的值
	func() {
		num++
		str = "anonymours..."

		fmt.Printf("anonymours：num = %d, str = %s\n", num, str)
	}()

	fmt.Printf("main: num = %d, str = %s\n", num, str)
	// 结果为：
	// anonymours：num = 4, str = anonymours...
	// main: num = 4, str = anonymours...

	// 结论：
	// 闭包是引用的方式捕获外部变量

	funcTest := test()
	fmt.Println(funcTest())
	fmt.Println(funcTest())
	fmt.Println(funcTest())
	// 结果为：
	// 1
	// 2
	// 3

	// 结论：
	// 只要闭包还在使用，那么它捕获的变量就在存在
}

// 定义一个函数，返回值为函数类型
func test() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
