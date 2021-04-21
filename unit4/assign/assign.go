package main

import (
	"fmt"
)

// 赋值运算符
// = += -= *= /= %= <<= >>= &= ^= |=
// 1、运算顺序从右至左
// 2、左边只能是变量，右边可以是变量，也可以是表达式（任何有值的都可以看做是表达式）

// 运算符优先级
// 1、括号，++，--
// 2、单目（右到左）
// 3、算术
// 4、移位
// 5、关系
// 6、位
// 7、逻辑
// 8、赋值（右到左）
// 9、逗号
func main()  {
	var i int
	i = 10 // 基本赋值
	fmt.Println(i)
	
	// a,b 交换
	a := 9
	b := 2
	fmt.Printf("before: a=%d b=%d\n", a, b)
	t := a
	a = b
	b = t
	fmt.Printf("after: a=%d b=%d\n", a, b)
		
	// 复合赋值
	a += 1 // <=> a = a + 1
	fmt.Println("a=", a)

	var x int = 10
	var y int = 20
	fmt.Printf("before: x=%d y=%d\n", x, y)
	x = x + y
	y = x - y // y = x + y - y = x
	x = x - y // x = x + y - x = y
	fmt.Printf("after: x=%d y=%d\n", x, y)

}