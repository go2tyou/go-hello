package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age int
}

func main()  {
	// 方式一

	// 方式二
	// p2 := Person{}
	// p2.Name = "tom"
	// p2.Age = 10
	p2 := Person{"marry", 20}
	fmt.Println("p2=", p2)

	// 方式三
	var p3 *Person = new(Person)
	(*p3).Name = "smith" // <=> p3.Name = "smith" 
	p3.Name = "john" 
	// -> go 的设计者，为了程序员使用方便，底层会对 p3.Name = "smith" 进行处理
	// -> 会给 p3 加上取值运算 (*p3).Name = "smith"
	(*p3).Age = 17
	fmt.Println("p3=", *p3) // p3= {john 17}

	// 方式四
	var p4 *Person = &Person{} // 可以直接赋初值 -> var p4 *Person = &Person{"marry", 20}
	// 因为 p4 是一个指针，因此标准的访问字段的方法
	// (*p4).Name = "scott"
	// -> go 的设计者，为了程序员使用方便，底层会对 p4.Name = "scott" 进行处理，会加上 (*p4)
	(*p4).Name = "scott"
	(*p4).Age = 99
	p4.Age = 9
	fmt.Println("p4=", *p4) // p4= {scott 9}
}