package main

import (
	"fmt"
)

// 定义一个人的结构体
type Person struct {
	id   int
	name string
	sex  byte
	age  int
}

// 定义一个学生的结构体，同时通过Person匿名字段，实现Student中可以继承Person中的属性
type Student struct {
	Person
	id   int // id字段与Person中的id字段同名
	addr string
}

func main() {

	// 定义并赋值Student变量
	var s1 Student
	s1.id = 123
	s1.name = "neil"
	s1.sex = 'm'
	s1.age = 26
	s1.addr = "wuhan"

	fmt.Printf("s1 = %+v\n", s1)
	// 结果为：
	// s1 = {Person:{id:0 name:neil sex:109 age:26} id:123 addr:wuhan}
	// 说明：对于同名字段如果直接赋值（即：s1.id），则会优先给该结构体的字段赋值

	// 如果要指定为匿名字段中的同名字段赋值，则进行如下操作
	s1.Person.id = 456
	fmt.Printf("s1 = %+v\n", s1)
	// 结果为：
	// s1 = {Person:{id:456 name:neil sex:109 age:26} id:123 addr:wuhan}
}
