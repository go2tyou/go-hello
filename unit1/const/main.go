package main

import (
	"fmt"
)

/*
 *	1、常量在定义时，必须初始化
 *	2、常量不能修改
 *	3、常量只能修饰 bool、属值类型(int、float 系列)、string 类型
 *	4、const identifier [type] = value -> type 可不写（类型推导）
 *	5、golang 中没有常量名必须字母大写的规范
 *	6、仍然通过首字母的大小来控制常量的访问范围，如 "time"
 */
func main()  {
	// const tax int // const declaration cannot have type without expression missing value in const declaration

	// const n int = 0
	// n = 100 // cannot assign to n

	// const (
	// 	a = 1
	// 	b = 2
	// )
	// fmt.Println(a, b)

	const (
		a = iota
		b
		c
	)
	fmt.Println(a, b, c) // 0 1 2
}