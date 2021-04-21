package main

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
func cal(x float64, y float64, operator byte) float64 {
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

func main() {
	var x float64 = 1
	var y float64 = 2
	var op byte = '+'
	// r := cal(1, 2, '+')
	r := cal(x, y, op)
	fmt.Printf("r=%f\n", r)
}