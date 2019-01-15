package main

import (
	"fmt"
)

func test(data map[int]string) {
	// 删除其中一个元素
	delete(data, 3)
}

func main() {
	data := map[int]string{1: "go", 2: "java", 3: "javascript"}
	fmt.Println("调用函数前：data = ", data)

	// map类型作函数参数
	test(data)
	fmt.Println("调用函数后：data = ", data)

	// 结果为：
	// 调用函数前：data =  map[1:go 2:java 3:javascript]
	// 调用函数后：data =  map[1:go 2:java]

	// 结论：
	// map作函数参数是引用传递
}
