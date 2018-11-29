package main

import (
	"fmt"
)

// 无参无返回值函数
func test01() {
	fmt.Println("test01...")
}

// 无参有一个返回值函数
func test02() (result int) {
	result = 123
	return
}

// 无参有多个返回值函数
func test03() (num int, str string) {
	num = 99
	str = "Hello go"
	return
}

// 有参有返回值函数
func test04(a int) (num int, str string) {
	num = a + 1
	str = "test"
	return
}

func main() {
	test01()

	a := test02()
	fmt.Println("a = ", a)

	b, c := test03()
	fmt.Println("b = ", b, ", c = ", c)

	e, f := test04(100)
	fmt.Println("e = ", e, ", f = ", f)
}
