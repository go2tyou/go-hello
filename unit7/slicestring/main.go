package main

import (
	"fmt"
)

// 1、string 的底层是一个 byte 数组，因此可以进行切片处理
// 2、string 是不可变的，不能通过 str[0] = 'a' 的方式来改变
// 3、如果需要修改 string -> []byte 或 []rune => 修改 => 重新转为 string
func main()  {
	// 测试 1
	str := "helloworld"
	// 使用切片获取
	slice1 := str[5:]
	fmt.Println("slice1=", slice1)

	// 测试 3
	arr := []byte(str)
	arr[0] = 'g'
	str = string(arr)
	fmt.Println("str=", str)

	// 中文乱码 -> []byte 字节来处理
	// arr[0] = '赢' // constant 36194 overflows byte
	arr1 := []rune(str)
	arr1[0] = '赢'
	str = string(arr1)
	fmt.Println("str=", str)
}