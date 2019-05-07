package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// 连接服务端
	conn, err1 := net.Dial("tcp", "127.0.0.1:8080")
	if err1 != nil {
		fmt.Println("err1 = ", err1)
		return
	}

	// 关闭连接
	defer conn.Close()

	// 开启协程，接收键盘的输入
	go func() {
		buffer := make([]byte, 1024)
		for {
			num, err2 := os.Stdin.Read(buffer)
			if err2 != nil {
				fmt.Println("err2 = ", err2)
				break
			}

			// 将内容发送给服务端
			conn.Write(buffer[:num])
		}
	}()

	// 接收服务端的内容
	buf := make([]byte, 1024)
	for {
		num, err3 := conn.Read(buf)
		if err3 != nil {
			fmt.Println("err3 = ", err3)
			break
		}

		fmt.Println("服务端内容: ", string(buf[:num]))
	}
}
