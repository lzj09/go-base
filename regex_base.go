package main

import (
	"fmt"
	// 需要用到正则相当的操作，需要导入regexp包
	"regexp"
)

func main() {
	// 正则表达式在go语句的用法
	// . 在正则中表示任意字符，那么以其为例子，说明正则的用法

	// 定义一个字符串，用于适配的正则表达式用的
	str := "nexil ddneil ne6il"

	// 定义一个正则规则，该方法的返回值为*Regexp
	reg := regexp.MustCompile("ne.il")

	// 利用正则，提取符合规则的字符串
	// 第2个参数，是用于提取符合规则的字符串个数，如果设置为-1，则表示提取所有
	// 该方法的返回值是2维切片
	res := reg.FindAllStringSubmatch(str, -1)

	// 直接输出结果
	fmt.Println("res = ", res)

	// 结果为：
	// res =  [[nexil] [ne6il]]
}
