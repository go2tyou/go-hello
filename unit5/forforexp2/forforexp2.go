package main

import (
	"fmt"
)

// 空心金字塔
func main() {

	var totalLevel int = 9

	// i 表示层数
	for i := 1; i <= totalLevel; i++ {
		for k := 1; k <= totalLevel - i; k++ {
			fmt.Print(" ")
		}
		// j 表示每层打印 * 的个数
		for j := 1; j <= 2 * i - 1; j++ {
			if j == 1 || j == 2 * i - 1 || i == totalLevel {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}