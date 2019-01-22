package main

import (
	"fmt"
)

// 定义一个人的结构体
type Person struct {
	name string
	sex  byte
	age  int
}

// 定义一个学生的结构体，同时通过Person匿名字段，实现Student中可以继承Person中的属性
type Student struct {
	Person
	id   int
	addr string
}

func main() {
	// 对含有匿名字段的结构体赋值方式
	// 方法一：顺序赋值
	s1 := Student{Person{"neil", 'm', 26}, 1, "wuhan"}
	// 通过%+v，可以打印出更加详细的信息
	fmt.Printf("s1 = %+v\n", s1)

	// 方法二：指定属性赋值，未赋值的属性则为该属性的默认值
	s2 := Student{Person: Person{name: "neil"}, addr: "wuhan"}
	fmt.Printf("s2 = %+v\n", s2)

	// 修改属性的方法
	s2.name = "chris"
	s2.sex = 'f'
	s2.age = 28
	s2.id = 2
	fmt.Printf("s2 = %+v\n", s2)

	// 结果为：
	// s1 = {Person:{name:neil sex:109 age:26} id:1 addr:wuhan}
	// s2 = {Person:{name:neil sex:0 age:0} id:0 addr:wuhan}
	// s2 = {Person:{name:chris sex:102 age:28} id:2 addr:wuhan}

	// 结论：
	// go语言当中是通过结构体的匿名字段来实现继承的
}
