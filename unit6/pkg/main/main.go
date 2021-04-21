package main

import (
	"fmt"
	"gocode/helloworld/unit6/pkg/utils"
)

// go 的每一个文件都属于一个包
// 包的本质：创建不同的文件夹，来存放程序文件
/*
 *	作用：
 *		1、区分相同名字的函数、变量等标识符
 *		2、当程序文件很多时，可以很好的管理目
 *		3、控制函数、变量的访问范围，即作用域
 */

/*
 *	使用细节：
 *		1、文件的包名通常和文件所在的文件名一致
 *		2、一个文件要使用其他包的函数或变量时，应该先引入对应的包
 *		3、package 在第一行，然后是 import
 *		4、在 import 包时，路径从 src 下开始定位的
 *		5、为了其他包能访问，函数名的首字母大写
 *		6、访问方式：包名.函数名
 *		7、包名太长，支持别名，但原来的包名就不能使用了
 *		8、同一包下，不能有相同的函数名、全局变量名
 *		9、如果要编译成一个可执行文件，需要将这个包声明为 package main
 *		10、go build -o xx\xxx.exe xxxx/mian
 *		11、build 会在 pkg 下生成 xxx.a 库文件
 */
func main() {
	var x float64 = 1
	var y float64 = 2
	var op byte = '+'
	r := utils.Cal(x, y, op)
	fmt.Printf("r=%f\n", r)
}
