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

// 定义一个学生的结构体，
type Student struct {
	*Person // 指针匿名字段
	id      int
	addr    string
}

func main() {
	// 对于含有指针匿名字段的赋值方法如下：

	// 方法一：对于指针变量，通过&进行赋值
	s1 := Student{Person: &Person{name: "neil", sex: 'm', age: 28}, id: 123, addr: "wuhan"}
	fmt.Println(s1.name, s1.sex, s1.age, s1.id, s1.addr)

	// 方法二：对于指针变量，通过new进行赋值
	var s2 Student
	s2.Person = new(Person)
	s2.name = "jane"
	s2.sex = 'f'
	s2.age = 18
	s2.id = 456
	s2.addr = "wh"
	fmt.Println(s2.name, s2.sex, s2.age, s2.id, s2.addr)

	// 结果为：
	// neil 109 28 123 wuhan
	// jane 102 18 456 wh
}
