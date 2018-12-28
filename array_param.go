package main

import (
	"fmt"
)

func main() {
	arrs := [5]int{1, 2, 3, 4, 5}

	// 调用modify01
	modify01(arrs)
	fmt.Println("main: arrs = ", arrs)

	// 调用modify02
	modify02(&arrs)
	fmt.Println("main: arrs = ", arrs)

	// 结果为：
	// modify01: arrs =  [2 3 4 5 6]
	// main: arrs =  [1 2 3 4 5]
	// modify02: arrs =  [2 3 4 5 6]
	// main: arrs =  [2 3 4 5 6]

	// 结论：
	// 对于数组当作参数，是值传递
	// 对于数组以地址当作参数，才是指针传递
}

// 数组作参数
func modify01(arrs [5]int) {
	// 遍历元数，并对每个元素加1
	for i, data := range arrs {
		arrs[i] = data + 1
	}
	fmt.Println("modify01: arrs = ", arrs)
}

// 数组地址作参数
func modify02(p *[5]int) {
	// 遍历元数，并对每个元素加1
	for i, data := range *p {
		(*p)[i] = data + 1
	}
	fmt.Println("modify02: arrs = ", *p)
}
