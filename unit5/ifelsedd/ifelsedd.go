package main

import (
	"fmt"
)

// 多分支
// else 不是必须的
// 只能有一个入口
func main() {
	var score int
	fmt.Println("input score:")
	fmt.Scanln(&score)
	if score == 100 {
		fmt.Printf("rufa!")
	} else if score > 80 && score <= 99 {
		fmt.Printf("cc")
	} else if score > 60 && score <= 79 {
		fmt.Printf("ri")
	} else {
		fmt.Printf("out!")
	}
}
