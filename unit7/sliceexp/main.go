package main

import (
	"fmt"
)

// 切片的使用
// 1、通过 make 方式创建切片可以指定切片的大小和容量
// 2、如果没有给切片各个元素赋值，那么就会使用默认值
// 3、通过 make 方式创建的切片对应的数组是由 make 底层维护，对外不可见，即只能通过 slice 访问元素
func main()  {
	// 方式一
	// 定义一个切片，然后让切片去引用一个已经创建好的数组

	// 方式二
	// var 切片名 []type = make([]type, len, [cap])
	// cap 可选，如果写了必须大于 len
	// 对于切片，必须 make 后才能使用
	var slice []float64 = make([]float64, 5, 10)
	fmt.Println("slice e=", slice)
	fmt.Println("slice len=", len(slice))
	fmt.Println("slice cap=", cap(slice))

	// 方式三
	// 定义一个切片，直接就指定具体数组，使用原理类似 make 的方式
	var slice3 []string = []string{"ddd", "fff", "rrr"}
	fmt.Println("slice3 e=", slice3)
	fmt.Println("slice3 len=", len(slice3))
	fmt.Println("slice3 cap=", cap(slice3))

	// 遍历方式一
	var arr = [5]int{1, 2, 3, 4, 5}
	s1 := arr[1:4]
	for i := 0; i < len(s1); i++ {
		fmt.Printf("s1[%v]=%v ", i, s1[i])
	}
	fmt.Println()

	// 遍历方式二
	for i, v := range s1 {
		fmt.Printf("i=%v v=%v\t", i, v)
	}
}