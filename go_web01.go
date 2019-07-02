package main

import (
	"fmt"
	"net/http"
)

// 处理请求方法
func hander(write http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(write, "hello world, %s!", request.URL.Path[0:])
}

func main() {
	// 设置路由处理方法
	http.HandleFunc("/", hander)

	// 设置监听端口
	http.ListenAndServe(":8000", nil)
}
