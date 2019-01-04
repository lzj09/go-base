package main

import (
	"fmt"
)

func main() {
	s1 := []int{1}
	fmt.Println("s1 = ", s1)
	fmt.Printf("s1: len = %d, cap = %d\n", len(s1), cap(s1))

	// append函数，是向切片末尾追加元素，并且返回一个新的切片
	s1 = append(s1, 2)
	fmt.Println("s1 = ", s1)
	fmt.Printf("s1: len = %d, cap = %d\n", len(s1), cap(s1))

	s1 = append(s1, 3)
	fmt.Println("s1 = ", s1)
	fmt.Printf("s1: len = %d, cap = %d\n", len(s1), cap(s1))

	// 结果为：
	// s1 =  [1]
	// s1: len = 1, cap = 1
	// s1 =  [1 2]
	// s1: len = 2, cap = 2
	// s1 =  [1 2 3]
	// s1: len = 3, cap = 4

	// 结论：
	// append是往切入末尾追加元素，并且返回一个新的切片，同时如果超出了容量大小，容量会翻倍增加

	s2 := []int{1, 2}
	s3 := []int{3, 4, 5, 6}
	// copy函数，是将s2内容复制到s3中相应的位置
	copy(s3, s2)
	fmt.Println("s3 = ", s3)

	// 结果为：
	// s3 =  [1 2 5 6]
}
