package main

import (
	"fmt"
)

func test1(num int) {
	test2(num - 1)
	fmt.Println("test1: num = ", num)
}

func test2(num int) {
	test3(num - 1)
	fmt.Println("test2: num = ", num)
}

func test3(num int) {
	fmt.Println("test3: num = ", num)
}

func main() {
	// 函数嵌套调用
	test1(100)

	// 结果为：
	// test3: num =  98
	// test2: num =  99
	// test1: num =  100
}
