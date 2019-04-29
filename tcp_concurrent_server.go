package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	// 监听端口
	listener, err1 := net.Listen("tcp", "127.0.0.1:8080")
	if err1 != nil {
		fmt.Println("err1 = ", err1)
		return
	}
	defer listener.Close()

	for {
		// 接收请求
		conn, err2 := listener.Accept()
		if err2 != nil {
			fmt.Println("err2 = ", err2)
			break
		}

		// 开启协助处理请求信息
		go handlerReq(conn)
	}
}

// 处理请求信息
func handlerReq(conn net.Conn) {
	defer conn.Close()
	// 获取客户端地址信息
	addr := conn.RemoteAddr()
	fmt.Println(addr, " 连接成功...")

	// 读取信息缓存
	buffer := make([]byte, 1024)

	for {
		num, err := conn.Read(buffer)
		if err != nil {
			fmt.Println(addr, " 报错：", err)
			break
		}
		content := string(buffer[:num])
		if "exit" == content {
			// 断开连接
			break
		}

		fmt.Println("client: ", content)
		// 往客户端发送信息
		conn.Write([]byte("server: " + strings.ToUpper(content)))
	}
}
