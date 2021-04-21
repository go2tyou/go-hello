package main

import "fmt"

func main() {
	// go 变量使用方式一
	// 指定变量类型，声明后若不赋值，使用默认值
	var i int
	fmt.Println("i=", i)

	// go 变量使用方式二
	// 根据值自动判断变量类型（类型推导）
	var num = 10.11
	fmt.Println("num=", num)

	// go 变量使用方式三
	// 省略 var，注意 := 左侧的变量不应该是已经声明过的，否则会编译错误
	// := 不能省略，否则错误
	str := "sss"
	fmt.Println("str=", str)


}
