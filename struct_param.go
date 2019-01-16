package main

import (
	"fmt"
)

// 定义学生的结构体
type Student struct {
	id   int
	name string
	sex  byte
	addr string
}

func test01(s Student) {
	// 修改id值
	s.id = 123
	fmt.Println("test01: s = ", s)
}

func test02(s *Student) {
	// 修改id值
	s.id = 123
	fmt.Println("test02: s = ", *s)
}

func main() {
	// 定义并赋值Student结构体
	s := Student{id: 101, name: "neil", sex: 'm', addr: "wuhan"}

	// 以Student结构体作为实参传递
	test01(s)
	fmt.Println("main: s = ", s)

	// 以Student结构体地址作为实参传递
	test02(&s)
	fmt.Println("main: s = ", s)

	// 结果为：
	// test01: s =  {123 neil 109 wuhan}
	// main: s =  {101 neil 109 wuhan}
	// test02: s =  {123 neil 109 wuhan}
	// main: s =  {123 neil 109 wuhan}

	// 结论：
	// 1、当直接以结构体作为实参传递时，是以值为传递的
	// 2、当直接以结构体地址作为实参传递时，是以指针为传递的
}
