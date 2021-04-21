package main

import (
	"fmt"
)

// switch 使用细节
// 1、case 与 switch  后是一个表达式（常量、变量、有返回值的函数、可以运算）
// 2、case 与 switch 表达式数据类型需保持一致
// 3、case 后面可以带多个表达式
// 4、case 后面的如果是常量值，则要求不能重复（编译器可能检测不出来）
// 5、default 语句不是必须的
// 6、switch 后面可以不带表达式，类似 if-else 使用
// 7、case 中也可以对范围进行判断
// 8、switch 后面也可以直接声明一个变量，分号结束，不推荐
// 9、switch 穿透 fallthrough
// 10、Type switch 用于判断某个 interface 变量中实际指向的变量类型
func main() {

	// // 测试 6
	// var age int = 10
	// switch {
	// case age == 10:
	// 	fmt.Println("age==10")
	// case age == 20:
	// 	fmt.Println("age==20")	
	// }

	// // 测试 9
	// var num int = 10
	// switch num {
	// case 10:
	// 	fmt.Println("ok")	
	// 	fallthrough // 默认只能穿透一层
	// case 20:
	// 	fmt.Println("ok..")	
	// case 30:
	// 	fmt.Println("ok.....")	
	// default:
	// 	fmt.Println("error")	
	// }

	// var char byte
	// fmt.Println("input please:")
	// fmt.Scanf("%c", &char)
	// switch char {
	// case 'a':
	// 	fmt.Println("A")	
	// case 'b':
	// 	fmt.Println("B")	
	// case 'c':
	// 	fmt.Println("C")	
	// default:	
	// 	fmt.Println("other")	
	// }

	// var score float64
	// fmt.Println("input please:")
	// fmt.Scanln(&score)
	// switch int(score / 60) {
	// case 1:
	// 	fmt.Println("nice")	
	// case 0:
	// 	fmt.Println("no")	
	// default:
	// 	fmt.Println("error")	
	// }

	var month byte
	fmt.Println("input please:")
	fmt.Scanln(&month)
	switch month {
	case 3,4,5 :
		fmt.Println("haru")	
	case 6,7,8:	
		fmt.Println("natsu")	
	case 9,10,11:
		fmt.Println("aki")	
	case 12,1,2:
		fmt.Println("huyu")
	default:
		fmt.Println("error")		
	}
}
