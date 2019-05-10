package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	// 开启监听
	listener, err1 := net.Listen("tcp", "127.0.0.1:8080")
	if err1 != nil {
		fmt.Println("net.Listen err1 = ", err1)
		return
	}
	defer listener.Close()

	// 接收请求
	conn, err2 := listener.Accept()
	if err2 != nil {
		fmt.Println("listener.Accept err2 = ", err2)
		return
	}
	defer conn.Close()

	// 接收文件名
	buffer := make([]byte, 1024)
	n, err3 := conn.Read(buffer)
	if err3 != nil {
		fmt.Println("conn.Read err3 = ", err3)
		return
	}
	fileName := string(buffer[:n])

	// 向客户端回复ok
	conn.Write([]byte("ok"))

	// 接收文件内容
	receiveFile(fileName, conn)
}

// 接收文件内容
func receiveFile(fileName string, conn net.Conn) {
	// 在当前目录下创建文件
	file, err1 := os.Create(fileName)
	if err1 != nil {
		fmt.Println("os.Create err1 = ", err1)
		return
	}
	defer file.Close()

	buffer := make([]byte, 1024)
	for {
		// 接收文件内容
		n, err2 := conn.Read(buffer)
		if err2 != nil {
			if err2 != io.EOF {
				fmt.Println("conn.Read err2 = ", err2)
			}
			return
		}

		file.Write(buffer[:n])
	}
}
