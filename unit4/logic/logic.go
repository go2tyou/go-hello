package main

import (
	"fmt"
)

// 逻辑运算符
// && || !
func main()  {
	var age int = 40

	if age > 30 && age < 50 {
		fmt.Println("ok1")
	}
	if age > 30 && age < 40 {
		fmt.Println("ok2")
	}

	if age > 30 || age < 50 {
		fmt.Println("ok3")
	}
	if age > 30 || age < 40 {
		fmt.Println("ok4")
	}

	if age > 30 {
		fmt.Println("ok5")
	}
	if !(age > 30) {
		fmt.Println("ok6")
	}

}