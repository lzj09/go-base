package main

import (
	"fmt"
)

// 定义Person结构体
type Person struct {
	id   string
	name string
	sex  byte
	age  int
}

// 为Person结构绑定方法setValue01
func (p Person) setValue01(id string, name string, sex byte, age int) {
	p.id = id
	p.name = name
	p.sex = sex
	p.age = age

	fmt.Println("setValue01: p = ", p)
}

// 为Person结构体绑定方法setValue02
func (p *Person) setValue02(id string, name string, sex byte, age int) {
	p.id = id
	p.name = name
	p.sex = sex
	p.age = age

	fmt.Println("setValue02: p = ", p)
}

func main() {
	// 定义并初始化Person
	p := Person{id: "1", name: "neil", sex: 'm', age: 28}
	fmt.Println("main: p = ", p)

	// 调用setValue01方法
	p.setValue01("2", "jane", 'f', 18)

	// 方法值的使用
	methodValue := p.setValue01       // 将p.setValue01赋值给变量methodValue
	methodValue("2", "jane", 'f', 18) // 等价于 p.setValue01(2, "jane", 'f', 18)

	// 调用setValue02方法
	p.setValue02("3", "chris", 'm', 20)

	// 方法表达式的使用
	methodExpr := (*Person).setValue02
	methodExpr(&p, "3", "chris", 'm', 20) // 注意：对于有参数的方法，此处的调用方法

	// 结果为：
	// main: p =  {1 neil 109 28}
	// setValue01: p =  {2 jane 102 18}
	// setValue01: p =  {2 jane 102 18}
	// setValue02: p =  &{3 chris 109 20}
	// setValue02: p =  &{3 chris 109 20}
}
