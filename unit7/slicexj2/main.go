package main

import (
	"fmt"
)

// 使用细节
/*
 *	6、用 append 内置函数，可以对切片进行动态追加
 *		6.1、切片 append 操作的本质就是对数组扩容
 *		6.2、go 底层会创建一个新的数组（安装扩容后大小）
 *		6.3、将 slice 原来包含的元素拷贝到新的数组
 *		6.4、slice 重新引用到新的数组
 *		6.5、新的数组在底层来维护，程序员不可见
 *	7、使用内置函数 copy 可以对切片进行拷贝
 *		7.1、copy(p1, p2) 参数的数据类型是切片
 *		7.2、p1、p2的数据空间是独立的，相互不影响
 *		7.3、拷贝到的切片大小不够任然可以拷贝
 *	8、切片是引用类型，所以在传递时，遵循引用传递机制
 */
func main()  {
	// 切片的追加
	var slice []int = []int{1, 2, 3}
	fmt.Println("slice=", slice)
	slice = append(slice, 4, 5)
	fmt.Println("slice=", slice)
	slice = append(slice, slice...) // 通过 append 将切片追加给切片
	fmt.Println("slice=", slice)

	// 切片的拷贝
	// 使用内置函数 copy
	var a []int = []int{1, 2, 3, 4, 5}
	var s = make([]int, 10)
	fmt.Println("s start=", s)
	copy(s, a)
	fmt.Println("s  copy=", s)
}