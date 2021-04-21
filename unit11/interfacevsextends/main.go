package main

import (
	"fmt"
)

type Monkey struct {
	Name string
}

type BirdAble interface {
	Flying()
}

type FishAble interface {
	Swimming()
}

func (this *Monkey) Climbing() {
	fmt.Println(this.Name, "Climbing...")
}

type LitteMonkey struct {
	Monkey // 继承
}

// 实现接口
func (this *LitteMonkey) Flying() {
	fmt.Println(this.Name, "Flying...")
}

func (this *LitteMonkey) Swimming() {
	fmt.Println(this.Name, "Swimming...")
}
// 1、当A结构体继承了B结构体，那么就继承了B结构体的字段和方法，并且可以直接使用
// 2、实现接口可以在不破坏继承关系的基础上，对一个结构体进行功能的扩展，实现接口是对继承的补充
/*
 *	1、继承：解决代码的复用性和可维护性
 *	2、接口：设计、灵活、解耦
 */
func main()  {
	monkey := LitteMonkey {
		Monkey {
			Name : "悟空",
		},
	}
	monkey.Climbing() // 悟空 Climbing...
	monkey.Flying() // 悟空 Flying...
	monkey.Swimming() // 悟空 Swimming...
}