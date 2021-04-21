package main

import (
	"fmt"
	"strconv"
)

/**
 * 注意：
 * 	   1、确保 string 类型能转成有效的数据
 *	   2、没转成功，结果为默认值
 */

// go 中基本数据类型与 string 类型互转
func main() {
	// string --> 基本数据类型
	var str string = "true"
	var b bool
	// ParseBool(str) 函数会返回两个值 (value bool, err error)
	// 只想要 val, 不想 err, 使用 _ 忽略
	b, _ = strconv.ParseBool(str)
	fmt.Printf("b type is %T b=%t\n", b, b)
	
	var str2 string = "1234"
	var num int64
	num, _ = strconv.ParseInt(str2, 10, 64)
	fmt.Printf("num type is %T b=%d\n", num, num)

	var str3 string = "123.456"
	var f1 float64
	f1, _ = strconv.ParseFloat(str3, 64)
	fmt.Printf("num type is %T b=%f\n", f1, f1)

}
