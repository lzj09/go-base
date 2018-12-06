package main

import (
	"fmt"
)

// 定义了一个两参数都整型，同时返回值也是整型的函数类型CalculateFunc
type CalculateFunc func(int, int) int

// 带有函数类型的参数的函数为回调函数
func calculate(a, b int, calculateFunc CalculateFunc) (result int) {
	result = calculateFunc(a, b)
	return
}

// 加法
func add(a, b int) (result int) {
	result = a + b
	return
}

// 减法
func minus(a, b int) (result int) {
	result = a - b
	return
}

// 乘法
func multiply(a, b int) (result int) {
	result = a * b
	return
}

// 除法
func division(a, b int) (result int) {
	result = a / b
	return
}

func main() {
	// 通过回调函数实现了多态
	res1 := calculate(8, 2, add)
	fmt.Println("res1 = ", res1)

	res2 := calculate(8, 2, minus)
	fmt.Println("res2 = ", res2)

	res3 := calculate(8, 2, multiply)
	fmt.Println("res3 = ", res3)

	res4 := calculate(8, 2, division)
	fmt.Println("res4 = ", res4)

	// 结果为：
	// res1 =  10
	// res2 =  6
	// res3 =  16
	// res4 =  4
}
