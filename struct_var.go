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

func main() {
	// 结构体成员变量赋值方式
	// 方法一：定义一个Student结构体变量，然后通过点（.）运算符进行成员变量赋值
	var s1 Student
	s1.id = 101
	s1.name = "neil"
	s1.sex = 'm'
	s1.addr = "wuhan"
	fmt.Println("s1 = ", s1)

	// 方法二：定义Student结构体指针，然后通过点（.）运算符进行成员变量赋值
	// 通过new关键字，得到Student结构体指针s2
	s2 := new(Student)
	s2.id = 101
	s2.name = "neil"
	s2.sex = 'm'
	s2.addr = "wuhan" // (*s2).addr = "wuhan"，这种方式赋值也是可以的，但是麻烦一些
	fmt.Println("s2 = ", *s2)

	// 结果为：
	// s1 =  {101 neil 109 wuhan}
	// s2 =  {101 neil 109 wuhan}
}
