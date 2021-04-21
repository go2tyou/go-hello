package main

import "fmt"

/**
 * 1、浮点数在机器中的存放形式：浮点数 = 符号位 + 指数位 + 尾数位 （浮点数都是有符号的）
 * 2、尾数部分可能丢失，造成精度损失
 * 3、浮点类型有固定的范围和字段长度，不受具体 OS 的影响
 * 4、浮点型默认 float64 类型
 * 5、开发中推荐使用 float64
 */

// go 中小数类型的使用
func main() {

	var price float32 = 88.99
	fmt.Println("price=", price)
	
	var num1 float32 = -123.0000901
	var num2 float64 = -123.0000901
	// num1= -123.00009
	fmt.Println("num1=", num1)
	// num2= -123.0000901
	fmt.Println("num2=", num2)
	
	var x = 1.1
	fmt.Printf("x 的类型为 %T \n", x)
	
	// 十进制数形式
	m := 5.12
	n := .123 // 必须有小数点
	fmt.Println("m=", m, "n=", n)
	
	// 科学计数法
	o := 5.1234e2 // 5.1234 * 10 的 2 次方
	p := 5.1234E2 // 5.1234 * 10 的 2 次方
	q := 5.1234e-2 // 5.1234 / 10 的 2 次方
	fmt.Println("o=", o)
	fmt.Println("p=", p)
	fmt.Println("q=", q)

}
