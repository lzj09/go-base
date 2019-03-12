package main

import (
	"fmt"
	"io"
	"os"
)

// 复制文件函数
// sourcePath：源文件的地址
// distPath：目的文件的地址
func copyFile(sourcePath, distPath string) {
	// 打开源文件
	sourceFile, sourceErr := os.Open(sourcePath)
	if sourceErr != nil {
		fmt.Println("sourceErr = ", sourceErr)
		return
	}

	// 创建目的文件
	distFile, distErr := os.Create(distPath)
	if distErr != nil {
		fmt.Println("distErr = ", distErr)
		return
	}

	// 定义读取文件的buffer
	buffer := make([]byte, 1024)

	for {
		n, err := sourceFile.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("err = ", err)
		}

		// 将读取出来的内容写入目的文件
		distFile.Write(buffer[:n])
	}

	fmt.Println("完成复制文件")

	// 关闭文件
	defer sourceFile.Close()
	defer distFile.Close()
}

// 复制文件
func main() {
	copyFile("./source.avi", "./dist.avi")
}
