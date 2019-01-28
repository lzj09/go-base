package main

import (
	"fmt"
)

// 定义方法格式：
// func (typeVar Type) methodName(paramName ParamType) (result ResultType) {
//     方法体
// }
// 定义方法的格式与定义函数很相似，只是多了需要指定绑定的类型（这种类型即是用Type声明的类型）

type MyInt int

// 在MyInt类型绑定了add方法
func (num1 MyInt) add(num2 MyInt) (res MyInt) {
	res = num1 + num2
	return
}

func main() {
	var num1 MyInt = 10

	res := num1.add(20)
	fmt.Println("res = ", res)

	// 结果为：
	// res =  30
}
