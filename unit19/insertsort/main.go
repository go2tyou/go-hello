package main

import (
	"fmt"
)

// 插入排序
func InsertSort(arr *[5]int) {
	// 完成第一次，给第二个元素找到合适的位置并插入
	insertVal := arr[1]
	insertIndex := 1 - 1
	// 从大到小
	for insertIndex >= 0 && arr[insertIndex] < insertVal {
		arr[insertIndex + 1] = arr[insertIndex] // 数据后移
		insertIndex--
	}
	// 插入
	if insertIndex + 1 != 1 {
		arr[insertIndex + 1] = insertVal
	}
	fmt.Println("第一次插入 =>", *arr)

	insertVal = arr[2]
	insertIndex = 2 - 1
	for insertIndex >= 0 && arr[insertIndex] < insertVal {
		arr[insertIndex + 1] = arr[insertIndex]
		insertIndex--
	}
	if insertIndex + 1 != 2 {
		arr[insertIndex + 1] = insertVal
	}
	fmt.Println("第二次插入 =>", *arr)

	insertVal = arr[3]
	insertIndex = 3 - 1
	for insertIndex >= 0 && arr[insertIndex] < insertVal {
		arr[insertIndex + 1] = arr[insertIndex]
		insertIndex--
	}
	if insertIndex + 1 != 3 {
		arr[insertIndex + 1] = insertVal
	}
	fmt.Println("第三次插入 =>", *arr)

	insertVal = arr[4]
	insertIndex = 4 - 1
	for insertIndex >= 0 && arr[insertIndex] < insertVal {
		arr[insertIndex + 1] = arr[insertIndex]
		insertIndex--
	}
	if insertIndex + 1 != 4 {
		arr[insertIndex + 1] = insertVal
	}
	fmt.Println("第四次插入 =>", *arr)
}

func InsertSort2(arr *[5]int) {

	for i := 1; i < len(arr); i++ {
		insertVal := arr[i]
		insertIndex := i - 1
		for insertIndex >= 0 && arr[insertIndex] < insertVal {
			arr[insertIndex + 1] = arr[insertIndex]
			insertIndex--
		}
		if insertIndex + 1 != i {
			arr[insertIndex + 1] = insertVal
		}
		fmt.Printf("第 %d 次插入 => %v\n", i, *arr)
	}
}

func main()  {
	arr := [5]int{23, 0, 12, 56, 34}
	fmt.Println("main before =>", arr)
	// InsertSort(&arr)
	InsertSort2(&arr)
	fmt.Println("main after =>", arr)
}