package main

import (
	"fmt"
)

// 冒泡排序
func BubbleSort(arr *[5]int)  {
	temp := 0
	for i := 0; i < len(arr) - 1; i++ {
		for j := 0; j < len(arr) - 1 - i; j++ {
			if (*arr)[j] > (*arr)[j + 1] {
				temp = (*arr)[j]
				(*arr)[j] = (*arr)[j + 1]
				(*arr)[j + 1] = temp
			}
		}	
	}
}

func main()  {
	arr := [5]int{24, 69, 80, 57, 13}
	fmt.Println("arr sort be=", arr)
	BubbleSort(&arr)
	fmt.Println("arr sort af=", arr)
}