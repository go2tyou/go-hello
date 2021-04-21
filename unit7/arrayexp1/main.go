package main

import (
	"fmt"
)

func main()  {
	// 字符序列
	var myChar [26]byte
	for i := 0; i < 26; i++ {
		myChar[i] = 'A' + byte(i)
	}
	for i := 0; i < 26; i++ {
		fmt.Printf("%c ", myChar[i])
	}
	fmt.Println()

	// 一个数组最大值
	var intArr = [...]int{1, -1, 99, 8, 9}
	maxVal := intArr[0]
	maxIndex := 0
	for i := 0; i < len(intArr); i++ {
		if maxVal < intArr[i] {
			maxVal = intArr[i]
			maxIndex = i
		}
	}
	fmt.Printf("maxVal=%v maxIndex=%v\n", maxVal, maxIndex)
	
	// 求一个数组的和和平均值
	var intArr2 = [...]int{1, -1, 99, 8, 9, 7, 5}
	sum := 0
	for _, val := range intArr2 {
		sum += val
	}
	fmt.Printf("sum=%v avg=%v\n", sum, float64(sum) / float64(len(intArr2)))
}