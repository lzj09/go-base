package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

// 传输文件的客户端
func main() {
	fmt.Println("请输入文件路径：")
	var path string
	fmt.Scan(&path)

	// 获取文件相关信息
	info, err1 := os.Stat(path)
	if err1 != nil {
		fmt.Println("os.Stat err1 = ", err1)
		return
	}

	fileName := info.Name()
	fmt.Println("fileName:", fileName)

	// 连接服务端
	conn, err2 := net.Dial("tcp", "127.0.0.1:8080")
	if err2 != nil {
		fmt.Println("net.Dial err2 = ", err2)
		return
	}
	// 关闭连接
	defer conn.Close()

	// 向服务端发送文件名
	_, err3 := conn.Write([]byte(fileName))
	if err3 != nil {
		fmt.Println("conn.Write err3 = ", err3)
		return
	}

	// 接收服务端的响应
	buffer := make([]byte, 1024)
	n, err4 := conn.Read(buffer)
	if err4 != nil {
		fmt.Println("conn.Read err4 = ", err4)
		return
	}
	if "ok" == string(buffer[:n]) {
		// 向服务端发送文件
		sendFile(path, conn)
	} else {
		fmt.Println("服务端响应失败")
	}
}

// 向服务端发送文件
func sendFile(path string, conn net.Conn) {
	// 打开文件
	file, err1 := os.Open(path)
	if err1 != nil {
		fmt.Println("os.Open err1 = ", err1)
		return
	}
	// 关闭文件
	defer file.Close()

	buffer := make([]byte, 1024)
	for {
		n, err2 := file.Read(buffer)
		if err2 != nil {
			if err2 != io.EOF {
				fmt.Println("file.Read err2 = ", err2)
			}
			break
		}

		// 往服务端发送
		conn.Write(buffer[:n])
	}
}
