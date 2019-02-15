package main

import (
	"fmt"
)

// 定义学生结构体
type Student struct {
	id   int
	name string
}

// 定义一个函数，参数为空接口类型
// 注意：空接口类型为：interface{}
func test(i interface{}) {
	fmt.Println("i = ", i)
}

func main() {
	// 对于空接口类型，即是任何类型都可以支持
	// 测试调用test函数，由于接收空接口类型，则表示可以接收任何类型的参数
	test(123)
	test("neil")
	test(Student{id: 1, name: "jane"})

	// 结果为：
	// i =  123
	// i =  neil
	// i =  {1 jane}

	// 下面通过两种方法来判断空接口对应的具体类型

	// 新建一个切片，类型为空接口类型
	arrs := make([]interface{}, 3)
	arrs[0] = 123
	arrs[1] = "neil"
	arrs[2] = Student{id: 1, name: "jane"}

	// 第1种方式：通过if来判断
	for index, data := range arrs {
		// 此处调用data.(int)的第1个返回值为具体值，第2个值是判断该值的类型是不是int，如果是则返回true,否则false
		if val, ok := data.(int); ok == true {
			fmt.Printf("arrs[%d]的类型为int, 值为%d\n", index, val)
		} else if val, ok := data.(string); ok == true {
			fmt.Printf("arrs[%d]的类型为string, 值为%s\n", index, val)
		} else if val, ok := data.(Student); ok == true {
			fmt.Printf("arrs[%d]的类型为Student, 值为%v\n", index, val)
		}
	}
	// 结果为：
	// arrs[0]的类型为int, 值为123
	// arrs[1]的类型为string, 值为neil
	// arrs[2]的类型为Student, 值为{1 jane}

	// 第2种方式：通过switch来判断
	for index, data := range arrs {
		switch val := data.(type) {
		case int:
			fmt.Printf("arrs[%d]的类型为int, 值为%d\n", index, val)
		case string:
			fmt.Printf("arrs[%d]的类型为string, 值为%s\n", index, val)
		case Student:
			fmt.Printf("arrs[%d]的类型为Student, 值为%v\n", index, val)
		}
	}
	// 结果为：
	// arrs[0]的类型为int, 值为123
	// arrs[1]的类型为string, 值为neil
	// arrs[2]的类型为Student, 值为{1 jane}
}
