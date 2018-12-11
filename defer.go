package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("...1...")

	fmt.Println("...2...")

	defer fmt.Println("...3...")

	a := 10
	b := 20

	// 无参的匿名函数
	defer func() {
		fmt.Printf("func(): a = %d, b = %d\n", a, b)
	}()

	// 有参的匿名函数
	defer func(a, b int) {
		fmt.Printf("func(a, b): a = %d, b = %d\n", a, b)
	}(a, b)

	a = 100
	b = 200

	fmt.Printf("main: a = %d, b = %d\n", a, b)

	// 结果为：
	// ...2...
	// main: a = 100, b = 200
	// func(a, b): a = 10, b = 20
	// func(): a = 100, b = 200
	// ...3...
	// ...1...

	// 结论：
	// 1、用defer声明的部分，是在函数运行结束前运行
	// 2、对于多个defer，是按照“先进后出”的原则，即先写的defer语句后执行
	// 3、对于闭包所捕获的外部变量，是获取其最终的结果
	// 4、对于defer声明的传递参数型函数，它会先传递参数后运行
}
