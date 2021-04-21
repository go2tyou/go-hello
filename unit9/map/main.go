package main

import (
	"fmt"
)

// map 映射
/*
 *	1、var map变量名 map[keytype]valuetype
 *	2、声明是不会分配内存的，初始化需要 make，分配内存后才能赋值和使用
 *	3、key 不可以重复，如果重复以最后的 key 对应的 value 为主
 *	4、value 可以重复
 *	5、map 的 key-value 是无序的
 *	6、make 默认大小为 1
 */
func main()  {
	var a map[string]string
	// a["no1"] = "ss" // assignment to entry in nil map
	// fmt.Println("a=", a) // a= map[]
	a = make(map[string]string, 10)
	a["no1"] = "ss"
	a["no2"] = "bb"
	a["no1"] = "nn" // key 不可以重复，如果重复以最后的 key 对应的 value 为主
	a["no3"] = "bb" // value 可以重复
	fmt.Println("a=", a) // a= map[no1:ss] // map 的 key-value 是无序的
}