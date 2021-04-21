package main

import (
	"fmt"
)

// 数组可以存放多个同一类型数据
// 在 go 中，数组是值类型

// 易维护、易扩展

// 定义：  var 数组名 [数组大小]数据类型
// 访问：  数组名[下标]

/*
 *	1、数组的地址可以通过数组名来获取 =%p &intArr
 *	2、数组的第一个元素的地址，就是数组的首地址 &intArr == &intArr[0]
 *	3、数组的各元素的地址间隔是按照数组的类型决定的 int64 -> 8 int32 -> 4
 */
func main()  {
	// 使用数组的方式
	var hens [6]float64 // 定义数组后，数组各个元素有默认值
	hens[0] = 1.0
	hens[1] = 2.0
	hens[2] = 4.0
	hens[3] = 3.4
	hens[4] = 5.0
	hens[5] = 10.0
	totalw := 0.0
	for i := 0; i < len(hens); i++ {
		totalw += hens[i]
	}
	avgw := fmt.Sprintf("%.2f", totalw / float64(len(hens)))
	fmt.Printf("totalw=%v avgw=%v\n", totalw, avgw)
}