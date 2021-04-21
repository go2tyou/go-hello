package main

import (
	"fmt"
)

// 全局匿名函数
var (
	// fun1 就是一个全局匿名函数
	Fun1 = func (n1 int, n2 int) int {
		return n1 * n2
	}
)

// 匿名函数
func main()  {
	// 在定义匿名函数时就直接调用，这种方式匿名函数只能调用一次

	// 求两个数的和，使用匿名函数的方式
	res := func (n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)
	fmt.Println("res=", res)

	// 先赋值给一个变量，在通过变量完成调用
	// 将匿名函数 func (n1 int, n2 int) int 赋值给 a 变量
	// 则 a 的数据类型就是函数类型，此时，可以通过 a 完成调用
	a := func (n1 int, n2 int) int {
		return n1 - n2
	}
	res2 := a(20, 10)
	res3 := a(30, 10)
	fmt.Println("res2=", res2)
	fmt.Println("res3=", res3)
	
	// 全局匿名函数使用
	res4 := Fun1(11, 11)
	fmt.Println("res4=", res4)
}