package main

import (
	"fmt"
)

// 二维数组
/*
 *	1、var 数组名 [大小][大小]类型
 *	2、var 数组名 [大小][大小]类型 = [大小][大小]类型{{初值...}, {初值...}, ...}
 *	3、var 数组名 [大小][大小]类型 = [...][大小]类型{{初值...}, {初值...}, ...}
 *	4、var 数组名 = [大小][大小]类型{{初值...}, {初值...}, ...}
 *	5、var 数组名 = [...][大小]类型{{初值...}, {初值...}, ...}
 */
func main()  {
	var arr [4][6]int
	fmt.Println("arr=", arr)
	arr[1][2] = 1
	arr[2][1] = 2
	arr[2][3] = 3
	fmt.Println("arr=", arr)
	
	// 遍历
	for i := 0; i < 4; i++ {
		// fmt.Printf("arr[%v]=%v\n", i, arr[i])
		for j := 0; j < 6; j++ {
			fmt.Print(arr[i][j])
		}
		fmt.Println()
	}

	var arr1 [2][3]int
	fmt.Printf("arr1[0]=%p arr1[1]=%p\n", &arr1[0], &arr1[1]) // arr1[0]=0xc000088000 arr1[1]=0xc000088018
	fmt.Printf("arr1[0][0]=%p arr1[1][0]=%p\n", &arr1[0][0], &arr1[1][0]) // arr1[0][0]=0xc000088000 arr1[1][0]=0xc000088018

	var arr2 [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("arr2=", arr2)
}