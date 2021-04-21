package main

import (
	"fmt"
)

// 切片
/*
 *	1、slice
 *	2、切片是数组的一个引用，因此切片是引用类型
 *	3、切片的使用和数组类似
 *	4、切片的长度是可以变化的，因此切片是一个可以动态变化的数组
 *	5、定义：var 切片名 []类型
 */
func main()  {
	var intArr = [5]int{11, 22, 33, 44, 55}
	// slice：切片名
	// intArr[1:3]：表示引用到 intArr 这个数组
	// 引用 intArr 数组的起始下标为 1，最后下标为 3，但是不包含 3
	slice := intArr[1:3]
	fmt.Println("intArr=", intArr)
	fmt.Println("slice e=", slice)
	fmt.Println("slice len=", len(slice))
	fmt.Println("slice cap=", cap(slice)) // 切片容量是动态变化的

	fmt.Printf("intArr[1]=%p\n", &intArr[1]) // intArr[1]=0xc00000a2a8
	fmt.Printf("slice[0]=%p\n", &slice[0]) // slice[0]=0xc00000a2a8
	slice[1] = 37
	fmt.Println("intArr=", intArr)
	fmt.Println("slice e=", slice)
}