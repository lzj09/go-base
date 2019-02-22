package main

import (
	"fmt"

	// 字符串相关处理的相关包
	"strconv"
)

// 字符串基本转换
func main() {

	// Append操作，将Bool、Float、Int、Quote等类型的内容添加到字节切片里面
	arrs := make([]byte, 0, 32)
	// 添加Bool
	arrs = strconv.AppendBool(arrs, true)

	// 添加Float
	// AppendFloat(dst []byte, f float64, fmt byte, prec, bitSize int)
	// 其中fmt参数，有效数字用于小数点之后的位数，可以有 'e', 'E' 和 'f'
	// prec表示有效数字，-1，表示越精确越好
	// bitSize的值可以是32、64分别表float32、float64
	arrs = strconv.AppendFloat(arrs, 2.32, 'f', -1, 32)

	// 添加Int
	// 最后一个参数，表示以10进制
	arrs = strconv.AppendInt(arrs, 14, 10)

	// 添加Quote
	arrs = strconv.AppendQuote(arrs, "hello")

	// 可以先将arrs转换成string，然后打印结果查看
	fmt.Println("arrs = ", string(arrs))
	// 结果为：
	// arrs =  true2.3214"hello"

	// 字符串转换成其它类型
	// 转换成Bool
	// 转换失败err1会有异常信息
	num1, err1 := strconv.ParseBool("true")
	if err1 == nil {
		fmt.Println("num1 = ", num1)
	} else {
		fmt.Println("err1 = ", err1)
	}

	// 转换Float
	num2, err2 := strconv.ParseFloat("1.23", 64)
	if err2 == nil {
		fmt.Println("num2 = ", num2)
	} else {
		fmt.Println("err2 = ", err2)
	}

	// 转换Int
	num3, err3 := strconv.ParseInt("123", 10, 32)
	if err3 == nil {
		fmt.Println("num3 = ", num3)
	} else {
		fmt.Println("err3 = ", err3)
	}
	// 也可以通过Atoi，相当于ParseInt(s, 10, 0)
	num4, err4 := strconv.Atoi("888")
	if err4 == nil {
		fmt.Println("num4 = ", num4)
	} else {
		fmt.Println("err4 = ", err4)
	}

	// 结果为：
	// num1 =  true
	// num2 =  1.23
	// num3 =  123
	// num4 =  888
}
