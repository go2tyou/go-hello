package main

import (
	"fmt"
)

type MethodUtils struct {
	// 结构体可以不要字段
}

func (m MethodUtils) print() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 8; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func (mu MethodUtils) print2(m int, n int) {
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}	
}

func (mu MethodUtils) area(len float64, width float64) float64 {
	return len * width
}

func main()  {
	var mu MethodUtils
	mu.print()
	fmt.Println()
	mu.print2(4, 6)
	res := mu.area(4.0, 2.7)
	fmt.Println("res=", res) // res= 10.8
}