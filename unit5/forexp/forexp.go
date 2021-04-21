package main

import (
	"fmt"
)

func main() {
	var max int = 100
	var count int = 0
	var sum int = 0
	for i := 1; i <= max; i++ {
		if i % 9 == 0 {
			count++
			sum += i
		}
	}
	fmt.Printf("count=%v sum=%v\n", count, sum)

	var n int = 6
	for i := 0; i <= n; i++ {
		fmt.Printf("%v + %v = %v\n", i, n - i, n)
	}
}