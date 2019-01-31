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
func (p *Person) printInfo() {
	fmt.Println("Person printInfo: ", p)
}

// 定义一个Student结构体并继承Person
type Student struct {
	Person
	id   int
	addr string
}

// 为Student结构体新增一个与Person结构体同名的方法
func (s *Student) printInfo() {
	fmt.Println("Student printInfo: ", s)
}

func main() {
	// 定义并初始化Student结构体
	s := Student{Person: Person{name: "jane", sex: 'f', age: 18}, id: 123, addr: "wh"}

	// 如果直接调用printInfo方法，会调用Student的printInfo方法
	s.printInfo() // 注意：上面printInfo方法是绑定的*Student，但是直接用s调用也是可以的，因为内部会转换成(&s).printInfo()

	// 如果想调用Person的printInfo方法，则此时必须显式调用
	s.Person.printInfo()

	// 结果为：
	// Student printInfo:  &{{jane 102 18} 123 wh}
	// Person printInfo:  &{jane 102 18}
}
