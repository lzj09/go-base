package main

import (
	"fmt"
	"regexp"
)

// 利用正则表达式提取邮箱的实例
func main() {
	// 定义一个字符串，里面含有所需的邮箱
	text := "北风凌冽，不再nckk.test@163.com散步。再后来，忙忙碌碌，这个书屋也被淡忘了。123@qq.com直到半年后的一天，晚饭后我独自散步，临出门时突然有要买一本书的冲动，便想起了那个温馨的小书屋。"

	// 定义一个匹配邮箱的正则规则
	reg := regexp.MustCompile(`[a-zA-Z0-9_.-]+@[a-zA-Z0-9]+(\.[a-zA-Z0-9-]+)*\.[a-zA-Z0-9]{2,6}`)

	// 利用正则，提取符合规则的字符串
	res := reg.FindAllStringSubmatch(text, -1)

	// 循环结果
	for _, data := range res {
		fmt.Println(data[0])
	}

	// 结果为：
	// nckk.test@163.com
	// 123@qq.com
}
