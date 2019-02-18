package main

// 异常接口，用于程序当中的异常信息的封装
// 使用时，需要导入errors包
import (
	"errors"
	"fmt"
)

// 定义一个除法函数，用于获取两个整合的商
// 如果分母为0则会导致程序出错，所以在程序处理时，需要处理这种异常
// 则定义的函数有两个返回值，第1个返回值为其结果，第2个返回值为异常信息
func divide(a, b int) (res int, err error) {
	err = nil

	if b == 0 {
		// 处理异常
		// 通过源码信息可知，调用errors.New函数，可以返回异常接口
		err = errors.New("分母不能为0")
	} else {
		res = a / b
	}

	return
}

func main() {

	// 测试
	res1, err1 := divide(22, 2)
	if err1 != nil {
		// 说明有异常，则打印异常信息
		fmt.Println("err1: ", err1)
	} else {
		// 打印运算结果
		fmt.Println("res1 = ", res1)
	}

	res2, err2 := divide(22, 0)
	if err2 != nil {
		// 说明有异常，则打印异常信息
		fmt.Println("err2: ", err2)
	} else {
		// 打印运算结果
		fmt.Println("res2 = ", res2)
	}

	// 结果为：
	// res1 =  11
	// err2:  分母不能为0
}
