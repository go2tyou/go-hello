package main

import (
	"fmt"
)

// 二维数组的遍历
func main()  {
	arr := [2][3]int{{1, 2, 3}, {4, 5, 6}}

	// for
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("%v\t", arr[i][j])
		}
		fmt.Println()
	}
	
	// for...range
	for i, v1 := range arr {
		for j, v2 := range v1 {
			fmt.Printf("arr[%v][%v]=%v\t", i, j, v2)
		}
		fmt.Println()
	}
}