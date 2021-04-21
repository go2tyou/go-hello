package main

import (
	"fmt"
)

// 嵌套分支
// 建议控制在 3 层内
func main() {
	// var score float64
	// var gender string

	// fmt.Println("input please:")
	// fmt.Println("成绩:")
	// fmt.Scanln(&score)
	// fmt.Println("性别:")
	// fmt.Scanln(&gender)
	// if score <= 8 {
	// 	if gender == "man" {
	// 		fmt.Println("man in")
	// 	} else {
	// 		fmt.Println("woman in")
	// 	}
	// } else {
	// 	fmt.Println("out")
	// }

	var month byte
	var age byte
	var price float64 = 60
	fmt.Println("input please:")
	fmt.Println("月份:")
	fmt.Scanln(&month)
	fmt.Println("年龄:")
	fmt.Scanln(&age)
	if month >= 4 && month <= 10 {
		if age > 60 {
			fmt.Printf("价格：%v", price / 3)
		} else if age >= 18 && age <= 60 {
			fmt.Printf("价格：%v", price)
		} else {
			fmt.Printf("价格：%v", price / 2)
		}
	} else {
		if age >= 18 && age < 60 {
			fmt.Printf("价格: 40")
		} else {
			fmt.Printf("价格: 20")
		}		
	}
}