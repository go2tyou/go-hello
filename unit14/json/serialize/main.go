package main

import (
	"fmt"
	"encoding/json"
)

// 希望序列化后 key 可自定义，使用 tag
type Monster struct {
	Name string  `json:"name"`
	Age int `json:"age"` 
	Birthday string
	Sal float64
	Skill string
}

func testStruct()  {
	monster := Monster {
		Name : "ssr",
		Age : 12,
		Birthday : "2011-11-11",
		Sal : 12.3,
		Skill : "huhuhu~",
	}
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Println("marshal err", err)
	}
	fmt.Printf("%s\n", data)
}

func testMap()  {
	var a map[string]interface{}
	a = make(map[string]interface{})
	a["name"] = "xxoo"
	a["age"] = 3
	a["address"] = "69xxoo96"
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Println("marshal err", err)
	}
	fmt.Printf("%s\n", data)
}

func testSlice()  {
	var slice []map[string]interface{}

	var m1 map[string]interface{}
	m1 = make(map[string]interface{})
	m1["name"] = "n"
	m1["age"] = 1
	m1["address"] = "qwer"
	slice = append(slice, m1)

	var m2 map[string]interface{}
	m2 = make(map[string]interface{})
	m2["name"] = "r"
	m2["age"] = 2
	m2["address"] = "uiop"
	slice = append(slice, m2)

	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Println("marshal err", err)
	}
	fmt.Printf("%s\n", data)
}

// 对基本数据类型序列化意义不大
func testFloat64()  {
	var num1 float64 = 23.4445
	data, err := json.Marshal(num1)
	if err != nil {
		fmt.Println("marshal err", err)
	}
	fmt.Printf("%s\n", data)
}

func main()  {
	testStruct() // {"Name":"ssr","Age":12,"Birthday":"2011-11-11","Sal":12.3,"Skill":"huhuhu~"}
	testMap() // {"address":"69xxoo96","age":3,"name":"xxoo"}
	testSlice() // [{"address":"qwer","age":1,"name":"n"},{"address":"uiop","age":2,"name":"r"}]
	testFloat64() // 23.4445
}