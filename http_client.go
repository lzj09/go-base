package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// 发送http请求
	resp, err1 := http.Get("http://www.baidu.com")
	if err1 != nil {
		fmt.Println("http.Get = ", err1)
		return
	}

	body := resp.Body
	defer body.Close()

	// 设置读取数据缓冲区
	buffer := make([]byte, 1024)

	// 网页内容
	var content string

	for {
		n, err := body.Read(buffer)
		if n == 0 {
			if err != nil && err != io.EOF {
				fmt.Println("body.Read = ", err)
			}
			break
		}

		content += string(buffer[:n])
	}

	fmt.Println("body = ", content)
}
