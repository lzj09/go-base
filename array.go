package main

import (
	"fmt"
)

func main() {
	// 数组定义及初始化
	// 给全部元素赋值
	arrs1 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("arrs1 = ", arrs1)

	// 给部分元素赋值，未赋值的默认为0
	arrs2 := [5]int{1, 2, 3}
	fmt.Println("arrs2 = ", arrs2)

	// 给指定元素赋值，未赋值的默认为0
	arrs3 := [5]int{1: 10, 4: 30}
	fmt.Println("arrs3 = ", arrs3)

	// 给每个元素遍历及赋值
	var arrs4 [5]int
	for i := 0; i < len(arrs4); i++ {
		arrs4[i] = i + 1
	}

	// 遍历数组并打印值
	for i, data := range arrs4 {
		fmt.Printf("arrs4[%d] = %d\n", i, data)
	}

	// 结果为：
	// arrs1 =  [1 2 3 4 5]
	// arrs2 =  [1 2 3 0 0]
	// arrs3 =  [0 10 0 0 30]
	// arrs4[0] = 1
	// arrs4[1] = 2
	// arrs4[2] = 3
	// arrs4[3] = 4
	// arrs4[4] = 5
}
