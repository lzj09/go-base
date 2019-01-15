package main

import (
	"fmt"
)

// 定义学生的结构体
type Student struct {
	id   int
	name string
	age  byte
	addr string
}

func main() {
	// 方法一：顺序初始化，每个元素都得初始化
	s1 := Student{101, "neil", 'm', "wuhan"}
	fmt.Println("s1 = ", s1)

	// 方法二：部分元素初始化，未初始化的值是该类型的默认值
	s2 := Student{id: 101, addr: "wuhan"}
	fmt.Println("s2 = ", s2)

	// 方法三：指针方法初始化，未初始化的值是该类型的默认值
	s3 := &Student{id: 101, addr: "wuhan"}
	fmt.Println("*s3 = ", *s3)

	// 结果为：
	// s1 =  {101 neil 109 wuhan}
	// s2 =  {101  0 wuhan}
	// *s3 =  {101  0 wuhan}
}
