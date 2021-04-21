package main

import (
	"fmt"
)

// 九九乘法表
func main() {

	// i 表示层数
	for i := 1; i <= 9; i++ {
		// j 表示表达式个数
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v = %v\t", j, i, i * j)
		}
		fmt.Println()
	}
}