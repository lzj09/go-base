package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

// 定义结构体
type Client struct {
	// 传递信息的管道
	C chan string
	// 名称
	Name string
	// 地址
	Addr string
}

// 定义全局map，用于存储已经连接的客户端信息
var onlineClient map[string]Client

// 定义全局广播信息的管道
var message chan string

// 向每个客户端广播信息
func BroadcastInfo() {
	for {
		msg := <-message
		for _, client := range onlineClient {
			client.C <- msg
		}
	}
}

// 处理连接
func HandlerConn(conn net.Conn) {
	defer conn.Close()

	// 获取连接地址信息
	addr := conn.RemoteAddr().String()

	// 封装Client信息
	client := Client{make(chan string), addr, addr}
	onlineClient[addr] = client

	message <- WrapMsg(client, "login")

	// 发送信息
	go SendMsg(client, conn)

	// 是否断开标志
	hasQuit := make(chan bool)
	// 是否超时标志
	hasTimeOut := make(chan bool)

	// 接收客户端的消息
	go func() {
		buffer := make([]byte, 1024)
		for {
			n, err := conn.Read(buffer)
			if n == 0 {
				hasQuit <- true
				fmt.Println("conn.Read err = ", err)
				return
			}

			msg := string(buffer[:n-2])
			if len(msg) == 3 && msg == "who" {
				conn.Write([]byte("User List:\n"))
				for _, cli := range onlineClient {
					conn.Write([]byte(cli.Name + "\n"))
				}
			} else if len(msg) >= 7 && msg[:6] == "rename" {
				name := strings.Split(msg, " ")[1]
				client.Name = name
				onlineClient[addr] = client

				conn.Write([]byte("rename ok\n"))
			} else {
				message <- WrapMsg(client, msg)
			}
			hasTimeOut <- true
		}
	}()

	for {
		select {
		case <-hasQuit:
			delete(onlineClient, addr)
			message <- WrapMsg(client, "logout")
			return
		case <-hasTimeOut:
		case <-time.After(60 * time.Second):
			delete(onlineClient, addr)
			message <- WrapMsg(client, "time out")
			return
		}

	}
}

// 发送信息
func SendMsg(client Client, conn net.Conn) {
	for msg := range client.C {
		conn.Write([]byte(msg))
	}
}

// 包装信息
func WrapMsg(client Client, msg string) (res string) {
	res = "[" + client.Name + "]: " + msg + "\n"
	return
}

func main() {
	// 监听端口
	listener, err1 := net.Listen("tcp", "127.0.0.1:8080")
	if err1 != nil {
		fmt.Println("net.Listen err1 = ", err1)
		return
	}
	defer listener.Close()

	// 分配空间onlineClient
	onlineClient = make(map[string]Client)
	message = make(chan string)

	// 向每个客户端广播信息
	go BroadcastInfo()

	// 循环获取连接
	for {
		conn, err2 := listener.Accept()
		if err2 != nil {
			fmt.Println("listener.Accept err2 = ", err2)
			continue
		}

		// 处理连接
		go HandlerConn(conn)
	}
}
