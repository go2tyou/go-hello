package main

import (
	"fmt"
)

// 1、对 for 循环来说，有 4 个要素
// 		循环变量初始化、循环条件、循环（操作）语句、循环变量迭代
// 2、for 循环的执行顺序
// 		1、执行变量初始化 -> 2、执行循环条件 -> 3、如果循环条件为真，就执行循环操作 ->
//		-> 4、执行循环变量迭代 -> 5、反复执行 2，3，4 步骤，循环条件为 false，退出循环

// 使用细节
// 1、循环条件是返回一个布尔值的表达式
// 2、可以将变量初始化和变量迭代写到其他位置
// 3、无限循环配合 break
// 4、字符串遍历 for-range
func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println("daisuki~", i)
	}

	j := 1
	for j <= 10 {
		fmt.Println("suki~", j)
		j++
	}

	// 无限循环配合 break
	k := 1
	for {
        if k <= 10 {
			fmt.Println("lululu~", k)	
		} else {
			break // 跳出 for 循环
		}     
		k++
	}

	// // 字符串遍历方式一
	// var str string = "hello,world!"
	// for i := 0; i < len(str); i++ {
	// 	fmt.Printf("%c\n", str[i])
	// }

	// 字符串中含有中文，需要 str 转为 切片
	// 字符串遍历方式一 -> 按字节方式遍历
	var str string = "hello,world!噜皮"
	str2 := []rune(str)
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c\n", str2[i])
	}

	// 字符串遍历方式二 for-range -> 按字符方式遍历
	str = "asdfghjkl问号"
	for index, val := range str {
		fmt.Printf("index=%d val=%c\n", index, val)
	} 
}