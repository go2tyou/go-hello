package main

import (
	"fmt"
)

/**
 * 1、字符串由单个字节连接起来的
 * 2、同一使用 UTF-8 编码
 * 3、字符串不可变
 * 4、可用两种形式表现
 *		双引号：会识别转义符号
 *		反引号：原生形式输出
 */

// go 中字符串类型的使用
func main() {
	var addr string = "你好啊老表"
	fmt.Println("addr=", addr)

	// var str = "hello"
	// str[0] = 'g' // cannot assign to str[0] --> 即不可变

	// 字符串拼接
	var xx = "hello " + "world"
	xx += " !"
	fmt.Println("xx=", xx)

	// 拼接操作很长，分行，每行最后要 + 号
	var yy = "lululu" + "lululu" +
		"lululu" + "lululu" +
		"lululu" + "lululu"
	fmt.Println("yy=", yy)

}
