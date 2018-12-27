package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 随机生成一个有10个元素的一维数组，然后对其进行排序
	var arrs [10]int

	len := len(arrs)

	// 设置随机种子
	rand.Seed(time.Now().UnixNano())

	// 对数组元素进随机赋值
	for i := 0; i < len; i++ {
		// 100以内的随机数
		arrs[i] = rand.Intn(100)
	}
	fmt.Println("排序前arrs: ", arrs)

	// 通过冒泡排序方式，进行升序排序
	for i := 0; i < len-1; i++ {
		for j := 0; j < len-1-i; j++ {
			if arrs[j] > arrs[j+1] {
				arrs[j], arrs[j+1] = arrs[j+1], arrs[j]
			}
		}
	}
	fmt.Println("排序后arrs: ", arrs)
}
