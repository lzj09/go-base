package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 随机初始化切片的值
func randData(s []int) {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < len(s); i++ {
		s[i] = rand.Intn(100)
	}
}

// 排序
func sortData(s []int) {
	len := len(s)
	for i := 0; i < len; i++ {
		for j := 0; j < len-i-1; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
}

func main() {
	// 创建一个切片
	s := make([]int, 5)

	// 随机初始化值
	randData(s)
	fmt.Println("排序前：s = ", s)

	// 排序
	sortData(s)
	fmt.Println("排序后：s = ", s)

	// 结果为：
	// 排序前：s =  [59 40 83 78 49]
	// 排序后：s =  [40 49 59 78 83]

	// 结论：
	// 切片作函数参数，是直接是地址传递
}
