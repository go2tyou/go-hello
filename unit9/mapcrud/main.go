package main

import (
	"fmt"
)

func main()  {
	cities := make(map[string]string)
	cities["no1"] = "xx"
	cities["no2"] = "yy"
	cities["no3"] = "zz"
	fmt.Println("cities=", cities) // cities= map[no1:xx no2:yy no3:zz]
	// no3 这个 key 已经存在，因此下句是修改
	cities["no3"] = "zz~"
	fmt.Println("cities=", cities) // cities= map[no1:xx no2:yy no3:zz~]

	// map 删除
	// delete(cities, "no1")
	// fmt.Println("cities=", cities) // cities= map[no2:yy no3:zz~]
	// no4 不存在，删除不会操作，也不会报错
	// delete(cities, "no4")
	// fmt.Println("cities=", cities)

	// 如果想删除所有 key
	// 1、遍历所有的 key，然后逐一删除
	// 2、直接 make 一个新的空间
	// cities = make(map[string]string)
	// fmt.Println("cities=", cities) // cities= map[]

	// map 查找
	val, ok := cities["no1"]
	if ok {
		fmt.Println("ok... no1=", val) // ok... no1= xx
	} else {
		fmt.Println("no...")
	}

	// map 遍历 -> for...range
	for k, v := range cities {
		fmt.Printf("k=%v v=%v\n", k, v)
	}
	fmt.Println("cities len=", len(cities)) // cities len= 3
}