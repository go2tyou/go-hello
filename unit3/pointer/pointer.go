package main

import (
	"fmt"
)

/**
 * 指针
 * 1、值类型，都有对应的指针类型，形式为 *数据类型
 * 2、值类型包括基本数据类型、数组、结构体
 */
func main() {
	// // 基本数据类型在内存布局
	// // i 的地址
	// var i int = 10
	// fmt.Println("i 的地址=", &i) // i 的地址= 0xc000012090
	
	// // 指针类型,指针变量存的是一个地址，这个地址指向的空间存的才是值
	// // 指针在内存的布局
	// // 1、ptr 是一个指针变量
	// // 2、ptr 本身的值是 &i
	// var ptr *int = &i
	// fmt.Printf("ptr=%v\n", ptr) // ptr=0xc000012090
	// fmt.Println("ptr 的地址=", &ptr) // ptr 的地址= 0xc000006030
	
	// // 获取指针类型所指向的值
	// fmt.Printf("ptr 指向的值=%v\n", *ptr) // ptr 指向的值=10

	var num int = 9
	fmt.Printf("num 的地址 %v\n", &num)

	var ptr *int
	ptr = &num
	*ptr = 10
	fmt.Println("num=", num)
}
