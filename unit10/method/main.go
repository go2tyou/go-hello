package main

import (
	"fmt"
)

type A struct {
	Num int
}

// 给 A 类型绑定一个方法
// test() 只能通过 A 类型的变量来调用，而不能直接调用，也不能通过其他类型来调用
// a 变量表示哪个 A 变量调用，这个 a 就是它的副本
// a 这个名字可自定义
func (a A) test() {
	fmt.Println("type A test()...Num=", a.Num)
}

// 方法
/*
 *	1、go 中的方法是作用在指定的数据类型上（即和指定的数据类型绑定），因此，自定义类型都可以有方法，而不仅仅是 struct
 */
func main()  {
	var a A
	a.Num = 10
	// 调用方法
	a.test() // type A test()...Num= 10
}