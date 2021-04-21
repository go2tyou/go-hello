package main

import (
	"fmt"
)

// slice of map
func main()  {
	var monsters []map[string]string
	monsters = make([]map[string]string, 2)
	if monsters[0] == nil {
		monsters[0] = make(map[string]string, 2)
		monsters[0]["name"] = "aaa"
		monsters[0]["age"] = "100"
	}
	if monsters[1] == nil {
		monsters[1] = make(map[string]string, 2)
		monsters[1]["name"] = "bbb"
		monsters[1]["age"] = "200"
	}

	// append 函数，可以动态增加
	newMonster :=  map[string]string {
		"name" : "ccc",
		"age" : "300",
	}
	monsters = append(monsters, newMonster) // monsters= [map[age:100 name:aaa] map[age:200 name:bbb] map[age:300 name:ccc]]
	fmt.Println("monsters=", monsters)
}