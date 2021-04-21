package main

import (
	"fmt"
)

// 算数运算符
func main()  {
	// 在 go 中，++ -- 只能独立使用
	var i int = 8
	var a int
	// a = i++ // syntax error: unexpected ++ at end of statement
	i++
	a = i
	// syntax error: unexpected >, expecting {
    // if i++ > 0 {
	// }
	fmt.Println(a)

	// 只有 a++ a--, 没有 ++a --a
	// ++i // syntax error: unexpected ++, expecting }

}