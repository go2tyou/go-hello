package main

import (
	"fmt"
)

type Point struct {
	x int
	y int
}

// 类型断言
// 由于接口是一般类型，不知道具体类型，如果要转成具体类型，就需要使用类型断言
func main()  {
	var a interface{}
	var point Point = Point{1, 2}
	a = point
	// 如何将 a 赋值给一个 Point 变量
	var b Point
	// b = a // cannot use a (type interface {}) as type Point in assignment: need type assertion
	b = a.(Point)
	fmt.Println(b) // {1 2}

	// 在进行类型断言时，如果类型不匹配，就会报 panic
	// 因此，要确保原来空接口指向的就是断言的类型
	var x interface{}
	var c float32 = 1.1
	x = c
	y := x.(float32)
	// y := x.(float64) // panic: interface conversion: interface {} is float32, not float64
	fmt.Printf("type y %T val=%v\n", y, y) // type y float32 val=1.1

	// 如何在进行类型断言时，带上检测机制，如果成功 ok，否则也不要报 panic
	var p interface{}
	var q float32 = 2.3
	p = q
	if res, ok := p.(float64); ok {
		fmt.Println("convert success")
		fmt.Printf("type res %T val=%v\n", res, res)
	} else {
		fmt.Println("convert fail")
	}
	fmt.Println("go go go...")
}