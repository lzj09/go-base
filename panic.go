package main

import "fmt"

// 测试方法1
func test01() {
	fmt.Println("test01...")
}

// 测试方法2
func test02() {
	// 显示调用panic函数
	panic("test02方法：发生panic异常...")
}

// 测试方法3
func test03() {
	fmt.Println("test03...")
}

// 程序当中，有些异常是致命的异常，出现之后会导致程序的中止运行
// 在go语言当中，panic异常，就是会导致程序的中止运行
func main() {

	// 调用测试方法
	test01()
	test02()
	test03()

	// 结果为：
	// test01...
	// panic: test02方法：发生panic异常...

	// goroutine 1 [running]:
	// main.test02()
	//        D:/Program/go/workspace_1.9.2/src/panic.go:13 +0x40
	// main.main()
	//        D:/Program/go/workspace_1.9.2/src/panic.go:27 +0x2c
	// exit status 2

	// 可以发现当发生panic异常后，程序中止了运行

	// 当数组发生下标越界时，同样是会报panic异常，只不过是内部封装好了的panic异常
	arrs := [3]int{1, 2, 3}
	fmt.Println("arrs[3] = ", arrs[3])
	// 报如下错误
	// .\panic.go:45:32: invalid array index 3 (out of bounds for 3-element array)
}
