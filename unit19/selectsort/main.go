package main

import (
	"fmt"
)

// 选择排序
func SelectSort(arr *[5]int) {
	// (*arr)[1] = 11
	// arr[1] = 12

	// 1、先完成将第一个最大值和 arr[0] 交换
	// 假设 arr[0] = max
	max := arr[0]
	maxIndex := 0
	// 2、遍历后面 [1, len(arr) - 1] 比较
	for i := 0 + 1; i < len(arr); i++ {
		// 找到真正的最大值
		if max < arr[i] {
			max = arr[i]
			maxIndex = i
		}
	}
	// 3、交换
	if maxIndex != 0 {
		arr[0], arr[maxIndex] = arr[maxIndex], arr[0]
	}
	fmt.Println("第一次 =>", *arr)

	max = arr[1]
	maxIndex = 1
	for i := 1 + 1; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
			maxIndex = i
		}
	}
	if maxIndex != 1 {
		arr[1], arr[maxIndex] = arr[maxIndex], arr[1]
	}
	fmt.Println("第二次 =>", *arr)

	max = arr[2]
	maxIndex = 2
	for i := 2 + 1; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
			maxIndex = i
		}
	}
	if maxIndex != 2 {
		arr[2], arr[maxIndex] = arr[maxIndex], arr[2]
	}
	fmt.Println("第三次 =>", *arr)

	max = arr[3]
	maxIndex = 3
	for i := 3 + 1; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
			maxIndex = i
		}
	}
	if maxIndex != 3 {
		arr[3], arr[maxIndex] = arr[maxIndex], arr[3]
	}
	fmt.Println("第四次 =>", *arr)
}

func SelectSort2(arr *[5]int) {
	// 1、先完成将第一个最大值和 arr[0] 交换
	// 假设 arr[0] = max
	for j := 0; j < len(arr) - 1; j++ {
		max := arr[j]
		maxIndex := j
		// 2、遍历后面 [1, len(arr) - 1] 比较
		for i := j + 1; i < len(arr); i++ {
			// 找到真正的最大值
			if max < arr[i] {
				max = arr[i]
				maxIndex = i
			}
		}
		// 3、交换
		if maxIndex != j {
			arr[j], arr[maxIndex] = arr[maxIndex], arr[j]
		}
		fmt.Printf("第%d次 => %v\n", j + 1, *arr)
	}
}

func main()  {
	arr := [5]int{10, 34, 19, 100, 80}
	fmt.Println("main before =>", arr)
	// SelectSort(&arr)
	SelectSort2(&arr)
	fmt.Println("main after =>", arr)
}