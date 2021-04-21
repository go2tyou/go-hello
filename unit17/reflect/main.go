package main

import (
	"reflect"
	"fmt"
)

func test1(b interface{})  {
	// 1、先获取 reflect.Type
	rType := reflect.TypeOf(b)
	fmt.Println("rType=", rType) // rType= int

	// 2、获取到 reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Println("rVal=", rVal) // rVal= 100
	// num := 2 + rVal // invalid operation: 2 + rVal (mismatched types int and reflect.Value)
	num := 2 + rVal.Int()
	fmt.Println("num=", num) // num= 102

	// 3、将 rVal -> interface{}
	iV := rVal.Interface()
	// 将 interface{} 通过类型断言转成需要的类型
	num2 := iV.(int)
	fmt.Println("num2=", num2) // num2= 100
}

type Student struct {
	Name string
	Age int
}

func test2(b interface{}) {
	rType := reflect.TypeOf(b)
	fmt.Println("struct rType=", rType) // struct rType= main.Student

	rVal := reflect.ValueOf(b)
	fmt.Println("struct rVal=", rVal) // struct rVal= {tom 20}

	// 获取变量对应的 kind -> 本质是一个常量
	// 1、rVal.Kind()
	kind1 := rVal.Kind()
	// 2、rType.Kind()
	kind2 := rType.Kind()
	fmt.Printf("kind1=%v kind2=%v\n", kind1, kind2) // kind1=struct kind2=struct

	iV := rVal.Interface()
	fmt.Printf("struct iV=%v iV=%T\n", iV, iV) // struct iV={tom 20} iV=main.Student

	// 还有 switch 形式断言
	stu, ok := iV.(Student)
	if ok {
		fmt.Println("stu.Name=", stu.Name) // stu.Name= tom
	}
}

/*
 *	1、反射可以在 运行时!!! 动态获取变量的各种信息
 *	2、通过反射，可以修改变量的值，可以调用关联的方法
 *	3、import "reflect"
 *	4、变量、interface{}、reflect.Value 可以相互转换
 */
func main()  {
	var num int = 100
	test1(num)

	stu := Student {
		Name : "tom",
		Age : 20,
	}
	test2(stu)
}