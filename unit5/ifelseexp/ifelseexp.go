package main

import (
	"fmt"
)

func main() {

	var n1 int32 = 11
	var n2 int32 = 12
	if n1 + n2 > 20 {
		fmt.Println("^_^")
	}

	var f1 float64 = 12.121
	var f2 float64 = 13.14
	if f1 > 10.0 && f2 < 20.0 {
		fmt.Printf("sum=%f\n", (f1 + f2))
	}

	var m int32 = 66
	var n int32 = 99
	if (m + n) % 3 == 0 && (m + n) % 5 == 0 {
		fmt.Println("ok~")
	} 

	var year int = 2020
	if (year % 4 == 0 && year % 100 != 0) || year % 400 == 0 {
		fmt.Println(year, "是闰年！")
	}
}
