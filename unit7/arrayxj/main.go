package main

import (
	"fmt"
)

// 长度是数组类型的一部分
// [3]int [4]int 不是一种数据类型
func test8(arr [3]int) {
	arr[0] = 88
}

func test9(arr *[3]int) {
	(*arr)[0] = 88
}

/*
	使用细节：
	1、数组是多个相同类型的数据，一个数组一旦定义，其长度是固定的，不能动态变化
	2、var arr []int 这是 arr 就是一个 slice 切片
	3、数组的元素可以是任何数据类型，但不能混用
	4、数组创建后，如果没有赋值，有默认值（零值）	数值数组：0		字符串数组：""		布尔数组：false	
	5、使用数组步骤：声明数组并开辟空间 -> 给数组各元素赋值（默认零值）-> 使用数组
	6、数组的下标是从 0 开始的
	7、数组下标必须在指定范围内使用，否则报 panic，数组越界
	8、go 的数组属值类型，在默认情况下是值传递，因此会进行值拷贝，数组间不会相互影响
	9、如果想在其他函数中，修改原来的数组，可以使用引用传递（指针方式）
	10、长度是数组类型的一部分，在传递函数参数时，需要考虑数组的长度
*/
func main()  {
	// 测试 1
	var arr1 [3]int
	arr1[0] = 1
	arr1[1] = 2
	// arr1[2] = 3.2 // constant 3.2 truncated to integer
	arr1[2] = 3
	// arr1[3] = 4 // invalid array index 3 (out of bounds for 3-element array)
	fmt.Println("arr1=", arr1)

	// 测试 4
	var arr4_1 [3]int32
	var arr4_2 [3]string
	var arr4_3 [3]float32
	var arr4_4 [3]bool
	fmt.Printf("arr4_1=%v arr4_2=%q arr4_3=%v arr4_4=%v\n", arr4_1, arr4_2, arr4_3, arr4_4) 
	// arr4_1=[0 0 0] arr4_2=["" "" ""] arr4_3=[0 0 0] arr4_4=[false false false]

	// 测试 8
	arr8 := [3]int{11, 22, 33}
	test8(arr8)
	fmt.Println("arr8=", arr8) // arr8= [11 22 33]

	// 测试 9
	arr9 := [3]int{11, 22, 33}
	test9(&arr9)
	fmt.Println("arr9=", arr9) // arr9= [88 22 33]
}