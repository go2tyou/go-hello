package main

import (
	"fmt"
)

// 1、函数中的变量是局部的，函数外不生效
func test()  {
	// n1 是 test 函数的局部变量，只能在 test 函数中使用
	// var n1 int = 10
}

// 2、基本数据类型和数组默认都是值传递的。在函数内修改，不会影响到原来的值
func test2(n1 int)  {
	n1 = n1 + 10
	fmt.Println("test2() n1=", n1)
}

// 3、如果希望修改 2 中的值，可以传入变量飞地址 &，函数内以指针的方式操作变量
func test3(n1 *int)  {
	fmt.Printf("test3() n1=%v\n", &n1) // 指针变量本身占一个空间
	*n1 = *n1 + 10
	fmt.Println("test3() n1=", *n1)
}

// 4、go 函数不支持传统的函数重载

// 5、函数也是一种数据类型，可以赋值给一个变量，则该变量就是函数类型的变量了，通过该变量可以对函数调用
func getSum(n1 int, n2 int) int {
	return n1 * n2
}

// 6、函数既然是一种数据类型，那么，函数可以作为形参，并且调用
func myFun(funvar func(int, int) int, num1 int, num2 int) int {
	return funvar(num1, num2)
}

// 7、go 支持自定义数据类型
type myFunType func(int, int) int // 这时 myFunType 就是 func(int, int) int
func myFun2(funvar myFunType, num1 int, num2 int) int {
	return funvar(num1, num2)
}

// 8、支持对函数返回值命名 -> 名字对应上就 Ok，返回参数列表顺序无所谓
func getSumAndSub(n1 int, n2 int) (sum int, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return
}

// 9、使用 _ 标识符作为占位符，忽略返回值

// 10、go 支持可变参数，必须放在形参列表最后
// args 是 slice 切片，通过 args[index] 可以访问到各个值
func sum(n1 int, args... int) int {
	sum := n1
	// 遍历 args
	for i := 0; i < len(args); i++ {
		sum += args[i] // args[0] 表示取出 args 切片的第一个元素值，其他以此类推
	}
	return sum
}

func main()  {
	// 测试 1
	// 这里不能使用 n1，因为是 test 函数的局部变量
	// fmt.Println("n1=", n1) // undefined: n1

	// 测试 2
	// n1 := 20
	// test2(n1)
	// fmt.Println("main() n1=", n1)

	// 测试 3
	// n1 := 20
	// test3(&n1)
	// fmt.Printf("main() n1=%v\n", &n1)
	// fmt.Println("main() n1=", n1)

	// 测试 5
	// a := getSum
	// fmt.Printf("a 的类型 %T\ngetSum 的类型 %T\n", a, getSum)
	// res := a(10, 20) // <=> res := getSum(10, 20)
	// fmt.Println("res=", res)

	// 测试 6
	// res2 := myFun(getSum, 50, 40)
	// fmt.Println("res2=", res2)

	// 测试 7
	// // 在 go 中 myInt 和 int 虽然都是 int 类型，但 go 认为 myInt 和 int 是两个类型
	// type myInt int // 给 int 取了别名
	// var num1 myInt
	// num1 = 40
	// var num2 int
	// num2 = int(num1)
	// // num2 = num1 // cannot use num1 (type myInt) as type int in assignment
	// fmt.Println("num1=", num1)
	// fmt.Println("num2=", num2)
	
	// res3 := myFun2(getSum, 30, 70)
	// fmt.Println("res3=", res3)

	// 测试 8
	// m, n := getSumAndSub(1, 2)
	// fmt.Printf("m=%v,n=%v\n", m, n)

	// 测试 10
	res4 := sum(10, 0, -3, 4, -1, 5)
	fmt.Println("res4=", res4)
}