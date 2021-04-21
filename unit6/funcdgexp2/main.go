package main

import (
	"fmt"
)

// 猴子吃桃
func f(n int) int {
	if n > 10 || n < 1 {
		return 0
	}
	if n == 10 {
		return 1
	} else {
		return (f(n + 1) + 1) * 2
	}
}

func main()  {
	fmt.Println("f=", f(1))
}