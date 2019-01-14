package main

import (
    "fmt"
)

func main() {
    data := map[int]string{1: "go", 2: "java", 3: "javascript"}

    // 循环遍历元素，返回第一个元素为键，第二元素为值
    for key, value := range data {
        fmt.Printf("%d ---> %s\n", key, value)
    }

    // 判断某一个键是否存在
    value, ok := data[2]
    if ok {
        fmt.Println("键存在，并且值为：", value)
    } else {
        fmt.Println("不存在的键")
    }

    // 删除元素
    delete(data, 3)
    fmt.Println("data = ", data)

    // 结果为：
    // 1 ---> go
    // 2 ---> java
    // 3 ---> javascript
    // 键存在，并且值为： java
    // data =  map[2:java 1:go]
}
