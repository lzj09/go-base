package main

import (
	"fmt"
	"os"
)

// 通过文件的路径获取文件的信息
func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("输入的信息错误")
		return
	}

	// 文件路径
	filePath := args[1]

	// 获取文件信息
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	fmt.Println("文件名：", fileInfo.Name())
	fmt.Println("文件大小：", fileInfo.Size())
}
