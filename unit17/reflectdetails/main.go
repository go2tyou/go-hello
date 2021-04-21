package main

import (
	"fmt"
	"reflect"
)

func test1(b interface{})  {
	rVal := reflect.ValueOf(b)
	fmt.Printf("rVal kind=%v\n", rVal.Kind()) // rVal kind=ptr
	// rVal.SetInt(20) // panic: reflect: reflect.Value.SetInt using unaddressable value
	rVal.Elem().SetInt(20)
}

/*
 *	1、reflect.Value.Kind() 获取变量的类别，返回的是一个常量
 *	2、Type 类型 Kind 类别，可能相同也可能不同
 *	3、变量 <-------> interface{} <-------> reflect.Value
 *	4、使用反射的方式来获取变量的值（并返回对应的类型），要求数据类型匹配
 *	5、通过反射来修改变量，注意当使用 SetXxx() 来设置需要通过对应的指针类型来完成
 *	   这样才能改变传入的变量的值，同时需要使用到 reflect.Value.Elem()
 */
func main()  {
	// 测试 5
	num := 10
	test1(&num)
	fmt.Println(num) // 20

	// // reflect.Value.Elem() 理解
	// n1 := 1
	// ptr *int = &n1
	// n2 := *ptr // reflect.Value.Elem()
}