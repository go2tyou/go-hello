package main

import (
	"fmt"
)

// 一个函数 test
func test(n1 int)  {
	n1 = n1 + 1
	fmt.Println("test n1=", n1)
}

/*
 *	return 语句：
 *		1、go 函数支持返回多个值
 *			func 函数名 (形参列表) (返回值类型列表) {
 *				语句...
 *				return 返回值列表	
 *			}
 *		2、如果返回多个值，在接收时希望忽略某个返回值，则使用下划线 _ 符号表示占位忽略	
 *		3、	如果返回值只有一个，(返回值类型列表) 可以不写 ()
 */

// 一个函数 getSum
func getSum(n1 int, n2 int) int {
	sum := n1 + n2
	fmt.Println("getSum sum=", sum)
	// 当函数有 return 时，就将结果返回给调用者
	// 谁调用我，返回给谁
	return sum
}

func getSumAndSub(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}

/*
 *	1、在调用一个函数时，会给该函数分配一个新的空间
 *	2、在每个函数对应的栈中，数据空间是独立的，不会混淆
 *	3、当一个函数执行完毕后，程序会销毁这个函数对应的栈空间
 */
func main()  {
	n1 := 10
	// 调用 test
	test(n1)
	fmt.Println("main n1=", n1)

	sum := getSum(1, 2)
	fmt.Println("main sum=", sum)

	r1, r2 := getSumAndSub(2, 1)
	fmt.Printf("r1=%v r2=%v\n", r1, r2)

	r3, _ := getSumAndSub(3, 5)
	fmt.Println("r3=", r3)
}