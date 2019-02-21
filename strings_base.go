package main

import (
	"fmt"
	"strings"
)

// 字符串当的一些基于函数用法
func main() {
	// Contains，用于判断字符串是否包含某字符串，如果包含返回true，否则返回false
	s1 := "hello golang"
	fmt.Println(strings.Contains(s1, "go"))
	fmt.Println(strings.Contains(s1, "world"))

	// Join，将切片中每一个元素，用指定的字符串连接起来
	arrs1 := []string{"wh", "bj", "gz"}
	fmt.Println(strings.Join(arrs1, "-"))

	// Index，获取字符串中某字符串的位置，如果没有，则返回-1
	fmt.Println(strings.Index(s1, "go"))
	fmt.Println(strings.Index(s1, "gotest"))

	// Repeat，指定字符串，重复的次数
	fmt.Println(strings.Repeat("golang", 2))

	// Split，将字符串根据指定字符串进行分割，返回一个切片
	s2 := "wh,bj,gz"
	arrs2 := strings.Split(s2, ",")
	for i, data := range arrs2 {
		fmt.Printf("arrs[%d] = %s\n", i, data)
	}

	// Trim，指定去除字符串两端的字符串
	s3 := "  golang     "
	fmt.Printf("+%s+\n", strings.Trim(s3, " "))

	// Fields，去除两端的空格，并以空格分割，返回一切片
	s4 := "   hello world  golang   "
	arrs3 := strings.Fields(s4)
	for i, data := range arrs3 {
		fmt.Printf("arrs[%d] = %s\n", i, data)
	}

	// 结果为：
	// true
	// false
	// wh-bj-gz
	// 6
	// -1
	// golanggolang
	// arrs[0] = wh
	// arrs[1] = bj
	// arrs[2] = gz
	// +golang+
	// arrs[0] = hello
	// arrs[1] = world
	// arrs[2] = golang
}
