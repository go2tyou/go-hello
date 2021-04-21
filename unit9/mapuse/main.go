package main

import (
	"fmt"
)

// map 的使用
func main()  {
	// 方式一
	// 声明 -> make -> 赋值和使用

	// 方式二
	cities := make(map[string]string)
	cities["no1"] = "xx"
	cities["no2"] = "yy"
	cities["no3"] = "zz"
	fmt.Println("cities=", cities) // cities= map[no1:xx no2:yy no3:zz]

	// 方式三
	heros := map[string]string{
		"h1" : "s",
		"h2" : "s+",
		"h3" : "ss",
		"h4" : "ss+",
	}
	fmt.Println("heros=", heros) // heros= map[h1:s h2:s+ h3:ss h4:ss+]
}