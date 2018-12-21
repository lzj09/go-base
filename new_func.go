package main

import (
	"fmt"
)

func main() {
	// a是int型变量
	a := 10
	// p是*int型变量，通过&符的方式
	p := &a

	fmt.Printf("a = %v, *p = %v\n", a, *p)

	// 通过new函数，定义一个*int型变量
	q := new(int)
	// 给q所指向的内存赋值为20
	*q = 20
	// 将*q的值赋值给b
	b := *q

	fmt.Printf("b = %v, *q = %v\n", b, *q)

	// 结果为：
	// a = 10, *p = 10
	// b = 20, *q = 20
}
