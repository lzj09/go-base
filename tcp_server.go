package main

import (
	"fmt"

	// 网络编程需要导入的包
	"net"
)

func main() {
	// 监听端口
	listener, err1 := net.Listen("tcp", "127.0.0.1:8080")
	if err1 != nil {
		fmt.Println("err1 = ", err1)
		return
	}
	defer listener.Close()

	conn, err2 := listener.Accept()
	if err2 != nil {
		fmt.Println("err2 = ", err2)
		return
	}
	defer conn.Close()

	buffer := make([]byte, 1024)
	num, err3 := conn.Read(buffer)
	if err3 != nil {
		fmt.Println("err3 = ", err3)
		return
	}
	fmt.Println("result: " + string(buffer[:num]))
}
