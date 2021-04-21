package main

import (
	"fmt"
	"strconv"
)

// go 中基本数据类型与 string 类型互转
func main() {
	// 基本数据类型 --> stirng
	var x int = 99
	var y float64 = 88.77
	var z bool = true
	var v byte = 'h'
	var str string
	// 方式一、fmt.Sprintf("%参数", 表达式)
	str = fmt.Sprintf("%d", x)
	fmt.Printf("str type is %T str=%q\n", str, str)

	str = fmt.Sprintf("%f", y)
	fmt.Printf("str type is %T str=%q\n", str, str)

	str = fmt.Sprintf("%t", z)
	fmt.Printf("str type is %T str=%q\n", str, str)

	str = fmt.Sprintf("%c", v)
	fmt.Printf("str type is %T str=%q\n", str, str)

	fmt.Printf("---------------------------------\n")

	//  方式二、使用 strconv 包的函数
	str = strconv.FormatInt(int64(x), 10)
	fmt.Printf("str type is %T str=%q\n", str, str)

	str = strconv.FormatFloat(y, 'f', 10, 64)
	fmt.Printf("str type is %T str=%q\n", str, str)

	str = strconv.FormatBool(z)
	fmt.Printf("str type is %T str=%q\n", str, str)

	str = strconv.Itoa(x)
	fmt.Printf("str type is %T str=%q\n", str, str)
}
