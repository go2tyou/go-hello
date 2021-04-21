package main

import (
	"fmt"
	"encoding/json"
)

type Point struct {
	x int
	y int
}

type Rect struct {
	leftUp, rightDown Point
}

type Rect2 struct {
	leftUp, rightDown *Point
}

type A struct {
	Num int
}

type B struct {
	Num float64
}

type Monster struct {
	Name string `json:"name"` // `json:"name"` 即 struct tag
	Age int `json:"age"`
	Skill string `json:"skill"`
}

/*
 *	1、结构体的所有字段在内存中是连续的
 *	2、结构体是用户单独定义的类型，和其他类型进行转换时需要有完全相同的字段
 *	3、结构体进行 type 重新定义（相当于取别名），go 认为是新的数据类型，但是相互间可以强转
 *	4、strut 的每个字段上，可以写上一个 tag，该 tag 可以通过反射机制获取，如序列化和反序列化
 */
func main()  {
	// 测试 1
	r1 := Rect{Point{1, 2}, Point{3, 4}}
	// r1 有 4 个 int，在内存中是连续分布的
	fmt.Printf("r1.leftUp.x addr=%p r1.leftUp.y addr=%p r1.rightDown.x addr=%p r1.rightDown.y addr=%p\n", 
		&r1.leftUp.x, &r1.leftUp.y, &r1.rightDown.x, &r1.rightDown.y)
	// r1.leftUp.x addr=0xc000010400 r1.leftUp.y addr=0xc000010408 r1.rightDown.x addr=0xc000010410 r1.rightDown.y addr=0xc000010418
	
	// r2 有 2 个 *Point 类型，这 2 个 *Point 类型本身的地址也是连续的，但是指向的地址不一定是连续的
	r2 := Rect2{&Point{10, 20}, &Point{30, 40}}
	fmt.Printf("r2.leftUp addr=%p r2.rightDown addr=%p\n", &r2.leftUp, &r2.rightDown) // r2.leftUp addr=0xc00003e1f0 r2.rightDown addr=0xc00003e1f8
	fmt.Printf("r2.leftUp addr=%p r2.rightDown addr=%p\n", r2.leftUp, r2.rightDown) // r2.leftUp addr=0xc0000120a0 r2.rightDown addr=0xc0000120b0

	// 测试 2
	var a A
	var b B
	// a = b // cannot use b (type B) as type A in assignment
	// a = A(b) // cannot convert b (type B) to type A -> 结构体的字段要完全一样（名称、个数、类型）
	fmt.Println(a, b)

	// 测试 4
	monster := Monster{"xxx", 500, "aowo~"}
	// json.Marshal 函数中使用反射
	jsonStr, _ := json.Marshal(monster)
	fmt.Println("jsonStr=", string(jsonStr)) // jsonStr= {"Name":"xxx","Age":500,"Skill":"aowo~"}
	// Name string `json:"name"` => jsonStr= {"name":"xxx","age":500,"skill":"aowo~"}
}