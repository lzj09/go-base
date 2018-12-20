package main

import (
	"fmt"
)

func main() {
	// 定义一个int型变量
	var num int

	// 定义一个*int的指针变量
	var p *int

	// &num表示取num的地址，并赋值给指针变量p
	p = &num

	fmt.Printf("num = %v, p = %v\n", num, p)

	// *p表示p指针变量所指向的地址内容，也即num的值
	*p = 20

	fmt.Printf("num = %v, p = %v\n", num, p)

	// 结果为：
	// num = 0, p = 0xc042048058
	// num = 20, p = 0xc042048058
}
