package main

import (
	"fmt"
)

func main() {
	a, b := 10, 20

	// 进行值传递
	swap01(a, b)
	fmt.Printf("main: a = %d, b = %d\n", a, b)

	// 进行地址传递
	swap02(&a, &b)
	fmt.Printf("main: a = %d, b = %d\n", a, b)

	// 结果为：
	// swap01: a = 20, b = 10
	// main: a = 10, b = 20
	// swap02: *a = 20, *b = 10
	// main: a = 20, b = 10

	// 注意：通过值传递的方式，不会改变原变量的值，但是通过地址传递会改原变量的值
}

// 交换两个变量的值，接收的int类型参数
func swap01(a, b int) {
	a, b = b, a
	fmt.Printf("swap01: a = %d, b = %d\n", a, b)
}

// 交换两个变量的值，接收的*int类型（地址类型）参数
func swap02(a, b *int) {
	*a, *b = *b, *a
	fmt.Printf("swap02: *a = %d, *b = %d\n", *a, *b)
}
