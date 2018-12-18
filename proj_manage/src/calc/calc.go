package calc

import (
	"fmt"
)

func init() {
	fmt.Println("calc init...")
}

// 加法运算
func Add(a, b int) (result int) {
	result = a + b
	return
}

// 减法运算
func Minus(a, b int) (result int) {
	result = a - b
	return
}

// 乘法运算
func multiply(a, b int) (result int) {
	result = a * b
	return
}
