package main

import (
	"fmt"
)

// 关系运算符
// 结果为 bool
// 用在 if 结构和循环结构
func main()  {
	var n1 int = 10
	var n2 int = 20
	// fmt.Println(n1 = n2) // syntax error: unexpected =, expecting comma or )
	fmt.Println(n1 == n2)

	flag := n1 > n2
	fmt.Println(flag)
}