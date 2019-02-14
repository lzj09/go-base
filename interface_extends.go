package main

import (
	"fmt"
)

// 定义一个用户接口
type IUser interface {
	// 定义say方法
	say()
}

// 定义一个人接口
type IPersoner interface {
	// 通过匿名字段方式继承IUser
	IUser

	// 定义run方法
	run()
}

// 定义学生结构体
type Student struct {
	id   int
	name string
}

// 给学生结构体绑定say方法
func (s *Student) say() {
	fmt.Printf("Student %s say...\n", s.name)
}

// 给学生结构体绑定run方法
func (s *Student) run() {
	fmt.Printf("Student %s run...\n", s.name)
}

func main() {
	// 定义IPersoner类型变量
	var personer IPersoner

	// 由于Student是绑定了IPersoner的方法（包括继承的方法）
	personer = &Student{id: 123, name: "neil"}
	personer.say()
	personer.run()

	// 由于IPersoner接口继承IUser接口，则IPersoner类型可以转换成IUser类型
	var user IUser
	user = personer
	user.say()

	// 结果为：
	// Student neil say...
	// Student neil run...
	// Student neil say...
}
