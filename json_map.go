package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// 将json字符串转换成map
	// json字符串
	json_str := `{"age":28,"course":["Java","Go","JavaScript"],"id":"001","name":"neil"}`

	// 定义一个map
	data := make(map[string]interface{}, 4)

	//Unmarshal方法的第1个参数是[]byte切片，第2个参数是空接口类型
	// 所以得将json字符串转换成[]byte切片，同时为了能够将结果赋值到student变量，所以第2个参数得地址传递
	// 返回值为错误信息
	err := json.Unmarshal([]byte(json_str), &data)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	// 打印出data的内容
	fmt.Printf("data = %+v\n", data)
	// 结果为：
	// data = map[name:neil age:28 course:[Java Go JavaScript] id:001]

	// 如果要获取id的具体值，id的值为string
	//	var id string
	//	id = data["id"]  // 此处会报：cannot use data["id"] (type interface {}) as type string in assignment: need type assertion
	//	fmt.Printf("id = %s/n", id)

	// 由上面的报错信息可知，如果要获取data里面的具体值，需要通过断言才能够得到
	for key, value := range data {
		switch item := value.(type) {
		case string:
			if key == "id" {
				fmt.Printf("id的值为%s\n", item)
			}
		}
	}
	// 这样是可以获取到id的值，就是比较麻烦，结果为：
	// id的值为001

}
