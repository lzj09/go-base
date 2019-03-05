package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	// 属性名称的首字母，只有大写才能被转到json中
	// 由于首字母是大写，导致json串中也是大写，这里可以通设置属性标签使其转换成小写
	Id     string   `json:"id"` // 在json中变成首字母小写
	Name   string   `json:"name"`
	Course []string `json:"course"`
	Age    int      `json:"age"`
}

func main() {
	// 利用json字符串转换成结构体
	// json字符串
	json_str := `{"age":28,"course":["Java","Go","JavaScript"],"id":"001","name":"neil"}`

	// 定义学生结构体
	var student Student

	//Unmarshal方法的第1个参数是[]byte切片，第2个参数是空接口类型
	// 所以得将json字符串转换成[]byte切片，同时为了能够将结果赋值到student变量，所以第2个参数得地址传递
	// 返回值为错误信息
	err := json.Unmarshal([]byte(json_str), &student)

	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	fmt.Printf("student = %+v\n", student)

	// 结果为：
	// student = {Id:001 Name:neil Course:[Java Go JavaScript] Age:28}
}
