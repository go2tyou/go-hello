package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 100; i++ {
		if i % 2 == 0 {
			continue
		} else {
			fmt.Printf("奇数是 %v\t", i)
		}
	}
	fmt.Println()

	var positiveCount int
	var negativeCount int
	var num int
	for {
		fmt.Println("input please:")	
		fmt.Scanln(&num)
		if num == 0 {
			break
		}
		if num > 0 {
			positiveCount++
			continue
		}
		negativeCount++
	}
	fmt.Printf("正数 %v 个 负数 %v 个", positiveCount, negativeCount)
}