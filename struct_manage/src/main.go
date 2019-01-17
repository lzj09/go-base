package main

import (
	"fmt"
	// 导入自定义的test包
	"test"
)

func main() {
	// 定义结构体变量
	// var s1 test.student
	// 此处会报cannot refer to unexported name test.student
	// 由于test包中student结构体首字母是小写，其它包无法访问
	// fmt.Println("s1 = ", s1)

	var s2 test.Student
	fmt.Println("s2 = ", s2)

	// 修改id值
	// s2.id = 123
	// 此处会报s2.id undefined (cannot refer to unexported field or method id)
	// id在Student结构体中是首字母是小写，其它包无法访问

	// 修改Name值
	s2.Name = "neil"

	fmt.Println("s2 = ", s2)

	// 结果为：
	// s2 =  {0 }
	// s2 =  {0 neil}
}
