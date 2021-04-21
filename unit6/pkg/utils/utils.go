package utils

import (
	"fmt"
)

// 函数：为完成某一功能的程序指令的集合
/*
 *	func 函数名(形参列表) (返回值列表) {
 *		执行语句...	
 *		return 返回值列表
 *	}
 */
 // 首字母 C 需要大写，才能被其他包的文件访问，类似其他语言 public
func Cal(x float64, y float64, operator byte) float64 {
	var r float64
	switch operator {
	case '+':
		r = x + y
	case '-':
		r = x - y
	case '*':
		r = x * y
	case '/':
		r = x / y
	default:
		fmt.Println("err")		
	}
	return r
}
