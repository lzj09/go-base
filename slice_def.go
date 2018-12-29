package main

import (
	"fmt"
)

func main() {
	// 数组定义及初始化
	arrs := [5]int{1, 2, 3, 4, 5}
	fmt.Println("arrs = ", arrs)
	fmt.Printf("len(arrs) = %d, cap(arrs) = %d\n", len(arrs), cap(arrs))

	// 切片的定义及初始化方式一：
	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Println("slice1 = ", slice1)
	fmt.Printf("len(slice1) = %d, cap(slice1) = %d\n", len(slice1), cap(slice1))

	// 切片的定义及初始化方式二：make(类型，长度，容量)
	slice2 := make([]int, 5, 10)
	fmt.Println("slice2 = ", slice2)
	fmt.Printf("len(slice2) = %d, cap(slice2) = %d\n", len(slice2), cap(slice2))

	// 结果为：
	// arrs =  [1 2 3 4 5]
	// len(arrs) = 5, cap(arrs) = 5
	// slice1 =  [1 2 3 4 5]
	// len(slice1) = 5, cap(slice1) = 5
	// slice2 =  [0 0 0 0 0]
	// len(slice2) = 5, cap(slice2) = 10

	// 切片与数组的区别:
	// 数组定义时，必须指定数组的长度，而切片不需要
}
