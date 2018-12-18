package main

import (
	// 如果需要用到不同目录的go方法，则需要导入相应包
	"calc"
	"fmt"
)

func init() {
	fmt.Println("main init...")
}

func main() {
	// 对于不同目录下的go方法，只能调用首字母为大写的方法
	a := calc.Add(10, 20)
	fmt.Println("a = ", a)

	b := calc.Minus(20, 10)
	fmt.Println("b = ", b)

	// 对于同目录下的go方法，可以直接调用
	test()

	// 结果为：
	// calc init...
	// main init...
	// a =  30
	// b =  10
	// test...

	// 注意：程序执行时，会先执行导入包的init方法，然后执行main包的init方法，再执行main方法

}
