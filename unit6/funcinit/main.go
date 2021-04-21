package main

import (
	"fmt"
)

var age = test()
// 为了看到全局变量先被初始化
func test() int {
	fmt.Println("test()...")
	return 90
}

// 每一个源文件都可以包含一个 init 函数
// 通常可以在 init() 中完成初始化工作
/*
	细节：
	1、如果一个文件中同时包含全局变量定义、init()、main()，那么执行流程：全局变量定义 -> init() -> main()
	2、init() 最主要的作用，就是完成初始化工作
*/
func init()  {
	fmt.Println("init()...")
}

func main()  {
	fmt.Println("main()...age=", age)
}