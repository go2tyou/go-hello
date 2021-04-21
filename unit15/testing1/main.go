package main

import (
	"fmt"
)

// 一个被测试的函数
func addUpper(n int) int {
	res := 0
	for i := 0; i <= n; i++ {
		res += i
	}
	return res
}

func main()  {
	// 传统测试方法 -> 直接在 main() 中看结果 
	// -> 如果要修改，项目正在运行则需要停止项目
	// -> 不利于管理，当测试多个函数或多个模块时，都需要写在 main() 中，不利于管理，思路不清晰
	// -> 引出单元测试 -> testing 测试框架
	res := addUpper(10)
	if res != 55 {
		fmt.Printf("err respect=%v res=%v\n", 55, res)
	} else {
		fmt.Printf("right respect=%v res=%v\n", 55, res)
	}
}