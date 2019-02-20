package main

import "fmt"

// 测试函数1
func test01() {
	fmt.Println("test01...")
}

// 测试函数2
func test02(num int) {
	// 通过defer关键字，使用recover函数
	// defer定义的函数，是在函数运行后执行的
	defer func() {
		// 如果recover()的返回值不为nil，则说明该函数发生了错误
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	arrs := [3]int{1, 2, 3}
	fmt.Println("test02: ", arrs[num])
}

// 测试函数3
func test03() {
	fmt.Println("test03...")
}

// 由于panic发生了，会导致程序中断，如果使用recover函数可以把错误信息显示出来，然后程序继续运行
func main() {

	// 调用函数
	test01()
	// 由于arrs[3]，是下标越界了，如果不用recover处理，则程序会中断
	test02(3)
	test03()

	// 结果为：
	// test01...
	// runtime error: index out of range
	// test03...

	// 注意：虽然test02发生了错误，但是程序还是继续在运行
}
