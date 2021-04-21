package main

import (
	"fmt"
)

// 二分查找
func BinaryFind(arr *[6]int, left int, right int, findVal int) {
	if left > right {
		fmt.Println("no...")
		return
	}
	mid := (left + right) / 2
	if (*arr)[mid] > findVal {
		// left, mid - 1
		BinaryFind(arr, left, mid - 1, findVal)
	} else if (*arr)[mid] < findVal {
		// mid + 1, right
		BinaryFind(arr, mid + 1, right, findVal)
	} else {
		// get
		fmt.Printf("get...%v index=%v\n", findVal, mid)
	}
}

func main()  {
	arr := [6]int{1, 8, 10, 89, 1234}
	BinaryFind(&arr, 0, len(arr) - 1, 10)
}