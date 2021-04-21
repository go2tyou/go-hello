package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64 = 3.0, 10.0, 66.0
	m := b * b - 4 * a * c
	if m > 0 {
		x1 := (-b + math.Sqrt(m) / 2 * a)
		x2 := (-b - math.Sqrt(m) / 2 * a)
		fmt.Printf("x1=%v x2=%v", x1, x2)
	} else if m == 0 {
		x := (-b + math.Sqrt(m) / 2 * a)
		fmt.Printf("x=%v", x)
	} else {
		fmt.Println("无解。。。")
	}

	var height int32
	var money float32
	var handsome bool
	fmt.Println("input please:") 
	fmt.Println("白:") 
	fmt.Scanln(&height)
	fmt.Println("富:") 
	fmt.Scanln(&money)
	fmt.Println("美:") 
	fmt.Scanln(&handsome)
	if height > 180 && money > 1.0 && handsome {
		fmt.Println("!!!")
	} else if height > 180 || money > 1.0 || handsome {
		fmt.Println("!")
	} else {
		fmt.Println("v_v")
	}

}
