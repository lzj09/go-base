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

// 为Persson结构体新增设置成员变量值的方法
func (p Person) setValue1(name string, sex byte, age int) {
	p.name = name
	p.sex = sex
	p.age = age
}

// 为Persson结构体指针新增设置成员变量值的方法
func (p *Person) setValue2(name string, sex byte, age int) {
	p.name = name
	p.sex = sex
	p.age = age
}

func main() {
	// 初始化并赋值Person
	p := Person{name: "neil", sex: 'm', age: 26}
	p.printInfo()

	p.setValue1("jane", 'f', 18)
	p.printInfo()

	(&p).setValue2("chris", 'm', 29)
	p.printInfo()

	// 结果为：
	// p =  {neil 109 26}
	// p =  {neil 109 26}
	// p =  {chris 109 29}

	// 结论：
	// 当结构体作为方法接收者时，是值传递
	// 当结构体指针作为方法的接收者时，是引用传递
}
