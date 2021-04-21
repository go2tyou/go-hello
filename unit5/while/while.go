package main

import (
	"fmt"
)

func main() {
	// go 中类似 while 实现
	var i int = 0
	for {
		if i >= 10 {
			break
		}
		fmt.Println("gg", i)
		i++
	}

	// go 中类似 do...while 实现
	// 至少执行一次
	var j int = 0
	for {
		fmt.Println("cc", j)
		j++
		if j >= 0 {
			break
		}
	}
}