package main

import (
	"fmt"
)

// 通过interface关键字定义一个接口
type IUser interface {
	// 里面只有方法声明，没有方法体
	say()
}

// 定义学生结构体
type Student struct {
	id   int
	name string
}

// 给学生结构体绑定say方法
func (s *Student) say() {
	fmt.Printf("Student say: id = %d, name = %s\n", s.id, s.name)
}

// 定义教师结构体
type Teacher struct {
	id     int
	name   string
	course string
}

// 给教师结构体绑定say方法
func (t *Teacher) say() {
	fmt.Printf("Teacher say: id = %d, name = %s, course = %s\n", t.id, t.name, t.course)
}

// 定义一个接收IUser为参数的函数，可以实现多态
func testSay(i IUser) {
	i.say()
}

func main() {
	s := &Student{id: 123, name: "neil"}
	t := &Teacher{id: 456, name: "jane", course: "english"}

	testSay(s)

	testSay(t)

	// 结果为：
	// Student say: id = 123, name = neil
	// Teacher say: id = 456, name = jane, course = english
}
