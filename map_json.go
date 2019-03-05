package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// 通过map转换成json字符串

	// 定义一个map
	// map的key为string类型，值有的是string，有的是int等，所以可以定义为interface{}
	data := make(map[string]interface{}, 4)
	data["id"] = "001"
	data["name"] = "neil"
	data["course"] = []string{"Java", "Go", "JavaScript"}
	data["age"] = 28

	buff, err := json.Marshal(data)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	fmt.Println("json = ", string(buff))

	// 结果为：
	// json =  {"age":28,"course":["Java","Go","JavaScript"],"id":"001","name":"neil"}

	// 说明：map也可以跟结构体一样转换成json字符串
}
