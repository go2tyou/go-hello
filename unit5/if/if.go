package main

import (
	"fmt"
)

// 单分支
func main() {
	// var age int
	// fmt.Println("input age:")
	// fmt.Scanln(&age)

	// {} 必须有
	// if age > 18 {
	// 	fmt.Println("ok")
	// }

	// 支持直接在 if 定义一个变量
	if age := 20; age > 18 {
		fmt.Println("ok")
	}
	
}