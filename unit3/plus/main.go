package main

import "fmt"

// go 中 + 的使用
func main() {

	var i, j = 1, 2
	var k = i + j
	fmt.Println("k=", k)
	
	var m, n = "hello", "world"
	var str = m + n
	fmt.Println("str=", str)
}
