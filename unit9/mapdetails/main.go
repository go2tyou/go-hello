package main

import (
	"fmt"
)

func modify(map1 map[int]int)  {
	map1[10] = 100
}

// 定义一个结构体
type Stu struct {
	Name string
	Age int
	Address string
}

// map 使用细节
/*
 *	1、map 是引用类型，遵循引用传递的机制
 *	2、map 的容量达到后，再增加元素，会自动扩容
 *	3、map 的 value 常用 struct 类型，更适合管理复杂的数据
 */
func main()  {
	// 测试 1
	map1 := make(map[int]int)
	map1[1] = 99
	map1[8] = 88
	map1[7] = 77
	map1[10] = 11
	modify(map1)
	fmt.Println("map1=", map1) // map1= map[1:99 7:77 8:88 10:100]

	// 测试 3
	stus := make(map[string]Stu, 10)
	stu1 := Stu{"tom", 17, "sh"}
	stu2 := Stu{"marry", 27, "sz"}
	stus["no1"] = stu1
	stus["no2"] = stu2
	fmt.Println("stus=", stus) // stus= map[no1:{tom 17 sh} no2:{marry 27 sz}]
	// 遍历
	for k, v := range stus {
		fmt.Printf("no=%v\n", k)
		fmt.Printf("name=%v\n", v.Name)
		fmt.Printf("age=%v\n", v.Age)
		fmt.Printf("address=%v\n", v.Address)
		fmt.Println()
	}

}