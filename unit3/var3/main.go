package main

import "fmt"

// 在 go 中函数外部定义变量就是全局变量
// 声明全局变量
var lsp = 777
var dalao = "wo"

// 也可一次声明
var (
	a = 222
	b = 333
	c = "www"
)

func main() {
	// go 如何一次性声明多个变量
	var n1, n2, n3 int

	fmt.Println("n1=", n1)
	fmt.Println("n2=", n2)
	fmt.Println("n3=", n3)

	// 方式二
	var x, y, z = 11, "dd", 44
	fmt.Println("x=", x)
	fmt.Println("y=", y)
	fmt.Println("z=", z)

	// 方式三（类型推导）
	o, p, q := 66, "sm", 99
	fmt.Println("o=", o)
	fmt.Println("p=", p)
	fmt.Println("q=", q)

	// 输出全局变量
	fmt.Println("lsp=", lsp)
	fmt.Println("dalao=", dalao)
	fmt.Println("a=", a, "b=", b, "c=", c)

}
