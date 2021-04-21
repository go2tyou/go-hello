package main

import (
	"fmt"
)

// 斐波那契数列
func fb(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return fb(n - 1) + fb(n - 2)
	}
}

func f(n int) int {
	if n == 1 {
		return 3
	} else {
		return 2 * f(n - 1) + 1
	}
}

func main()  {
	r := fb(3)
	fmt.Println("r=", r)
	
	fmt.Println("f=", f(1))
	fmt.Println("f=", f(5))
}