package main

import "fmt"
/**
 * 1、字符在 ASCII 表，直接用 byte 保存
 * 2、字符对应码值大于 255, 考虑用 int
 * 3、要按照字符形式输出，使用 fmt.Printf("c1=%c", c1)
 * 4、字符使用 UTF-8 编码
 * 5、字符的本质是一个整数
 * 6、字符类型是可以运算的（看 5）
 */

// go 中字符的使用
func main() {

	var c1 byte = 'a'
	var c2 byte = '0'
	// 当直接输出 byte 值，即是输出了字符对应的十进制码值
	// c1= 97 
	fmt.Println("c1=", c1)
	// c2= 48
	fmt.Println("c2=", c2)
	// 格式化输出
	fmt.Printf("c1=%c c2=%c \n", c1, c2)
	
	// var c3 byte = '北' // constant 21271 overflows byte
	var c3 int = '北'
	fmt.Printf("c3=%c", c3)
}
