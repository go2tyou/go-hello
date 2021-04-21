package main

import "fmt"

/**
 * 1、包名与目录保持一致
 * 2、变量名、函数名、常量，采用驼峰法命名
 * 3、首字母大写是公有的，首字母小写是私有的
 */

// go 中标识符的使用
func main() {

	// 严格区分大小写
	var num int = 10
	var Num int = 20
	fmt.Printf("num=%v Num=%v\n", num, Num)
	
	// 不能包含空格
	// var ab c int = 30 // syntax error: unexpected int at end of statement
	
	// _ 在 go 中是一个特殊标识符，称为空标识符
	// 可以代表任何其他的标识符，但是它对应的值会被忽略（比如忽略某个返回值）
	// 仅能作为占位符使用，不能作为标识符使用
	// var _ int = 40
	// fmt.Printf("_=%v\n", _) // cannot use _ as value

	// 不能使用保留关键字
	// var if int = 50 // syntax error: unexpected if, expecting name
}
