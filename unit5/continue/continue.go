package main

import (
	"fmt"
)

// 用于结束本次循环，继续执行下一次循环
// 支持标签，同 break，指明要跳过的是哪个循环
func main() {
	
	// label2:
	for i := 0; i < 4; i++ {
		// label1: // 设置一个标签，名字可自定义
		for j := 0; j < 10; j++ {
			if j == 2 {
				continue
				// continue label1
				// continue label2
			}
			fmt.Println("j=", j)
		}
	}
}