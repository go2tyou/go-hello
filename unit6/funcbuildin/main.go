package main

import (
	"fmt"
)

// 内置函数
func main()  {
	// 1、len：求长度
	// 2、new：用来分配内存，主要用来分配值类型(int float32 struct)，返回的是指针
	num1 := 100
	fmt.Printf("num1=%v num1Type=%T num1=%v\n", num1, num1, &num1) // num1=100 num1Type=int num1=0xc000012090
	
	num2 := new(int)
	fmt.Printf("num2=%v num2Type=%T num2=%v\n", num2, num2, &num2) // num2=0xc0000120b0 num2Type=*int num2=0xc000006030
	fmt.Printf("num2=%v\n", *num2) // num2=0
	*num2 = 100
	fmt.Printf("num2=%v\n", *num2) // num2=100

	// 3、make 用来分配内存，主要用来分配引用类型(channel、map、slice)
}