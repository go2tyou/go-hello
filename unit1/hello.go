// go 文件的后缀是 .go
// 在 go 中，每个文件都必须归属一个包
package main

// 表示引入一个包，包名 “fmt”，引入后就可使用 fmt 包的函数
import "fmt"

// func 是一个关键字，表示一个函数
// main 是函数名，是一个主函数，即为程序的入口
func main() {
	fmt.Println("hello,world!") // 表示调用 fmt 包的函数 Println 输出 "hello,world!"
}

/**
* 1、文件扩展名为 .go
* 2、main() 为程序入口
* 3、go 语言严格区分大小写
* 4、每个语句后不需要分号
* 5、每一行一条语句
* 6、定义过的变量或者 import 的包必须使用
* 7、大括号要成对出现，且和函数在同一行
 */
