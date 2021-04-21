package main

import (
	"fmt"
)

// 跳出所在的方法或函数
// 如果在 main() 直接终止
func main()  {
	var n int = 1
	fmt.Println("ok1")
	if n < 3 {
		return
	}
	fmt.Println("ok2")
}