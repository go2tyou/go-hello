package main

import (
	"fmt"
)

type AInterfance interface {
	Say()
}

type BInterfance interface {
	Hello()
}

type Stu struct {
	Name string
}

func (stu Stu) Say() {
	fmt.Println("stu Say()...")
}

type integer int

func (i integer) Say() {
	fmt.Println("integer Say()...i=", i)
}

type Monster struct {

}

func (m Monster) Hello() {
	fmt.Println("Monster Hello()...")
}

func (m Monster) Say() {
	fmt.Println("Monster Say()...")
}

type T interface {

}

/*
 *	1、接口本身不能创建实例，但是可以指向一个实现了该接口的自定义类型的变量（实例）
 *	2、接口中所有的方法都没有方法体，即都是没有实现的方法
 *	3、在 go 中，一个自定义类型需要将某个接口的所有方法都实现，才说这个自定义类型实现了该接口
 *	4、只要是自定义数据类型，就可以实现接口，不仅仅是结构体类型
 *	5、一个自定义类型可以实现多个接口
 *	6、go 接口中不能有任何变量
 *	7、一个接口（A接口）可以继承多个别的接口（B，C接口），这时如果要实现A接口，也必须将B，C接口的方法全部实现
 *	8、interface 类型默认是一个指针（引用类型），如果没有对 interface 初始化就使用，那么会输出 nil
 *	9、空接口 interface{} 没有任何方法，所以所有类型都实现了空接口，即可以把任何一个变量赋值给空接口
 *	10、一个接口继承多个接口时，方法不能重复
 */
func main()  {

	// 测试 1
	var stu Stu
	var a AInterfance = stu
	a.Say() // stu Say()...

	// 测试 4
	var i integer = 10
	var b AInterfance = i
	b.Say() // integer Say()...i= 10

	// 测试 5
	var monster Monster
	var aa AInterfance = monster
	var bb BInterfance = monster
	aa.Say() // Monster Say()...
	bb.Hello() // Monster Hello()...

	// 测试 9
	var t T = stu
	var t2 interface{} = stu
	fmt.Println(t) // {}
	fmt.Println(t2) // {}
	var num float64 = 1.2
	t = num
	t2 = num
	fmt.Println(t) // 1.2
	fmt.Println(t2) // 1.2
}