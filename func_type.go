package main

import (
	"fmt"
)

// 两整型相加
func add(a, b int) (result int) {
	result = a + b
	return
}

// 两整型相减
func minus(a, b int) (result int) {
	result = a - b
	return
}

// 定义了一个两参数都整型，同时返回值也是整型的函数类型FuncTest
type FuncTest func(int, int) int

func main() {
	var funcTest FuncTest

	// 将加法函数赋值给funcTest
	funcTest = add
	res1 := funcTest(10, 10) // 等价于 res1 := add(10, 10)
	fmt.Println("res1 = ", res1)

	// 将减法函数赋值给funcTest
	funcTest = minus
	res2 := funcTest(10, 10) // 等价于 res2 := minus(10, 10)
	fmt.Println("res2 = ", res2)

	// 结果为：
	// res1 = 20
	// res2 = 0
}
