package main

import (
	"fmt"
	"io"
	"os"
)

// 定义一个写文件的函数
// 参数为文件路径
func write(path string) {
	// 创建文件
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	for i := 0; i < 20; i++ {
		file.WriteString(fmt.Sprintf("i的值为 = %d\n", i))
	}

	// 关闭文件
	defer file.Close()
}

// 定义一个读取文件
func read(path string) {
	// 打开文件
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	res := ""

	// 定义每个获取文件内容
	buffer := make([]byte, 32)
	for {
		n, err1 := file.Read(buffer)

		if err1 != nil {
			if err1 == io.EOF {
				res += string(buffer[:n])
				break
			}

			fmt.Println("err1 = ", err1)
			return
		}

		res += string(buffer[:n])
	}

	fmt.Println("res: ", res)

	// 关闭文件
	defer file.Close()
}

// 测试文件读写操作
func main() {
	// 创建文件
	path := "./test.txt"

	// 写文件
	write(path)

	// 读文件
	read(path)

}
