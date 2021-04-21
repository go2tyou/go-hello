package main

import (
	"fmt"
)

type MethodUtils struct {
	// 结构体可以不要字段
}

func (mu *MethodUtils) JudgeNum(num int) {
	if num % 2 == 0 {
		fmt.Println(num, "is ou~")
	} else {
		fmt.Println(num, "is ji~")
	}
}

func (mu *MethodUtils) Print(m int, n int, char string) {
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(char)
		}
		fmt.Println()
	}
}

type Calcuator struct {
	Num1 float64
	Num2 float64
}

func (cal *Calcuator) getRes(op byte) float64 {
	res := 0.0
	switch op {
	case '+':
		res = cal.Num1 + cal.Num2
	case '-':
		res = cal.Num1 - cal.Num2
	case '*':
		res = cal.Num1 * cal.Num2
	case '/':
		res = cal.Num1 / cal.Num2
	default:
		fmt.Println("运算符不对")	
	}
	return res
}

func main()  {
	var mu MethodUtils
	mu.JudgeNum(10)
	mu.JudgeNum(1)
	mu.Print(3, 4, "`")

	var cal Calcuator
	cal.Num1 = 1.2
	cal.Num2 = 2.3
	res := cal.getRes('+')
	// res = cal.getRes('%')
	fmt.Println("res=", res) // res= 3.5
}