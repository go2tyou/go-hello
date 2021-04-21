package main

import (
	"fmt"
)

// 双分支
// 只会执行其中的一个分支
func main() {
	var age int
	fmt.Println("input age:")
	fmt.Scanln(&age)

	if age > 18 {
		fmt.Println("ok")
	} else {
		fmt.Println("out")
	}
}
