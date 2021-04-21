package main

import "fmt"

/**
 * 1、范围大 <=> 范围小
 * 2、被转换的是变量存储的数据（即值），变量本身的数据类型并没有变化
 * 3、int64 -> int8，编译时不会报错，只是转换结果按溢出处理（转换时要考虑范围）
 */

// go 中基本数据类型的转换
func main() {

	var i int32 = 100
	// i => float
	var j float32 = float32(i)
	var k int8 = int8(i)
	// var h int64 = i // cannot use i (type int32) as type int64 in assignment
	var h int64 = int64(i)
	fmt.Printf("i=%v j=%v k=%v h=%v\n", i, j, k, h)

	fmt.Printf("i typeof is %T\n", i)

	var num1 int64 = 9999999
	var num2 int8 = int8(num1)
	fmt.Println("num2=", num2) // num2= 127 和预期值不一致
}
