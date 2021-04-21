package main

import (
	"fmt"
)

// 键盘输入语句
func main()  {
	// 方式一
	var name string
	var age byte
	var sal float32
	var isPass bool
	// fmt.Println("请输入姓名：")
	// fmt.Scanln(&name)
	// fmt.Println("请输入年龄：")
	// fmt.Scanln(&age)
	// fmt.Println("请输入薪水：")
	// fmt.Scanln(&sal)
	// fmt.Println("是否通过：")
	// fmt.Scanln(&isPass)
	// fmt.Printf("名字：%v\n年龄：%v\n薪水：%v\n是否通过：%v\n", name, age, sal, isPass)
	
	// 方式二
	// 可以按指定格式输入
	fmt.Println("请输入xxx,使用空格隔开")
	fmt.Scanf("%s %d %f %t", &name, &age, &sal, &isPass)
	fmt.Printf("名字：%v\n年龄：%v\n薪水：%v\n是否通过：%v\n", name, age, sal, isPass)

}