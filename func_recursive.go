package main

import (
	"fmt"
)

func test(num int) {

	if num == 0 {
		// 终止递归的条件
		return
	}

	test(num - 1)
	fmt.Println("num = ", num)
}

func main() {
	// 函数递归调用
	test(5)

	// 结果为：
	// num =  1
	// num =  2
	// num =  3
	// num =  4
	// num =  5
}
