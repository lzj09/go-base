package test

// 定义一个结构体，首字母小写，则只能同包级别访问
type student struct {
	id   int
	name string
}

// 定义一个结构体，首字母大写，则其它包级别可以访问
type Student struct {
	// 定义的成员变量，首字母小写，则只能同包级别访问
	id int

	// 定义的成员变量，首字母大写，则其它包级别可以访问
	Name string
}
