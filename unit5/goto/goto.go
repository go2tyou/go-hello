package main

import (
	"fmt"
)

// 一般不主张使用 goto 语句
func main()  {

	var n int = 1
	fmt.Println("ok1")
	if n < 7 {
		goto label
	}
	fmt.Println("ok2")
	fmt.Println("ok3")
	fmt.Println("ok4")
	fmt.Println("ok5")
	fmt.Println("ok6")
	label:
	fmt.Println("ok7")
}