package main

import (
	"fmt"
)

// 算数运算符
func main()  {
	// 如果运算的数都是整数，那么结果，去掉小数部分，保留整数部分
	fmt.Println(10/4) // 2

	var n1 float32 = 10 / 4
	fmt.Println(n1) // 2

	// 想保留小数
	var n2 float32 = 10.0 / 4
	fmt.Println(n2) // 2.5

	// a % b = a - a / b * b
	fmt.Println(10 % 3) // 1
	fmt.Println(-10 % 3) // -1
	fmt.Println(10 % -3) // 1
	fmt.Println(-10 % -3) // -1

	// ++ --
	var i int = 10
	i++
	fmt.Println("i=", i)
	i--
	fmt.Println("i=", i)
}