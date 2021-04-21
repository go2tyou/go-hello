package main

import (
	"fmt"
)

func main() {

	var kurasu int = 3
	var gakuse int = 5
	var totalSum float64 = 0.0
	var passCount int = 0
	for j := 1; j <= kurasu; j++ {
		sum := 0.0
		for i := 1; i <= gakuse; i++ {
			var score float64
			fmt.Printf("第 %d 班 第 %d 个:\n", j, i)
			fmt.Scanln(&score)
			sum += score
			if score >= 60 {
				passCount++
			}
		}
		fmt.Printf("第 %d 班 总：%v 平均：%v\n", j, sum, sum / float64(gakuse))
		totalSum += sum
	}
	fmt.Printf("总总：%v 总总平均：%v\n", totalSum, totalSum / float64((kurasu * gakuse)))
	fmt.Printf("及格：%v\n", passCount)
}