package main

import (
	"fmt"
	"sort"
)

func main()  {
	map1 := make(map[int]int, 10)
	map1[10] = 100
	map1[1] = 10
	map1[4] = 34
	map1[8] = 75
	fmt.Println("map1=", map1) // map1= map[1:10 4:34 8:75 10:100]

	// 按 key 的顺序排序输出
	// 1、现将 map 的 key 放入到切片中
	// 2、对切片排序
	// 3、遍历切片，然后按照 key 来输出 map 的值
	var keys []int
	for k, _ := range map1 {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	fmt.Println("keys=", keys)
	for _, k := range keys {
		fmt.Printf("map1[%v]=%v\n", k, map1[k])
	}
}