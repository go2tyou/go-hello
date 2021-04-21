package main

// 函数的传递方式：
// 		本质 -> 传递给函数的都是变量的副本
// 1、值传递  -> 值拷贝
// 		基本数据类型、数组、结构体
// 2、引用传递 -> 地址拷贝
// 		指针、silce 切片、map、管道、interface

// 变量作用域：
// 1、局部变量
// 		仅限于函数内部（变量名大写也没卵用）
// 2、全局变量
// 		整个包有效。如果首字母大写，整个程序都有效
// 3、变量在代码块，如 for/if，作用域只在该代码块有效
// 4、就近原则
// 5、赋值语句不能在函数体外
// Name := "sr" // syntax error: non-declaration statement outside function body <=> var Name string Name = "tom"

import (
	"fmt"
)

func printPyramid(totalLevel int) {
	// i 表示层数
	for i := 1; i <= totalLevel; i++ {
		for k := 1; k <= totalLevel - i; k++ {
			fmt.Print(" ")
		}
		// j 表示每层打印 * 的个数
		for j := 1; j <= 2 * i - 1; j++ {
			if j == 1 || j == 2 * i - 1 || i == totalLevel {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func printMulti(n int) {
	// i 表示层数
	for i := 1; i <= n; i++ {
		// j 表示表达式个数
		for j := 1; j <= i; j++ {
			fmt.Printf("%v x %v = %v\t", j, i, i * j)
		}
		fmt.Println()
	}
}

func main()  {
	printPyramid(8)
	printMulti(9)
}