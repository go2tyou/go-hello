package main

import (
	"fmt"
)

// 算数运算符
func main()  {
	var days int = 97
	var week int = days / 7
	var day int = days % 7
	fmt.Printf("%d week, %d day\n", week, day)

	var h float32 = 134.2
	var s float32 = 5.0 / 9 * (h - 100)
	fmt.Printf("s=%f\n", s)
}