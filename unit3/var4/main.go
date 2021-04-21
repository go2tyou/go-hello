package main

import "fmt"

// 变量使用的注意事项
func main() {

	// 该区域的数据值可以在同一类型范围内不断变化
	var i int = 10
	i = 20
	i = 30
	fmt.Println("i=", i)
	// i = 1.2 // 错误，不能改变数据类型

	// 变量在同一个作用域（在一个函数或代码块）内不能重名
	// var i int = 40
	// i := 50
}
