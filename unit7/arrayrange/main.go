package main

import (
	"fmt"
)

// for index, value := range array() {...}
/*
 *	1、index 是数组的下标
 *	2、value 是在该下标位置的值
 *	3、index、value 仅在 for 循环内部可见的局部变量
 *	4、不想使用下标，可以使用 _
 *	5、index、value 的名称不是固定的，可自定义
 */
func main()  {
	var heroes = [3]string{"怂", "孬", "猥"}	
	for index, value := range heroes {
		fmt.Printf("index=%v value=%v\n", index, value)
	}
}