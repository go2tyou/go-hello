package main

import (
	"fmt"
)

// 使用细节
/*
 *	1、切片初始化：var slice = arr[startIndex:endIndex]
 *	2、切片初始化时，不能越界。范围 [0-len(arr)] 之间，但是可以动态增长
 *		arr[0:end] <=> arr[:end]
 *		arr[start:len(arr)] <=> arr[start:]
 *		arr[0:len(arr)] <=> arr[:]
 *	3、cap 是一个内置函数，用于统计切片的容量，即最大可以存放多少个元素
 *	4、切片定义后还不能使用，因为本身是一个空的，需要将其引用到一个数组或者 make 一个空间供切片来使用
 *	5、切片可以继续切片
 */
func main()  {
	var arr = [5]int{1, 2, 3, 4, 5}
	slice1 := arr[1:4]
	slice2 := slice1[1:2]
	fmt.Println("slice1=", slice1)
	fmt.Println("slice2=", slice2)
}