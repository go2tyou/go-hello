package main

import (
	"fmt"
)

func main()  {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
		if sum > 20 {
			fmt.Printf("大于 20 时，当前 %v\n", i)
			break
		}
	}

	var name string
	var pwd string
	var chance = 3
	for i := 1; i <= 3; i++ {
		fmt.Println("input name:")
		fmt.Scanln(&name)
		fmt.Println("input pwd:")
		fmt.Scanln(&pwd)

		if name == "cc" && pwd == "666" {
			fmt.Println("sucess!")
			break
		} else {
			chance--
			if chance == 0 {
				break
			}
			fmt.Printf("还有 %v 次机会\n", chance)
		}
	}

	if chance == 0 {
		fmt.Printf("没有机会了，出门左拐！\n")
	}
}