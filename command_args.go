package main

import (
	"fmt"
	"os"
)

func main() {
	// 获取命令行参数
	args := os.Args

	for i, item := range args {
		fmt.Printf("args[%d] = %s\n", i, item)
	}

	// 首先build文件，使其构建成可执行文件，即：
	// go build command_args.go

	// 然后执行并输入参数，即：
	// command_args.exe 10 20

	// 结果为：
	// args[0] = command_args.exe
	// args[1] = 10
	// args[2] = 20

	// 注意：它会将本身command_args.exe作为第1个参数
}
