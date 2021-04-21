package main

import (
	"fmt"
)

type Person struct {
	Name string
}

func (p Person) getSum(n1 int, n2 int) int {
	return n1 + n2
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

// 为了提高效率，通常方法和结构体的指针类型绑定
func (c *Circle) area2() float64 {
	// 标准访问
	// return 3.14 * (*c).radius * (*c).radius
	return 3.14 * c.radius * c.radius
}

// int float32... 都可以有方法
type integer int

func (i integer) print() {
	fmt.Println("i=", i)
}

func (i *integer) change() {
	*i = *i + 1
}

// 如果一个类型实现了 String() 这个方法，那么 fmt.Println() 就会调用这个变量的 String()
type Student struct {
	Name string
	Age int
}

func (stu *Student) String() string {
	str := fmt.Sprintf("Name=[%v] Age=[%v]", stu.Name, stu.Age)
	return str
}

// 方法的调用和传参机制
/*
 *	1、在通过一个变量去调用方法时，其调用机制和函数一样
 *	2、不一样的地方是，变量调用方法时，该变量本身也会作为一个参数传递到方法（变量是值类型值拷贝，引用类型地址拷贝）
 */
func main()  {
	var p Person
	p.Name = "ssr"
	n1 := 10
	n2 := 20
	res := p.getSum(n1, n2)
	fmt.Println("res=", res) // res= 30

	var c Circle
	c.radius = 4.0
	ar := c.area()
	fmt.Println("ar=", ar) // ar= 50.24

	var c2 Circle
	c2.radius = 5.0
	// 标准写法
	// ar2 := (&c2).area()
	// 编译器支持
	ar2 := c2.area()
	fmt.Println("ar2=", ar2) // ar2= 78.5

	var i integer = 10
	i.print() // i= 10
	i.change()
	fmt.Println("i=", i) // i= 11

	stu := Student {
		Name : "tom",
		Age : 20,
	}
	fmt.Println(stu) // {tom 20}
	// fmt.Println(&stu) // &{tom 20}
	// 实现了 String()
	fmt.Println(&stu) // Name=[tom] Age=[20]
}