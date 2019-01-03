package main

import (
	"fmt"
)

func main() {
	// 定义并初始化切片
	slices := []int{1, 2, 3, 4, 5, 6, 7, 8}

	// 切片截取格式为：slices[low:high:max]
	// 表示：从切片slices中，从low开始位置截取，截取的长度为(hight - low)，容量为(max - low)
	// 如果max未填，容量为(cap(slices) - low)

	s1 := slices[2:4:8]
	fmt.Println("s1 = ", s1)
	fmt.Printf("s1: len = %d, cap = %d\n", len(s1), cap(s1))

	s2 := slices[2:5]
	fmt.Println("s2 = ", s2)
	fmt.Printf("s2: len = %d, cap = %d\n", len(s2), cap(s2))

	// s3 := slices[2:4:9]
	// fmt.Println("s3 = ", s1)
	// fmt.Printf("s3: len = %d, cap = %d\n", len(s3), cap(s3))
	// 注意：此时slices[2:4:9]会报slice bounds out of range错误，
	// 因为此cap为（9 - 2 = 7），超过了（cap(slices) - 2 = 6）
}
