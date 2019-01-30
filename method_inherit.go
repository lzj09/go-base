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

// 为Person结构体新增一个打印方法
func (p Person) printInfo() {
	fmt.Println("p = ", p)
}

// 定义一个Student结构体并继承Person
type Student struct {
	Person
	id   int
	addr string
}

func main() {
	// 定义并初始化Student结构体
	s := Student{Person: Person{name: "jane", sex: 'f', age: 18}, id: 123, addr: "wh"}

	s.printInfo()

	// 结果：
	// p =  {jane 102 18}

	// 结论：
	// 通过匿名字段可以继承其属性，也可以继承其的方法
}
