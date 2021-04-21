package main

import (
	"fmt"
	"encoding/json"
)

type Monster struct {
	Name string  `json:"name"`
	Age int `json:"age"` 
	Birthday string
	Sal float64
	Skill string
}

func unmarshalStruct()  {
	str := "{\"name\":\"ssr\",\"age\":12,\"Birthday\":\"2011-11-11\",\"Sal\":12.3,\"Skill\":\"huhuhu~\"}"
	var monster Monster
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Println("unmarsal err", err)
	}
	fmt.Println("monster=", monster)
}

func unmarshalMap() {
	str := "{\"address\":\"69xxoo96\",\"age\":3,\"name\":\"xxoo\"}"
	var m map[string]interface{}
	// 反序列化不需要 make
	err := json.Unmarshal([]byte(str), &m)
	if err != nil {
		fmt.Println("unmarsal err", err)
	}
	fmt.Println("m=", m)
}

func unmarshalSlice() {
	str := "[{\"address\":\"qwer\",\"age\":1,\"name\":\"n\"}," +
		"{\"address\":\"uiop\",\"age\":2,\"name\":\"r\"}]"
	var slice []map[string]interface{}
	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Println("unmarsal err", err)
	}
	fmt.Println("slice=", slice)
}

/*
	1、在反序列化一个 json 字符串时，要确保反序列化的数据类型和原来序列化前的数据类型一致
	2、如果 json 字符串是通过程序获取到的，则不需要对 "" 号转义处理
*/
func main()  {
	unmarshalStruct() // monster= {ssr 12 2011-11-11 12.3 huhuhu~}
	unmarshalMap() // m= map[address:69xxoo96 age:3 name:xxoo]
	unmarshalSlice() // slice= [map[address:qwer age:1 name:n] map[address:uiop age:2 name:r]]
}