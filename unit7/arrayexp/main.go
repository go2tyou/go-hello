package main

import (
	"fmt"
)

func main() {
	// var score [5]float64
	// for i := 0; i < len(score); i++ {
	// 	fmt.Printf("input %d=\n", i + 1)
	// 	fmt.Scanln(&score[i])
	// }
	// for i := 0; i < len(score); i++ {
	// 	fmt.Printf("score[%d]=%v\n", i, score[i])
	// }

	// 四种数组初始化方式
	var numArr01 [3]int = [3]int{1, 2, 3}
	fmt.Printf("numArr01=%v\n", numArr01)

	var numArr02 = [3]int{1, 2, 3}
	fmt.Printf("numArr02=%v\n", numArr02)

	var numArr03 = [...]int{1, 2, 3}
	fmt.Printf("numArr03=%v\n", numArr03)

	var numArr04 = [...]int{1: 100, 2: 200, 0: 300}
	fmt.Printf("numArr04=%v\n", numArr04)

	numArr05 := [...]int{1: 100, 2: 200, 0: 300}
	fmt.Printf("numArr05=%v\n", numArr05)
}
