package main

import (
	"fmt"
)

// 进制
func main()  {
	var i int = 5
	// 二进制输出
	fmt.Printf("i=%b\n", i)

	// 八进制 --> 0-7; 满 8 进 1; 以数字 0 开头
	var j int = 011
	fmt.Println("j=", j)

	// 十六进制 --> 0-9及A-F; 满 16 进 1; 以 0x 或 0X 开头
	var k int = 0x11
	fmt.Println("k=", k)
}