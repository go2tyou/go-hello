package main

import (
	"fmt"
)

// 累加器
// AddUpper 是一个函数，返回的数据类型是 func (int) int
func AddUpper() func (int) int {
	// 返回的是一个匿名函数，但是这个匿名函数引用到函数外的 n，因此，这个匿名函数就和 n 形成一个整体，构成闭包
	// 当反复调用时，因为 n 只初始化一次，因此每调用一次就进行累计
	var n int = 10
	var str string = "hello"
	return func (x int) int {
		n = n + x
		str += "a"
		fmt.Println("str=", str)
		return n
	}
}

// 闭包就是一个函数和与其相关的引用环境组合的一个整体
// 闭包的关键，就是要分析出返回的函数它引用到哪些变量，因为函数与它引用的变量构成
func main()  {
	f := AddUpper()
	fmt.Println(f(1)) 
	fmt.Println(f(2))
	fmt.Println(f(3))
}