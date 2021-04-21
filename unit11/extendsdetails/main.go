package main

import (
	"fmt"
)

type A struct {
	Name string
	age int
}

func (a * A) SayOk() {
	fmt.Println("A SayOk...", a.Name)
}

func (a * A) hello() {
	fmt.Println("A hello...", a.Name)
}

type B struct {
	A
	Name string
}

func (b * B) SayOk() {
	fmt.Println("B SayOk...", b.Name)
}

type Goods struct {
	Name string
	Price float64
}

type Brand struct {
	Name string
	Address string
}

type TV struct {
	Goods
	Brand
}

type TV2 struct {
	*Goods
	*Brand
}

type Monster struct {
	Name string
	Age int
}

type E struct {
	Monster
	int // 匿名字段是基本数据类型，此类型的匿名字段有且仅有一个
	n int
}

/*
 *	1、结构体可以使用匿名结构体的所有的字段和方法，即与首字母大小写无关
 *	2、匿名结构体访问可以简化
 *	3、当结构体和匿名结构体有相同的字段或者方法时，编译器采用就近访问原则访问，
 *	   如果希望访问匿名结构体的字段和方法，可以通过匿名结构体名来区分
 *	4、结构体嵌入两个（多个）匿名结构体（实现了多重继承），如果两个匿名结构体有相同的字段和方法（同时结构体本身没有相同的名字和方法），
 *	   在访问时，就必须指明匿名结构体名字，否则编译报错
 *	5、如果一个结构体嵌套了一个有名结构体，这种模式就是组合，如果是组合关系，那么，
 *	   在访问组合的结构体字段或方法时，必须带上结构体的名字
 *	6、嵌套匿名结构体后，也可以在创建结构体变量（实例）时，直接指定各个匿名结构体的字段值         
 */
func main()  {
	// 测试一
	// var b B
	// b.A.Name = "tom"
	// b.A.age = 19
	// b.A.SayOk()
	// b.A.hello()

	// 测试二
	// 直接通过 b 访问字段或方法时，其执行流程
	// 编译器判断 b 对应的类型有没有 Name，有直接调用 B 类型的 Name 字段
	// 如果没有就去看 B 中嵌入的结构体 A 有没有声明 Name，有就调用，没有继续查找...，如果找不到就报错
	// b.Name = "smith"
	// b.age = 20
	// b.SayOk()
	// b.hello()

	// 测试三
	var b B
	b.Name = "jack"	
	b.age = 23
	b.SayOk() // B SayOk... jack
	b.hello() // A hello...
	b.A.Name = "marry"
	b.A.SayOk() // A SayOk... marry

	// 测试六
	tv1 := &TV{
		Goods{"电视机1", 5000.0},
		Brand{"海尔", "沙东"},
	}
	fmt.Println(*tv1) // {{电视机1 5000} {海尔 沙东}}

	tv2 := &TV{
		Goods{
			Name : "电视机2",
			Price : 4999.99,
		},
		Brand{
			Name : "夏普",
			Address : "北京",
		},
	}
	fmt.Println(*tv2) // {{电视机2 4999.99} {夏普 北京}}

	tv3 := &TV2{
		&Goods{
			Name : "电视机3",
			Price : 999.99,
		},
		&Brand{
			Name : "创维",
			Address : "成都",
		},
	}
	fmt.Println(*tv3.Goods, *tv3.Brand) // {电视机3 999.99} {创维 成都}

	// 匿名字段是基本数据类型的使用
	var e E
	e.Name = "xxx"
	e.Age = 300
	e.int = 20
	e.n = 30
	fmt.Println("e=", e) // e= {{xxx 300} 20 30}
}