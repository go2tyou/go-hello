package main

import "fmt"

func main() {
	// 定义变量
	var i int
	// 给 i 赋值
	i = 10
	// 使用变量
	fmt.Println("i=", i)
}

/**
 * 1、变量 = 变量名 + 值 + 数据类型
 * 2、go 的变量如果没有赋初值，编译器会使用默认值 
 * 3、声明变量
 * 		基本语法：var 变量名 数据类型
 * 4、初始化变量
 * 5、给变量赋值
 */