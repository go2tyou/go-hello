package main

import (
	"fmt"
)

func sum(n1 int, n2 int) int {
	// 当执行到 defer 时，暂时不执行，会将 defer 后面的语句压入到独立的栈（defer 栈）
	// 当函数执行完毕后，再从 defer 栈，按照先入后出的方式出栈执行s
	defer fmt.Println("ok1 n1=", n1)
	defer fmt.Println("ok2 n2=", n2)

	// 验证细节 3
	// n1++
	// n2++
	// /// 结果
	// // ok3 res= 32
	// // ok2 n2= 20
	// // ok1 n1= 10
	// // res= 32

	res := n1 + n2
	fmt.Println("ok3 res=", res)
	return res
}

// 为什么使用 defer
// 为了在函数执行完毕后，及时释放资源
/*
 *	细节：
 *	1、当 go 执行到一个 defer 时，不会立即执行 defer 后的语句，而是将 defer 后的语句压入到一个栈中，然后继续执行函数下一个语句
 *	2、当函数执行完毕后，再从 defer 栈中，依次从栈顶取出语句执行（遵循先入后出）
 *	3、在 defer 将语句放入到栈时，也会将相关的值拷贝，同时入栈
 *	4、创建资源后（比如打开文件、获取数据库连接、锁），可以执行 xxx.Close()
 *  5、在 defer 后，可以继续使用创建资源
 *  6、当函数执行完毕后，系统会依次从 defer 栈中，取出语句，关闭资源
 *  7、不在意关闭资源的时机了
 */
func main()  {
	res := sum(10, 20)
	fmt.Println("res=", res)
}

/// 运行结果：
// ok3 res= 30
// ok2 n2= 20
// ok1 n1= 10
// res= 30