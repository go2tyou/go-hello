package main

import (
	"fmt"
)

// 如果结构体的字段类型：指针、slice、map 的零值都是 nil，即还没有分配空间
// 如果需要使用这样的字段，需要先 make 才能使用
// 结构体是值类型
type Person struct {
	Name string
	Age int
	Scores [5]float64
	ptr *int
	slice []int
	map1 map[string]string
}

type Monster struct {
	Name string
	Age int
}

func main()  {
	var p1 Person
	fmt.Println("p1=", p1) // p1= { 0 [0 0 0 0 0] <nil> [] map[]}

	// 使用 slice 前，一定要先 make
	// p1.slice[0] = 100 // index out of range [0] with length 0
	p1.slice = make([]int, 10)
	p1.slice[0] = 100
	fmt.Println("p1=", p1) // p1= { 0 [0 0 0 0 0] <nil> [100 0 0 0 0 0 0 0 0 0] map[]}

	p1.map1 = make(map[string]string)
	p1.map1["key1"] = "tom"
	fmt.Println("p1=", p1) // p1= { 0 [0 0 0 0 0] <nil> [100 0 0 0 0 0 0 0 0 0] map[key1:tom]}

	var monster1 Monster
	monster1.Name = "xxx"
	monster1.Age = 500

	monster2 := monster1 // 结构体是值类型，默认是值拷贝
	monster2.Name = "yyy"

	fmt.Println("monster1=", monster1) // monster1= {xxx 500}
	fmt.Println("monster2=", monster2) // monster2= {yyy 500}
}