package main

import (
	"fmt"
	"gocode/helloworld/unit11/encapsulate/model"
)

// 封装
func main()  {
	p := model.NewPerson("smith")
	fmt.Println(*p) // {smith 0 0}
	// p.age = 18 // p.age undefined (cannot refer to unexported field or method age)
	p.SetAge(18)
	p.SetSal(5000.0)
	fmt.Println("Name=", p.Name, "age=", p.GetAge(), "sal=", p.GetSal()) // Name= smith age= 18 sal= 5000

	acc := model.NewAccount("xmzsss", "666666", 250.0)
	fmt.Println(*acc) // {xmzsss 666666 250}
}