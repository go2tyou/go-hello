package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 数组反转打印
func main()  {
	var intArr [5]int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(intArr); i++ {
		intArr[i] = rand.Intn(100) // [0, 100)
	}
	fmt.Println("intArr=", intArr)
	
	temp := 0
	len := len(intArr)
	for i := 0; i < len / 2; i++ {
		temp = intArr[len - 1 - i]
		intArr[len - 1 - i] = intArr[i]
		intArr[i] = temp
	}
	fmt.Println("intArr=", intArr)
}