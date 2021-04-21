package main

import (
	"fmt"
)

// 1、每一个 case 分支都是唯一的，从上到下逐一测试，直到匹配为止
// 2、匹配项后不需要加 break，默认有，即在默认情况下，当程序执行完 case 语句块后，就会直接退出该 swith 控制结构
// 3、case 后的表达式可以有多个，使用逗号间隔
func main() {
	fmt.Println("input please:")
	var c byte
	fmt.Scanf("%c", &c)
	switch c {
	case 'a':
		fmt.Println("one")
	case 'b':
		fmt.Println("two")
	default:
		fmt.Println("error")
	}
}
