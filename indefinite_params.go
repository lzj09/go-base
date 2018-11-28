package main

import (
	"fmt"
)

func test01(args ...int) {
  // 全部传递
  //	test02(args...)

  // 将从下标为2的元素起（包含该元素），到最后的元素传递
	//	test02(args[2:]...)

  // 将从下标为0元素起，到下标为2的元素传递（不包含该元素）
	test02(args[:2]...)
}

func test02(args ...int) {
	for _, data := range args {
		fmt.Println("data = ", data)
	}
}

func main() {
	test01(1, 2, 3, 4)
}
