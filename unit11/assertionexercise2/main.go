package main

import (
	"fmt"
)

type Student struct {

}

func TypeJudge(items... interface{})  {
	for index, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("index=%v type bool x=%v\n", index, x)
		case float32:
			fmt.Printf("index=%v type float32 x=%v\n", index, x)
		case float64:
			fmt.Printf("index=%v type float64 x=%v\n", index, x)
		case int, int32, int64:
			fmt.Printf("index=%v type int x=%v\n", index, x)
		case string:
			fmt.Printf("index=%v type string x=%v\n", index, x)
		case Student:
			fmt.Printf("index=%v type Student x=%v\n", index, x)
		case *Student:
			fmt.Printf("index=%v type *Student x=%v\n", index, x)
		default:
			fmt.Printf("index=%v type unknow x=%v\n", index, x)
		}
	}
}

func main()  {
	var n1 float32 = 1.1
	var n2 float64 = 1.1
	var n3 int32 = 10
	var name string = "ssr"
	address := "sr"
	n4 := 200
	stu1 := Student{}
	stu2 := &Student{}
	TypeJudge(n1, n2, n3, name, address, n4, stu1, stu2)
	// index=0 type float32 x=1.1
	// index=1 type float64 x=1.1
	// index=2 type int x=10
	// index=3 type string x=ssr
	// index=4 type string x=sr
	// index=5 type int x=200
	// index=6 type Student x={}
	// index=7 type *Student x=&{}
}