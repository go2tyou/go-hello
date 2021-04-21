package main

import (
	"fmt"
)

func test(n int)  {
	if n > 2 {
		n--
		test(n)
	}
	fmt.Println("n=", n)
}

func test2(n int)  {
	if n > 2 {
		n--
		test2(n)
	} else {
		fmt.Println("n2=", n)
	}
}

/*
 *	函数递归原则：
 *		1、执行一个函数时，就创建一个受保护的独立函数（新函数栈）
 *		2、函数的局部变量是独立的，不会相互影响
 *		3、递归必须向退出递归条件逼近，否则就是无限循环调用
 *		4、当一个函数执行完毕或者遇到 return，就会返回，遵守谁调用就将结果返回给谁。同时，该函数本身也会被系统销毁
 */
func main()  {
	test(4)
	test2(4)
}