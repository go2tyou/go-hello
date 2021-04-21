package main

import (
	"fmt"
)

type Cat struct {
	Name string
	Age int
	Color string
}

// 结构体
/*
 *	1、将一类事物的特征提取出来，形成一个新的数据类型，就是一个结构体
 *	2、通过这个结构体，可以创建多个变量（实例对象）
 *	3、事物可以是。。。
 */
/*
 *	1、结构体是自定义的数据类型，代表一类事物
 *	2、结构体变量（实例）是具体的，实际的，代表一个具体变量
 */
 /*
	type 结构体名称 struct {
		field1 type
		field2 type
	}
 */
func main()  {
	var cat1 Cat
	fmt.Println("cat1=", cat1) // cat1= { 0 }
	fmt.Printf("cat1=%p\n", &cat1) // cat1=0xc00006e330 直接指向，说明为值类型
	cat1.Name = "小白"
	cat1.Age = 3
	cat1.Color = "白色"
	fmt.Println("cat1=", cat1) // cat1= {小白 3 白色}
}