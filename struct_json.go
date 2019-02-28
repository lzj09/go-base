package main

import (
	"fmt"
	// 需要处理json相关操作，则需要导入此包
	"encoding/json"
)

// 定义学生结构体
type Student1 struct {
	// 属性名称的首字母，只有大写才能被转到json中
	Id     string
	Name   string
	Course []string
	Age    int
}

type Student2 struct {
	// 属性名称的首字母，只有大写才能被转到json中
	// 由于首字母是大写，导致json串中也是大写，这里可以通设置属性标签使其转换成小写
	Id     string   `json:"id"` // 在json中变成首字母小写
	Name   string   `json:"name"`
	Course []string `json:"course"`
	Age    int      `json:"age,string"` // 在json中变成首字母小写，同时值变成字符串
}

// 通过结构体构建json数据
func main() {

	s1 := Student1{Id: "001", Name: "neil", Course: []string{"Java", "Go", "JavaScript"}, Age: 28}

	// 此方法返回2个值，第1值是字节切片，第2值是错误信息
	json1, err1 := json.Marshal(s1)
	if err1 == nil {
		fmt.Println(string(json1))
	} else {
		fmt.Println("err = ", err1)
	}

	// 结果为：
	// {"Id":"001","Name":"neil","Course":["Java","Go","JavaScript"],"Age":28}

	s2 := Student2{Id: "001", Name: "neil", Course: []string{"Java", "Go", "JavaScript"}, Age: 28}

	// 此方法返回2个值，第1值是字节切片，第2值是错误信息
	json2, err2 := json.Marshal(s2)
	if err2 == nil {
		fmt.Println(string(json2))
	} else {
		fmt.Println("err2 = ", err2)
	}

	// 结果为：
	// {"id":"001","name":"neil","course":["Java","Go","JavaScript"],"age":"28"}
}
