package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// 抓取网页信息
func GetWebInfo(url string) (result string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	body := resp.Body
	defer body.Close()

	// 缓存数据
	buffer := make([]byte, 1024)
	for {
		n, err1 := body.Read(buffer)
		if n == 0 {
			if err1 != io.EOF {
				err = err1
			}
			break
		}

		result += string(buffer[:n])
	}

	return
}

func main() {
	var url string
	fmt.Printf("请输入需要抓取的url:")
	fmt.Scan(&url)

	// 抓取网页信息
	result, err1 := GetWebInfo(url)
	if err1 != nil {
		fmt.Println("GetWebInfo err1 = ", err1)
		return
	}

	// 保存网页
	file, err2 := os.Create("web-info.html")
	if err2 != nil {
		fmt.Println("os.Create err2 = ", err2)
		return
	}
	file.WriteString(result)
	file.Close()
}
