package main

import (
	"fmt"
	"gocode/helloworld/unit10/factory/model"
)

// 工厂模式
// 结构体首字母大写，引入后，可以直接使用
// 结构体首字母小写，引入后，不能直接使用
func main()  {
	// var stu = model.Student {
	// 	Name : "tom",
	// 	Score : 78.9,
	// }
	// fmt.Println(stu) // {tom 78.9}

	var stu1 = model.NewStudent("ssr", 99.99)
	fmt.Println(*stu1) // {ssr 99.99}
}