package main

import (
	"fmt"
	"net"
)

func main() {
	// 连接服务端
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	defer conn.Close()

	// 往服务端写内容
	conn.Write([]byte("测试服务端..."))
}
