package main

import (
    "fmt"
)

func main() {
    // 定义方法一：map的格式为：map[keyType]valueType
    var data1 map[int]string
    fmt.Println("data1 = ", data1)

    // 定义方法二：make(map[keyType]valueType, len)，同时len也可以不指定
    data2 := make(map[int]string)
    fmt.Println("data2 = ", data2)

    // 定义并初始化
    data3 := map[int]string{1: "go", 2: "java", 3: "javascript"}
    fmt.Println("data3 = ", data3)

    // 会报：panic: assignment to entry in nil map错误，因为data1是没有初始化的
    //    data1[1] = "hello"

    // 为map新增值，或者修改值，当key存在时，就是修改该值
    data2[1] = "data"

    // 对于data3[2]之前的值为java，此操作即修改该值，即变成为javaweb
    data3[2] = "javaweb"

    fmt.Println("data1 = ", data1)
    fmt.Println("data2 = ", data2)
    fmt.Println("data3 = ", data3)

    // 结果为：
    // data1 =  map[]
    // data2 =  map[]
    // data3 =  map[2:java 3:javascript 1:go]
    // data1 =  map[]
    // data2 =  map[1:data]
    // data3 =  map[3:javascript 1:go 2:javaweb]

    // 注意：
    // map的key必须是唯一的，同时内容是无序的
}
